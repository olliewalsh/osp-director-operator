//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	k8s_cni_cncf_iov1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/openstack-k8s-operators/osp-director-operator/api/shared"
	"k8s.io/apimachinery/pkg/runtime"
	"kubevirt.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrsForBackup) DeepCopyInto(out *CrsForBackup) {
	*out = *in
	in.OpenStackBaremetalSets.DeepCopyInto(&out.OpenStackBaremetalSets)
	in.OpenStackClients.DeepCopyInto(&out.OpenStackClients)
	in.OpenStackControlPlanes.DeepCopyInto(&out.OpenStackControlPlanes)
	in.OpenStackMACAddresses.DeepCopyInto(&out.OpenStackMACAddresses)
	in.OpenStackNets.DeepCopyInto(&out.OpenStackNets)
	in.OpenStackNetAttachments.DeepCopyInto(&out.OpenStackNetAttachments)
	in.OpenStackNetConfigs.DeepCopyInto(&out.OpenStackNetConfigs)
	in.OpenStackProvisionServers.DeepCopyInto(&out.OpenStackProvisionServers)
	in.OpenStackVMSets.DeepCopyInto(&out.OpenStackVMSets)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrsForBackup.
func (in *CrsForBackup) DeepCopy() *CrsForBackup {
	if in == nil {
		return nil
	}
	out := new(CrsForBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hash) DeepCopyInto(out *Hash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hash.
func (in *Hash) DeepCopy() *Hash {
	if in == nil {
		return nil
	}
	out := new(Hash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NAD != nil {
		in, out := &in.NAD, &out.NAD
		*out = make(map[string]k8s_cni_cncf_iov1.NetworkAttachmentDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	in.IPStatus.DeepCopyInto(&out.IPStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPStatus) DeepCopyInto(out *IPStatus) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPStatus.
func (in *IPStatus) DeepCopy() *IPStatus {
	if in == nil {
		return nil
	}
	out := new(IPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBackup) DeepCopyInto(out *OpenStackBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBackup.
func (in *OpenStackBackup) DeepCopy() *OpenStackBackup {
	if in == nil {
		return nil
	}
	out := new(OpenStackBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBackupList) DeepCopyInto(out *OpenStackBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBackupList.
func (in *OpenStackBackupList) DeepCopy() *OpenStackBackupList {
	if in == nil {
		return nil
	}
	out := new(OpenStackBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBackupSpec) DeepCopyInto(out *OpenStackBackupSpec) {
	*out = *in
	in.Crs.DeepCopyInto(&out.Crs)
	in.ConfigMaps.DeepCopyInto(&out.ConfigMaps)
	in.Secrets.DeepCopyInto(&out.Secrets)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBackupSpec.
func (in *OpenStackBackupSpec) DeepCopy() *OpenStackBackupSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackBackupStatus) DeepCopyInto(out *OpenStackBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackBackupStatus.
func (in *OpenStackBackupStatus) DeepCopy() *OpenStackBackupStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlane) DeepCopyInto(out *OpenStackControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlane.
func (in *OpenStackControlPlane) DeepCopy() *OpenStackControlPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneList) DeepCopyInto(out *OpenStackControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneList.
func (in *OpenStackControlPlaneList) DeepCopy() *OpenStackControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneProvisioningStatus) DeepCopyInto(out *OpenStackControlPlaneProvisioningStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneProvisioningStatus.
func (in *OpenStackControlPlaneProvisioningStatus) DeepCopy() *OpenStackControlPlaneProvisioningStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneProvisioningStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneSpec) DeepCopyInto(out *OpenStackControlPlaneSpec) {
	*out = *in
	if in.VirtualMachineRoles != nil {
		in, out := &in.VirtualMachineRoles, &out.VirtualMachineRoles
		*out = make(map[string]OpenStackVirtualMachineRoleSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.OpenStackClientNetworks != nil {
		in, out := &in.OpenStackClientNetworks, &out.OpenStackClientNetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalServiceVIPs != nil {
		in, out := &in.AdditionalServiceVIPs, &out.AdditionalServiceVIPs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneSpec.
func (in *OpenStackControlPlaneSpec) DeepCopy() *OpenStackControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneStatus) DeepCopyInto(out *OpenStackControlPlaneStatus) {
	*out = *in
	if in.VIPStatus != nil {
		in, out := &in.VIPStatus, &out.VIPStatus
		*out = make(map[string]HostStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(shared.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ProvisioningStatus = in.ProvisioningStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneStatus.
func (in *OpenStackControlPlaneStatus) DeepCopy() *OpenStackControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSet) DeepCopyInto(out *OpenStackVMSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSet.
func (in *OpenStackVMSet) DeepCopy() *OpenStackVMSet {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackVMSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSetDisk) DeepCopyInto(out *OpenStackVMSetDisk) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSetDisk.
func (in *OpenStackVMSetDisk) DeepCopy() *OpenStackVMSetDisk {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSetDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSetList) DeepCopyInto(out *OpenStackVMSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackVMSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSetList.
func (in *OpenStackVMSetList) DeepCopy() *OpenStackVMSetList {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackVMSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSetProvisioningStatus) DeepCopyInto(out *OpenStackVMSetProvisioningStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSetProvisioningStatus.
func (in *OpenStackVMSetProvisioningStatus) DeepCopy() *OpenStackVMSetProvisioningStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSetProvisioningStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSetSpec) DeepCopyInto(out *OpenStackVMSetSpec) {
	*out = *in
	out.RootDisk = in.RootDisk
	if in.AdditionalDisks != nil {
		in, out := &in.AdditionalDisks, &out.AdditionalDisks
		*out = make([]OpenStackVMSetDisk, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapDNS != nil {
		in, out := &in.BootstrapDNS, &out.BootstrapDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNSSearchDomains != nil {
		in, out := &in.DNSSearchDomains, &out.DNSSearchDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionStrategy != nil {
		in, out := &in.EvictionStrategy, &out.EvictionStrategy
		*out = new(v1.EvictionStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSetSpec.
func (in *OpenStackVMSetSpec) DeepCopy() *OpenStackVMSetSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVMSetStatus) DeepCopyInto(out *OpenStackVMSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(shared.ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ProvisioningStatus = in.ProvisioningStatus
	if in.VMpods != nil {
		in, out := &in.VMpods, &out.VMpods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VMHosts != nil {
		in, out := &in.VMHosts, &out.VMHosts
		*out = make(map[string]HostStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVMSetStatus.
func (in *OpenStackVMSetStatus) DeepCopy() *OpenStackVMSetStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackVMSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackVirtualMachineRoleSpec) DeepCopyInto(out *OpenStackVirtualMachineRoleSpec) {
	*out = *in
	out.RootDisk = in.RootDisk
	if in.AdditionalDisks != nil {
		in, out := &in.AdditionalDisks, &out.AdditionalDisks
		*out = make([]OpenStackVMSetDisk, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionStrategy != nil {
		in, out := &in.EvictionStrategy, &out.EvictionStrategy
		*out = new(v1.EvictionStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackVirtualMachineRoleSpec.
func (in *OpenStackVirtualMachineRoleSpec) DeepCopy() *OpenStackVirtualMachineRoleSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackVirtualMachineRoleSpec)
	in.DeepCopyInto(out)
	return out
}
