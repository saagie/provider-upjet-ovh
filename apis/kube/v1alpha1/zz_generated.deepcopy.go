//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionpluginsObservation) DeepCopyInto(out *AdmissionpluginsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionpluginsObservation.
func (in *AdmissionpluginsObservation) DeepCopy() *AdmissionpluginsObservation {
	if in == nil {
		return nil
	}
	out := new(AdmissionpluginsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionpluginsParameters) DeepCopyInto(out *AdmissionpluginsParameters) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionpluginsParameters.
func (in *AdmissionpluginsParameters) DeepCopy() *AdmissionpluginsParameters {
	if in == nil {
		return nil
	}
	out := new(AdmissionpluginsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiserverObservation) DeepCopyInto(out *ApiserverObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiserverObservation.
func (in *ApiserverObservation) DeepCopy() *ApiserverObservation {
	if in == nil {
		return nil
	}
	out := new(ApiserverObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiserverParameters) DeepCopyInto(out *ApiserverParameters) {
	*out = *in
	if in.Admissionplugins != nil {
		in, out := &in.Admissionplugins, &out.Admissionplugins
		*out = make([]AdmissionpluginsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiserverParameters.
func (in *ApiserverParameters) DeepCopy() *ApiserverParameters {
	if in == nil {
		return nil
	}
	out := new(ApiserverParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationObservation) DeepCopyInto(out *CustomizationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationObservation.
func (in *CustomizationObservation) DeepCopy() *CustomizationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationParameters) DeepCopyInto(out *CustomizationParameters) {
	*out = *in
	if in.Apiserver != nil {
		in, out := &in.Apiserver, &out.Apiserver
		*out = make([]ApiserverParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationParameters.
func (in *CustomizationParameters) DeepCopy() *CustomizationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictions) DeepCopyInto(out *IpRestrictions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictions.
func (in *IpRestrictions) DeepCopy() *IpRestrictions {
	if in == nil {
		return nil
	}
	out := new(IpRestrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpRestrictions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictionsList) DeepCopyInto(out *IpRestrictionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IpRestrictions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictionsList.
func (in *IpRestrictionsList) DeepCopy() *IpRestrictionsList {
	if in == nil {
		return nil
	}
	out := new(IpRestrictionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpRestrictionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictionsObservation) DeepCopyInto(out *IpRestrictionsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictionsObservation.
func (in *IpRestrictionsObservation) DeepCopy() *IpRestrictionsObservation {
	if in == nil {
		return nil
	}
	out := new(IpRestrictionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictionsParameters) DeepCopyInto(out *IpRestrictionsParameters) {
	*out = *in
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.KubeID != nil {
		in, out := &in.KubeID, &out.KubeID
		*out = new(string)
		**out = **in
	}
	if in.KubeIDRef != nil {
		in, out := &in.KubeIDRef, &out.KubeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeIDSelector != nil {
		in, out := &in.KubeIDSelector, &out.KubeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictionsParameters.
func (in *IpRestrictionsParameters) DeepCopy() *IpRestrictionsParameters {
	if in == nil {
		return nil
	}
	out := new(IpRestrictionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictionsSpec) DeepCopyInto(out *IpRestrictionsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictionsSpec.
func (in *IpRestrictionsSpec) DeepCopy() *IpRestrictionsSpec {
	if in == nil {
		return nil
	}
	out := new(IpRestrictionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRestrictionsStatus) DeepCopyInto(out *IpRestrictionsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRestrictionsStatus.
func (in *IpRestrictionsStatus) DeepCopy() *IpRestrictionsStatus {
	if in == nil {
		return nil
	}
	out := new(IpRestrictionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kube) DeepCopyInto(out *Kube) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kube.
func (in *Kube) DeepCopy() *Kube {
	if in == nil {
		return nil
	}
	out := new(Kube)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kube) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeList) DeepCopyInto(out *KubeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kube, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeList.
func (in *KubeList) DeepCopy() *KubeList {
	if in == nil {
		return nil
	}
	out := new(KubeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeObservation) DeepCopyInto(out *KubeObservation) {
	*out = *in
	if in.ControlPlaneIsUpToDate != nil {
		in, out := &in.ControlPlaneIsUpToDate, &out.ControlPlaneIsUpToDate
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsUpToDate != nil {
		in, out := &in.IsUpToDate, &out.IsUpToDate
		*out = new(bool)
		**out = **in
	}
	if in.NextUpgradeVersions != nil {
		in, out := &in.NextUpgradeVersions, &out.NextUpgradeVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NodesURL != nil {
		in, out := &in.NodesURL, &out.NodesURL
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeObservation.
func (in *KubeObservation) DeepCopy() *KubeObservation {
	if in == nil {
		return nil
	}
	out := new(KubeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeParameters) DeepCopyInto(out *KubeParameters) {
	*out = *in
	if in.Customization != nil {
		in, out := &in.Customization, &out.Customization
		*out = make([]CustomizationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworkConfiguration != nil {
		in, out := &in.PrivateNetworkConfiguration, &out.PrivateNetworkConfiguration
		*out = make([]PrivateNetworkConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateNetworkID != nil {
		in, out := &in.PrivateNetworkID, &out.PrivateNetworkID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.UpdatePolicy != nil {
		in, out := &in.UpdatePolicy, &out.UpdatePolicy
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeParameters.
func (in *KubeParameters) DeepCopy() *KubeParameters {
	if in == nil {
		return nil
	}
	out := new(KubeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSpec) DeepCopyInto(out *KubeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSpec.
func (in *KubeSpec) DeepCopy() *KubeSpec {
	if in == nil {
		return nil
	}
	out := new(KubeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStatus) DeepCopyInto(out *KubeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStatus.
func (in *KubeStatus) DeepCopy() *KubeStatus {
	if in == nil {
		return nil
	}
	out := new(KubeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataObservation) DeepCopyInto(out *MetadataObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataObservation.
func (in *MetadataObservation) DeepCopy() *MetadataObservation {
	if in == nil {
		return nil
	}
	out := new(MetadataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataParameters) DeepCopyInto(out *MetadataParameters) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataParameters.
func (in *MetadataParameters) DeepCopy() *MetadataParameters {
	if in == nil {
		return nil
	}
	out := new(MetadataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolObservation) DeepCopyInto(out *NodePoolObservation) {
	*out = *in
	if in.AvailableNodes != nil {
		in, out := &in.AvailableNodes, &out.AvailableNodes
		*out = new(float64)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CurrentNodes != nil {
		in, out := &in.CurrentNodes, &out.CurrentNodes
		*out = new(float64)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.SizeStatus != nil {
		in, out := &in.SizeStatus, &out.SizeStatus
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UpToDateNodes != nil {
		in, out := &in.UpToDateNodes, &out.UpToDateNodes
		*out = new(float64)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolObservation.
func (in *NodePoolObservation) DeepCopy() *NodePoolObservation {
	if in == nil {
		return nil
	}
	out := new(NodePoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolParameters) DeepCopyInto(out *NodePoolParameters) {
	*out = *in
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(bool)
		**out = **in
	}
	if in.DesiredNodes != nil {
		in, out := &in.DesiredNodes, &out.DesiredNodes
		*out = new(float64)
		**out = **in
	}
	if in.FlavorName != nil {
		in, out := &in.FlavorName, &out.FlavorName
		*out = new(string)
		**out = **in
	}
	if in.KubeID != nil {
		in, out := &in.KubeID, &out.KubeID
		*out = new(string)
		**out = **in
	}
	if in.KubeIDRef != nil {
		in, out := &in.KubeIDRef, &out.KubeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeIDSelector != nil {
		in, out := &in.KubeIDSelector, &out.KubeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxNodes != nil {
		in, out := &in.MaxNodes, &out.MaxNodes
		*out = new(float64)
		**out = **in
	}
	if in.MinNodes != nil {
		in, out := &in.MinNodes, &out.MinNodes
		*out = new(float64)
		**out = **in
	}
	if in.MonthlyBilled != nil {
		in, out := &in.MonthlyBilled, &out.MonthlyBilled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]TemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolParameters.
func (in *NodePoolParameters) DeepCopy() *NodePoolParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolStatus) DeepCopyInto(out *NodePoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolStatus.
func (in *NodePoolStatus) DeepCopy() *NodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(NodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkConfigurationObservation) DeepCopyInto(out *PrivateNetworkConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkConfigurationObservation.
func (in *PrivateNetworkConfigurationObservation) DeepCopy() *PrivateNetworkConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateNetworkConfigurationParameters) DeepCopyInto(out *PrivateNetworkConfigurationParameters) {
	*out = *in
	if in.DefaultVrackGateway != nil {
		in, out := &in.DefaultVrackGateway, &out.DefaultVrackGateway
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworkRoutingAsDefault != nil {
		in, out := &in.PrivateNetworkRoutingAsDefault, &out.PrivateNetworkRoutingAsDefault
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateNetworkConfigurationParameters.
func (in *PrivateNetworkConfigurationParameters) DeepCopy() *PrivateNetworkConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateNetworkConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecObservation) DeepCopyInto(out *SpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecObservation.
func (in *SpecObservation) DeepCopy() *SpecObservation {
	if in == nil {
		return nil
	}
	out := new(SpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecParameters) DeepCopyInto(out *SpecParameters) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]map[string]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]*string, len(*in))
				for key, val := range *in {
					var outVal *string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(string)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
		}
	}
	if in.Unschedulable != nil {
		in, out := &in.Unschedulable, &out.Unschedulable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecParameters.
func (in *SpecParameters) DeepCopy() *SpecParameters {
	if in == nil {
		return nil
	}
	out := new(SpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateObservation) DeepCopyInto(out *TemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateObservation.
func (in *TemplateObservation) DeepCopy() *TemplateObservation {
	if in == nil {
		return nil
	}
	out := new(TemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateParameters) DeepCopyInto(out *TemplateParameters) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]MetadataParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]SpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateParameters.
func (in *TemplateParameters) DeepCopy() *TemplateParameters {
	if in == nil {
		return nil
	}
	out := new(TemplateParameters)
	in.DeepCopyInto(out)
	return out
}
