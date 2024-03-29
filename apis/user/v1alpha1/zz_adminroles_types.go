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

type AdminRolesObservation struct {

	// The list of Okta user admin roles, e.g. ["APP_ADMIN", "USER_ADMIN"] See API Docs.
	// User Okta admin roles - ie. ['APP_ADMIN', 'USER_ADMIN']
	AdminRoles []*string `json:"adminRoles,omitempty" tf:"admin_roles,omitempty"`

	// When this setting is enabled, the admins won't receive any of the default Okta
	// administrator emails. These admins also won't have access to contact Okta Support and open support cases on behalf of your org.
	// When this setting is enabled, the admins won't receive any of the default Okta administrator emails
	DisableNotifications *bool `json:"disableNotifications,omitempty" tf:"disable_notifications,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Okta user ID.
	// ID of a Okta User
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type AdminRolesParameters struct {

	// The list of Okta user admin roles, e.g. ["APP_ADMIN", "USER_ADMIN"] See API Docs.
	// User Okta admin roles - ie. ['APP_ADMIN', 'USER_ADMIN']
	// +kubebuilder:validation:Optional
	AdminRoles []*string `json:"adminRoles,omitempty" tf:"admin_roles,omitempty"`

	// When this setting is enabled, the admins won't receive any of the default Okta
	// administrator emails. These admins also won't have access to contact Okta Support and open support cases on behalf of your org.
	// When this setting is enabled, the admins won't receive any of the default Okta administrator emails
	// +kubebuilder:validation:Optional
	DisableNotifications *bool `json:"disableNotifications,omitempty" tf:"disable_notifications,omitempty"`

	// Okta user ID.
	// ID of a Okta User
	// +crossplane:generate:reference:type=github.com/healthcarecom/provider-okta/apis/user/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/healthcarecom/provider-okta/apis/user/v1alpha1.UserID()
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// AdminRolesSpec defines the desired state of AdminRoles
type AdminRolesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AdminRolesParameters `json:"forProvider"`
}

// AdminRolesStatus defines the observed state of AdminRoles.
type AdminRolesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AdminRolesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AdminRoles is the Schema for the AdminRoless API. Resource to manage a set of admin roles for a specific user.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type AdminRoles struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.adminRoles)",message="adminRoles is a required parameter"
	Spec   AdminRolesSpec   `json:"spec"`
	Status AdminRolesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdminRolesList contains a list of AdminRoless
type AdminRolesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdminRoles `json:"items"`
}

// Repository type metadata.
var (
	AdminRoles_Kind             = "AdminRoles"
	AdminRoles_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AdminRoles_Kind}.String()
	AdminRoles_KindAPIVersion   = AdminRoles_Kind + "." + CRDGroupVersion.String()
	AdminRoles_GroupVersionKind = CRDGroupVersion.WithKind(AdminRoles_Kind)
)

func init() {
	SchemeBuilder.Register(&AdminRoles{}, &AdminRolesList{})
}
