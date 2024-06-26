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

package v1beta1

import (
	"github.com/openstack-k8s-operators/osp-director-operator/api/shared"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenStackProvisionServerSpec defines the desired state of OpenStackProvisionServer
type OpenStackProvisionServerSpec struct {
	// The port on which the Apache server should listen
	Port int `json:"port"`
	// An optional interface to use instead of the cluster's default provisioning interface (if any)
	Interface string `json:"interface,omitempty"`
	// URL for RHEL qcow2 image (compressed as gz, or uncompressed)
	BaseImageURL string `json:"baseImageUrl"`
	// Container image URL for init container that downloads the RHEL qcow2 image (baseImageUrl)
	DownloaderImageURL string `json:"downloaderImageUrl,omitempty"`
	// Container image URL for the main container that serves the downloaded RHEL qcow2 image (baseImageUrl)
	ApacheImageURL string `json:"apacheImageUrl,omitempty"`
	// Container image URL for the sidecar container that discovers provisioning network IPs
	AgentImageURL string `json:"agentImageUrl,omitempty"`
}

// OpenStackProvisionServerStatus defines the observed state of OpenStackProvisionServer
type OpenStackProvisionServerStatus struct {
	// Surfaces status in GUI
	Conditions shared.ConditionList `json:"conditions,omitempty" optional:"true"`
	// Holds provisioning status for this provision server
	ProvisioningStatus OpenStackProvisionServerProvisioningStatus `json:"provisioningStatus,omitempty"`
	// IP of the provisioning interface on the node running the ProvisionServer pod
	ProvisionIP string `json:"provisionIp,omitempty"`
	// URL of provisioning image on underlying Apache web server
	LocalImageURL string `json:"localImageUrl,omitempty"`
}

// OpenStackProvisionServerProvisioningStatus represents the overall provisioning state of all BaremetalHosts in
// the OpenStackProvisionServer (with an optional explanatory message)
type OpenStackProvisionServerProvisioningStatus struct {
	State  shared.ProvisioningState `json:"state,omitempty"`
	Reason string                   `json:"reason,omitempty"`
}

// IsReady - Is this resource in its fully-configured (quiesced) state?
func (instance *OpenStackProvisionServer) IsReady() bool {
	return instance.Status.ProvisioningStatus.State == shared.ProvisioningState(shared.ProvisionServerCondTypeProvisioned)
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=osprovserver;osprovservers
// +operator-sdk:csv:customresourcedefinitions:displayName="OpenStack ProvisionServer"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.provisioningStatus.state",description="Status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.provisioningStatus.reason",description="Reason"

// OpenStackProvisionServer used to serve custom images for baremetal provisioning with Metal3
type OpenStackProvisionServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenStackProvisionServerSpec   `json:"spec,omitempty"`
	Status OpenStackProvisionServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpenStackProvisionServerList contains a list of OpenStackProvisionServer
type OpenStackProvisionServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenStackProvisionServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenStackProvisionServer{}, &OpenStackProvisionServerList{})
}
