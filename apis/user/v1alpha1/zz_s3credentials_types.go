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

type S3CredentialsObservation struct {

	// the Access Key ID
	AccessKeyID *string `json:"accessKeyId,omitempty" tf:"access_key_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of a public cloud project's user.
	InternalUserID *string `json:"internalUserId,omitempty" tf:"internal_user_id,omitempty"`
}

type S3CredentialsParameters struct {

	// The ID of the public cloud project. If omitted,
	// the OVH_CLOUD_PROJECT_SERVICE environment variable is used.
	// Service name of the resource representing the ID of the cloud project.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// The ID of a public cloud project's user.
	// The user ID
	// +crossplane:generate:reference:type=github.com/saagie/provider-ovh/apis/user/v1alpha1.User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// S3CredentialsSpec defines the desired state of S3Credentials
type S3CredentialsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3CredentialsParameters `json:"forProvider"`
}

// S3CredentialsStatus defines the observed state of S3Credentials.
type S3CredentialsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3CredentialsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3Credentials is the Schema for the S3Credentialss API. Creates an S3 Credential for a user in a public cloud project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type S3Credentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3CredentialsSpec   `json:"spec"`
	Status            S3CredentialsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3CredentialsList contains a list of S3Credentialss
type S3CredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Credentials `json:"items"`
}

// Repository type metadata.
var (
	S3Credentials_Kind             = "S3Credentials"
	S3Credentials_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3Credentials_Kind}.String()
	S3Credentials_KindAPIVersion   = S3Credentials_Kind + "." + CRDGroupVersion.String()
	S3Credentials_GroupVersionKind = CRDGroupVersion.WithKind(S3Credentials_Kind)
)

func init() {
	SchemeBuilder.Register(&S3Credentials{}, &S3CredentialsList{})
}
