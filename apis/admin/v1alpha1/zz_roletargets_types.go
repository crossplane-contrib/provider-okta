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

type RoleTargetsObservation struct {

	// List of app names (name represents set of app instances) or a combination of app name and app instance ID (like 'salesforce' or 'facebook.0oapsqQ6dv19pqyEo0g3')
	Apps []*string `json:"apps,omitempty" tf:"apps,omitempty"`

	// List of group IDs
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of a role
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// Type of the role that is assigned to the user and supports optional targets
	RoleType *string `json:"roleType,omitempty" tf:"role_type,omitempty"`

	// User associated with the role
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type RoleTargetsParameters struct {

	// List of app names (name represents set of app instances) or a combination of app name and app instance ID (like 'salesforce' or 'facebook.0oapsqQ6dv19pqyEo0g3')
	// +kubebuilder:validation:Optional
	Apps []*string `json:"apps,omitempty" tf:"apps,omitempty"`

	// List of group IDs
	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// Type of the role that is assigned to the user and supports optional targets
	// +kubebuilder:validation:Optional
	RoleType *string `json:"roleType,omitempty" tf:"role_type,omitempty"`

	// User associated with the role
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

// RoleTargetsSpec defines the desired state of RoleTargets
type RoleTargetsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleTargetsParameters `json:"forProvider"`
}

// RoleTargetsStatus defines the observed state of RoleTargets.
type RoleTargetsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleTargetsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleTargets is the Schema for the RoleTargetss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type RoleTargets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roleType)",message="roleType is a required parameter"
	Spec   RoleTargetsSpec   `json:"spec"`
	Status RoleTargetsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleTargetsList contains a list of RoleTargetss
type RoleTargetsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleTargets `json:"items"`
}

// Repository type metadata.
var (
	RoleTargets_Kind             = "RoleTargets"
	RoleTargets_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleTargets_Kind}.String()
	RoleTargets_KindAPIVersion   = RoleTargets_Kind + "." + CRDGroupVersion.String()
	RoleTargets_GroupVersionKind = CRDGroupVersion.WithKind(RoleTargets_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleTargets{}, &RoleTargetsList{})
}
