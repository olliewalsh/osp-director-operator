/*


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

// Generated by:
//
// operator-sdk create webhook --group osp-director --version v1beta1 --kind OpenStackProvisionServer --programmatic-validation
//

package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var openstackprovisionserverlog = logf.Log.WithName("openstackprovisionserver-resource")

// SetupWebhookWithManager - register this webhook with the controller manager
func (r *OpenStackProvisionServer) SetupWebhookWithManager(mgr ctrl.Manager, defaults OpenStackProvisionServerDefaults) error {

	if webhookClient == nil {
		webhookClient = mgr.GetClient()
	}

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-osp-director-openstack-org-v1beta1-openstackprovisionserver,mutating=false,failurePolicy=fail,sideEffects=None,groups=osp-director.openstack.org,resources=openstackprovisionservers,verbs=create;update;delete,versions=v1beta1,name=vopenstackprovisionserver.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &OpenStackProvisionServer{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackProvisionServer) ValidateCreate() error {
	openstackprovisionserverlog.Info("validate create", "name", r.Name)

	if err := checkBackupOperationBlocksAction(r.Namespace, APIActionCreate); err != nil {
		return err
	}

	return r.validateCr()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackProvisionServer) ValidateUpdate(old runtime.Object) error {
	openstackprovisionserverlog.Info("validate update", "name", r.Name)

	return r.validateCr()
}

func (r *OpenStackProvisionServer) validateCr() error {
	existingPorts, err := r.GetExistingProvServerPorts(&Condition{}, webhookClient)
	if err != nil {
		return err
	}

	for name, port := range existingPorts {
		if name != r.Name && port == r.Spec.Port {
			return fmt.Errorf("port %d is already in use", port)
		}
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackProvisionServer) ValidateDelete() error {
	openstackprovisionserverlog.Info("validate delete", "name", r.Name)

	return checkBackupOperationBlocksAction(r.Namespace, APIActionDelete)
}

//+kubebuilder:webhook:path=/mutate-osp-director-openstack-org-v1beta1-openstackprovisionserver,mutating=true,failurePolicy=fail,sideEffects=None,groups=osp-director.openstack.org,resources=openstackprovisionservers,verbs=create;update,versions=v1beta1,name=mopenstackprovisionserver.kb.io,admissionReviewVersions={v1,v1beta1}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OpenStackProvisionServer) Default() {
	openstackephemeralheatlog.Info("default", "name", r.Name)
}
