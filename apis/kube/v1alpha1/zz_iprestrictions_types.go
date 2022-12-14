/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IpRestrictionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IpRestrictionsParameters struct {

	// List of CIDR authorized to interact with the managed Kubernetes cluster.
	// List of IP restrictions for the cluster
	// +kubebuilder:validation:Required
	Ips []*string `json:"ips" tf:"ips,omitempty"`

	// The id of the managed Kubernetes cluster.
	// Kube ID
	// +crossplane:generate:reference:type=github.com/saagie/provider-ovh/apis/kube/v1alpha1.Kube
	// +kubebuilder:validation:Optional
	KubeID *string `json:"kubeId,omitempty" tf:"kube_id,omitempty"`

	// Reference to a Kube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDRef *v1.Reference `json:"kubeIdRef,omitempty" tf:"-"`

	// Selector for a Kube in kube to populate kubeId.
	// +kubebuilder:validation:Optional
	KubeIDSelector *v1.Selector `json:"kubeIdSelector,omitempty" tf:"-"`

	// The id of the public cloud project. If omitted,
	// the OVH_CLOUD_PROJECT_SERVICE environment variable is used.
	// Service name
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// IpRestrictionsSpec defines the desired state of IpRestrictions
type IpRestrictionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpRestrictionsParameters `json:"forProvider"`
}

// IpRestrictionsStatus defines the observed state of IpRestrictions.
type IpRestrictionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpRestrictionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpRestrictions is the Schema for the IpRestrictionss API. Apply IP restrictions to an OVHcloud Managed Kubernetes cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type IpRestrictions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpRestrictionsSpec   `json:"spec"`
	Status            IpRestrictionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpRestrictionsList contains a list of IpRestrictionss
type IpRestrictionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpRestrictions `json:"items"`
}

// Repository type metadata.
var (
	IpRestrictions_Kind             = "IpRestrictions"
	IpRestrictions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpRestrictions_Kind}.String()
	IpRestrictions_KindAPIVersion   = IpRestrictions_Kind + "." + CRDGroupVersion.String()
	IpRestrictions_GroupVersionKind = CRDGroupVersion.WithKind(IpRestrictions_Kind)
)

func init() {
	SchemeBuilder.Register(&IpRestrictions{}, &IpRestrictionsList{})
}
