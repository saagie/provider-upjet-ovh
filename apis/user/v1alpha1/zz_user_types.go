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

type RolesObservation struct {

	// A description associated with the user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// id of the role
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// name of the role
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// list of permissions associated with the role
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type RolesParameters struct {
}

type UserObservation struct {

	// the date the user was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// id of the role
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of roles associated with the user.
	Roles []RolesObservation `json:"roles,omitempty" tf:"roles,omitempty"`

	// the status of the user. should be normally set to 'ok'.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// the username generated for the user. This username can be used with
	// the Openstack API.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// A description associated with the user.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// a convenient map representing an openstack_rc file.
	// Note: no password nor sensitive token is set in this map.
	// +kubebuilder:validation:Optional
	OpenstackRc map[string]*string `json:"openstackRc,omitempty" tf:"openstack_rc,omitempty"`

	// The name of a role. See role_names.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// A list of role names. Values can be:
	// +kubebuilder:validation:Optional
	RoleNames []*string `json:"roleNames,omitempty" tf:"role_names,omitempty"`

	// The id of the public cloud project. If omitted,
	// the OVH_CLOUD_PROJECT_SERVICE environment variable is used.
	// Service name of the resource representing the id of the cloud project.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Creates a user in a public cloud project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}