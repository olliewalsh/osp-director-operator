/*
Copyright 2024 Red Hat

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

package common

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-logr/logr"
	crypto_ssh "golang.org/x/crypto/ssh"
	corev1 "k8s.io/api/core/v1"
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GitRemoteOptions
type GitRemoteOptions struct {
	URL      string
	Auth     transport.AuthMethod
	CABundle []byte
}

func getGitSecret(
	ctx context.Context,
	gitSecret types.NamespacedName,
	client client.Client,
	log logr.Logger,
) (*corev1.Secret, error) {
	// Check if this Secret already exists
	foundSecret := &corev1.Secret{}
	err := client.Get(ctx, gitSecret, foundSecret)
	if err != nil {
		return nil, err
	}
	return foundSecret, nil
}

func BuildGitScriptEnvVars(
	ctx context.Context,
	gitSecret types.NamespacedName,
	client client.Client,
	log logr.Logger,
) ([]corev1.EnvVar, error) {
	foundSecret, err := getGitSecret(ctx, gitSecret, client, log)
	if err != nil {
		if k8s_errors.IsNotFound(err) {
			log.Error(err, "GitRepo secret was not found.")
			return nil, err
		}
	}
	log.Info("GitRepo foundSecret")

	gitURL := string(foundSecret.Data["git_url"])
	gitEndpoint, err := transport.NewEndpoint(gitURL)
	if err != nil {
		log.Info(fmt.Sprintf("parse git url failed: %s\n", err.Error()))
		return nil, err
	}

	gitScriptEnvVars := []corev1.EnvVar{
		{
			Name:  "GIT_URL",
			Value: gitURL,
		},
	}

	if strings.HasPrefix(gitEndpoint.Protocol, "http") {
		gitScriptEnvVars = append(gitScriptEnvVars, corev1.EnvVar{
			Name: "GIT_API_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: gitSecret.Name,
					},
					Key: "git_api_key",
				},
			},
		})
	} else {
		gitScriptEnvVars = append(gitScriptEnvVars, corev1.EnvVar{
			Name: "GIT_ID_RSA",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: gitSecret.Name,
					},
					Key: "git_ssh_identity",
				},
			},
		})
	}
	return gitScriptEnvVars, nil
}

func BuildGitOptions(
	ctx context.Context,
	gitSecret types.NamespacedName,
	CABundle []byte,
	client client.Client,
	log logr.Logger,
) (*GitRemoteOptions, error) {
	// Check if this Secret already exists
	foundSecret, err := getGitSecret(ctx, gitSecret, client, log)
	if err != nil {
		if k8s_errors.IsNotFound(err) {
			log.Error(err, "GitRepo secret was not found.")
			return nil, err
		}
	}
	log.Info("GitRepo foundSecret")

	pkey := foundSecret.Data["git_ssh_identity"]
	apikey := string(foundSecret.Data["git_api_key"])

	gitURL := string(foundSecret.Data["git_url"])
	gitEndpoint, err := transport.NewEndpoint(gitURL)
	if err != nil {
		log.Info(fmt.Sprintf("parse git url failed: %s\n", err.Error()))
		return nil, err
	}

	remoteOptions := GitRemoteOptions{
		URL: gitURL,
	}

	if strings.HasPrefix(gitEndpoint.Protocol, "http") {
		// Username is typically irrelevant for apikey auth
		userName := "notused"
		// But use username from url if it is set e.g https://foo@mygit.example.com/...
		if gitEndpoint.User != "" {
			userName = gitEndpoint.User
		}
		if gitEndpoint.Password != "" {
			// Not secure, password would be visible on git command line from create-playbooks.sh
			err := errors.New("git url includes password")
			log.Info(fmt.Sprintf("insecure git url (use git_api_key): %s\n", err.Error()))
			return nil, err
		}

		remoteOptions.Auth = &http.BasicAuth{
			Username: userName,
			Password: apikey,
		}

	} else {
		publicKeys, err := ssh.NewPublicKeys(gitEndpoint.User, pkey, "")
		publicKeys.HostKeyCallback = crypto_ssh.InsecureIgnoreHostKey()
		if err != nil {
			log.Info(fmt.Sprintf("generate publickeys failed: %s\n", err.Error()))
			return nil, err
		}
		remoteOptions.Auth = publicKeys
	}
	remoteOptions.CABundle = CABundle
	return &remoteOptions, nil
}
