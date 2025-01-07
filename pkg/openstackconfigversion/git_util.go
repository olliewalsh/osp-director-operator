/*
Copyright 2021 Red Hat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstackconfigversion

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	git "github.com/go-git/go-git/v5"
	config "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/capability"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/go-logr/logr"
	ospdirectorv1beta1 "github.com/openstack-k8s-operators/osp-director-operator/api/v1beta1"
	"github.com/openstack-k8s-operators/osp-director-operator/pkg/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Patch is an implementation of fdiff.Patch interface
type Patch struct {
	message     string
	filePatches []diff.FilePatch
}

// FilePatches .
func (p Patch) FilePatches() []diff.FilePatch {
	return p.filePatches
}

// Message .
func (p Patch) Message() string {
	return p.message
}

// filterPatches internal function which is used to strip out '_server_id', ansible inventory, and config changes from our git diff results
func filterPatches(filePatches []diff.FilePatch) []diff.FilePatch {
	var retPatches []diff.FilePatch

	m1 := regexp.MustCompile(`_server_id:`)
	for _, fp := range filePatches {
		from, to := fp.Files()

		if len(fp.Chunks()) == 0 {
			continue
		}
		if from == nil {
			// New File is created.
			retPatches = append(retPatches, fp)
		} else if to == nil {
			// File is deleted.
			retPatches = append(retPatches, fp)

		} else if from.Path() == to.Path() {

			if strings.HasSuffix("tripleo-ansible/tripleo-ansible-inventory.yaml", from.Path()) || strings.HasSuffix("tripleo-ansible/ansible.cfg", from.Path()) {
				continue
			}

			var chunks []diff.Chunk
			for _, chunk := range fp.Chunks() {
				chunk.Type()
				s := chunk.Content()
				if len(s) == 0 {
					continue
				}

				if chunk.Type() == diff.Add {

					if !m1.Match([]byte(s)) {
						chunks = append(chunks, chunk)
					}
				}

			}
			if len(chunks) > 0 {
				retPatches = append(retPatches, fp)
			}
		} else {
			// file is renamed
			retPatches = append(retPatches, fp)
		}
	}

	return retPatches
}

// SyncGit func
func SyncGit(
	ctx context.Context,
	inst *ospdirectorv1beta1.OpenStackConfigGenerator,
	remoteOptions *common.GitRemoteOptions,
	client client.Client,
	log logr.Logger,
) (map[string]ospdirectorv1beta1.OpenStackConfigVersion, error) {

	configVersions := make(map[string]ospdirectorv1beta1.OpenStackConfigVersion)

	gitCloneOptions := git.CloneOptions{
		URL:          remoteOptions.URL,
		Auth:         remoteOptions.Auth,
		CABundle:     remoteOptions.CABundle,
		ProxyOptions: remoteOptions.ProxyOptions,
	}

	repo, err := git.Clone(memory.NewStorage(), nil, &gitCloneOptions)
	// if Azure DevOps is used it can fail with as azure is not compatible to go-git, https://github.com/go-git/go-git/pull/613
	// "2023-08-04T13:16:19.264Z        INFO    controllers.OpenStackConfigGenerator    Failed to create Git repo: empty git-upload-pack given"
	// retry with workaround setting capability.ThinPack
	if err != nil && errors.Is(err, transport.ErrEmptyUploadPackRequest) {
		log.Info(fmt.Sprintf("Failed to create Git repo: %s\n", err.Error()))
		log.Info("retrying with capability.ThinPack transport capability, required for Azure DevOps")
		transport.UnsupportedCapabilities = []capability.Capability{
			capability.ThinPack,
		}

		repo, err = git.Clone(memory.NewStorage(), nil, &gitCloneOptions)
	}
	// Failed to create Git repo: URL field is required
	if err != nil {
		log.Info(fmt.Sprintf("Failed to create Git repo: %s\n", err.Error()))
		return nil, err
	}

	// Create the remote with repository URL
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteOptions.URL},
	})

	refs, err := rem.List(&git.ListOptions{
		Auth:         remoteOptions.Auth,
		CABundle:     remoteOptions.CABundle,
		ProxyOptions: remoteOptions.ProxyOptions,
	})
	if err != nil {
		log.Info(fmt.Sprintf("Failed to list remote: %s\n", err.Error()))
		return nil, err
	}

	m1 := regexp.MustCompile(`/`)
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			if ref.Name() == "refs/heads/master" || ref.Name() == "HEAD" {
				continue
			}
			commit, err := repo.CommitObject(ref.Hash())
			if err != nil {
				log.Info(fmt.Sprintf("Failed to get commit object: %s\n", err.Error()))
				return nil, err
			}

			latest, err := repo.Tag("latest")
			if err != nil {
				log.Info(fmt.Sprintf("Failed to get 'latest' tag: %s\n. No git diffs will be calculated.", err.Error()))
			}
			var configVersion ospdirectorv1beta1.OpenStackConfigVersion

			if latest != nil {
				commitLatest, err := repo.CommitObject(latest.Hash())
				if err != nil {
					return nil, err
				}
				patch, err := commit.PatchContext(ctx, commitLatest)

				filterPatch := Patch{
					message:     patch.Message(),
					filePatches: filterPatches(patch.FilePatches()),
				}
				if err != nil {
					return nil, err
				}
				buffer := bytes.NewBuffer(nil)
				e := diff.NewUnifiedEncoder(buffer, 0)
				err = e.Encode(filterPatch)
				if err != nil {
					return nil, err
				}
				diff := truncateDiff(buffer.String(), log)
				configVersion = ospdirectorv1beta1.OpenStackConfigVersion{
					ObjectMeta: metav1.ObjectMeta{
						Name:      m1.Split(ref.Name().String(), -1)[2],
						Namespace: inst.Namespace,
					},
					Spec: ospdirectorv1beta1.OpenStackConfigVersionSpec{Hash: m1.Split(ref.Name().String(), -1)[2], Diff: diff, ConfigGeneratorName: inst.Name}}
			} else {
				configVersion = ospdirectorv1beta1.OpenStackConfigVersion{
					ObjectMeta: metav1.ObjectMeta{
						Name:      m1.Split(ref.Name().String(), -1)[2],
						Namespace: inst.Namespace,
					},
					Spec: ospdirectorv1beta1.OpenStackConfigVersionSpec{Hash: m1.Split(ref.Name().String(), -1)[2], Diff: "", ConfigGeneratorName: inst.Name}}
			}
			configVersions[ref.Hash().String()] = configVersion
		}

	}

	return configVersions, nil
}

// truncateDiff  truncate the diff size to less than 512KB
func truncateDiff(diff string, log logr.Logger) string {
	if len(diff) > 524800 {
		log.Info(fmt.Sprintf("Git diff was truncated to 512KB.\n %s", diff[:524800]))
		return diff[:524800]
	}
	return diff
}
