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

type AdmissionpluginsObservation struct {
}

type AdmissionpluginsParameters struct {

	// +kubebuilder:validation:Optional
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ApiserverObservation struct {
}

type ApiserverParameters struct {

	// +kubebuilder:validation:Optional
	Admissionplugins []AdmissionpluginsParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type CustomizationObservation struct {
}

type CustomizationParameters struct {

	// +kubebuilder:validation:Optional
	Apiserver []ApiserverParameters `json:"apiserver,omitempty" tf:"apiserver,omitempty"`
}

type KubeObservation struct {

	// True if control-plane is up to date.
	ControlPlaneIsUpToDate *bool `json:"controlPlaneIsUpToDate,omitempty" tf:"control_plane_is_up_to_date,omitempty"`

	// Managed Kubernetes Service ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// True if all nodes and control-plane are up to date.
	IsUpToDate *bool `json:"isUpToDate,omitempty" tf:"is_up_to_date,omitempty"`

	// Kubernetes versions available for upgrade.
	NextUpgradeVersions []*string `json:"nextUpgradeVersions,omitempty" tf:"next_upgrade_versions,omitempty"`

	// Cluster nodes URL.
	NodesURL *string `json:"nodesUrl,omitempty" tf:"nodes_url,omitempty"`

	// Cluster status. Should be normally set to 'READY'.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Management URL of your cluster.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type KubeParameters struct {

	// Customer customization object
	// +kubebuilder:validation:Optional
	Customization []CustomizationParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// The name of the kubernetes cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The private network configuration
	// +kubebuilder:validation:Optional
	PrivateNetworkConfiguration []PrivateNetworkConfigurationParameters `json:"privateNetworkConfiguration,omitempty" tf:"private_network_configuration,omitempty"`

	// OpenStack private network ID to use.
	// Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// a valid OVHcloud public cloud region ID in which the kubernetes
	// cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
	// Changing this value recreates the resource.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The id of the public cloud project. If omitted,
	// the OVH_CLOUD_PROJECT_SERVICE environment variable is used.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	// +kubebuilder:validation:Optional
	UpdatePolicy *string `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	// kubernetes version to use.
	// Changing this value updates the resource. Defaults to latest available.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PrivateNetworkConfigurationObservation struct {
}

type PrivateNetworkConfigurationParameters struct {

	// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
	// +kubebuilder:validation:Required
	DefaultVrackGateway *string `json:"defaultVrackGateway" tf:"default_vrack_gateway,omitempty"`

	// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
	// +kubebuilder:validation:Required
	PrivateNetworkRoutingAsDefault *bool `json:"privateNetworkRoutingAsDefault" tf:"private_network_routing_as_default,omitempty"`
}

// KubeSpec defines the desired state of Kube
type KubeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubeParameters `json:"forProvider"`
}

// KubeStatus defines the observed state of Kube.
type KubeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kube is the Schema for the Kubes API. Creates a kubernetes managed cluster in a public cloud project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type Kube struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeSpec   `json:"spec"`
	Status            KubeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeList contains a list of Kubes
type KubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kube `json:"items"`
}

// Repository type metadata.
var (
	Kube_Kind             = "Kube"
	Kube_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Kube_Kind}.String()
	Kube_KindAPIVersion   = Kube_Kind + "." + CRDGroupVersion.String()
	Kube_GroupVersionKind = CRDGroupVersion.WithKind(Kube_Kind)
)

func init() {
	SchemeBuilder.Register(&Kube{}, &KubeList{})
}
