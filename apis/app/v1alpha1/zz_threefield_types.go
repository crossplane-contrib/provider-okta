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

type ThreeFieldObservation struct {

	// Custom error page URL
	AccessibilityErrorRedirectURL *string `json:"accessibilityErrorRedirectUrl,omitempty" tf:"accessibility_error_redirect_url,omitempty"`

	// Custom login page URL
	AccessibilityLoginRedirectURL *string `json:"accessibilityLoginRedirectUrl,omitempty" tf:"accessibility_login_redirect_url,omitempty"`

	// Enable self service
	AccessibilitySelfService *bool `json:"accessibilitySelfService,omitempty" tf:"accessibility_self_service,omitempty"`

	// Application notes for admins.
	AdminNote *string `json:"adminNote,omitempty" tf:"admin_note,omitempty"`

	// Displays specific appLinks for the app
	AppLinksJSON *string `json:"appLinksJson,omitempty" tf:"app_links_json,omitempty"`

	// Display auto submit toolbar
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty" tf:"auto_submit_toolbar,omitempty"`

	// Login button field CSS selector
	ButtonSelector *string `json:"buttonSelector,omitempty" tf:"button_selector,omitempty"`

	// Application credentials scheme
	CredentialsScheme *string `json:"credentialsScheme,omitempty" tf:"credentials_scheme,omitempty"`

	// Application notes for end users.
	EnduserNote *string `json:"enduserNote,omitempty" tf:"enduser_note,omitempty"`

	// Extra field CSS selector
	ExtraFieldSelector *string `json:"extraFieldSelector,omitempty" tf:"extra_field_selector,omitempty"`

	// Value for extra form field
	ExtraFieldValue *string `json:"extraFieldValue,omitempty" tf:"extra_field_value,omitempty"`

	// Do not display application icon on mobile app
	HideIos *bool `json:"hideIos,omitempty" tf:"hide_ios,omitempty"`

	// Do not display application icon to users
	HideWeb *bool `json:"hideWeb,omitempty" tf:"hide_web,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Pretty name of app.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Local path to logo of the application.
	Logo *string `json:"logo,omitempty" tf:"logo,omitempty"`

	// URL of the application's logo
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// Name of the app.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Login password field CSS selector
	PasswordSelector *string `json:"passwordSelector,omitempty" tf:"password_selector,omitempty"`

	// Allow user to reveal password
	RevealPassword *bool `json:"revealPassword,omitempty" tf:"reveal_password,omitempty"`

	// Shared password, required for certain schemes.
	SharedPassword *string `json:"sharedPassword,omitempty" tf:"shared_password,omitempty"`

	// Shared username, required for certain schemes.
	SharedUsername *string `json:"sharedUsername,omitempty" tf:"shared_username,omitempty"`

	// Sign on mode of application.
	SignOnMode *string `json:"signOnMode,omitempty" tf:"sign_on_mode,omitempty"`

	// Status of application.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Login URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// A regex that further restricts URL to the specified regex
	URLRegex *string `json:"urlRegex,omitempty" tf:"url_regex,omitempty"`

	// Username template
	UserNameTemplate *string `json:"userNameTemplate,omitempty" tf:"user_name_template,omitempty"`

	// Push username on update
	UserNameTemplatePushStatus *string `json:"userNameTemplatePushStatus,omitempty" tf:"user_name_template_push_status,omitempty"`

	// Username template suffix
	UserNameTemplateSuffix *string `json:"userNameTemplateSuffix,omitempty" tf:"user_name_template_suffix,omitempty"`

	// Username template type
	UserNameTemplateType *string `json:"userNameTemplateType,omitempty" tf:"user_name_template_type,omitempty"`

	// Login username field CSS selector
	UsernameSelector *string `json:"usernameSelector,omitempty" tf:"username_selector,omitempty"`
}

type ThreeFieldParameters struct {

	// Custom error page URL
	// +kubebuilder:validation:Optional
	AccessibilityErrorRedirectURL *string `json:"accessibilityErrorRedirectUrl,omitempty" tf:"accessibility_error_redirect_url,omitempty"`

	// Custom login page URL
	// +kubebuilder:validation:Optional
	AccessibilityLoginRedirectURL *string `json:"accessibilityLoginRedirectUrl,omitempty" tf:"accessibility_login_redirect_url,omitempty"`

	// Enable self service
	// +kubebuilder:validation:Optional
	AccessibilitySelfService *bool `json:"accessibilitySelfService,omitempty" tf:"accessibility_self_service,omitempty"`

	// Application notes for admins.
	// +kubebuilder:validation:Optional
	AdminNote *string `json:"adminNote,omitempty" tf:"admin_note,omitempty"`

	// Displays specific appLinks for the app
	// +kubebuilder:validation:Optional
	AppLinksJSON *string `json:"appLinksJson,omitempty" tf:"app_links_json,omitempty"`

	// Display auto submit toolbar
	// +kubebuilder:validation:Optional
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty" tf:"auto_submit_toolbar,omitempty"`

	// Login button field CSS selector
	// +kubebuilder:validation:Optional
	ButtonSelector *string `json:"buttonSelector,omitempty" tf:"button_selector,omitempty"`

	// Application credentials scheme
	// +kubebuilder:validation:Optional
	CredentialsScheme *string `json:"credentialsScheme,omitempty" tf:"credentials_scheme,omitempty"`

	// Application notes for end users.
	// +kubebuilder:validation:Optional
	EnduserNote *string `json:"enduserNote,omitempty" tf:"enduser_note,omitempty"`

	// Extra field CSS selector
	// +kubebuilder:validation:Optional
	ExtraFieldSelector *string `json:"extraFieldSelector,omitempty" tf:"extra_field_selector,omitempty"`

	// Value for extra form field
	// +kubebuilder:validation:Optional
	ExtraFieldValue *string `json:"extraFieldValue,omitempty" tf:"extra_field_value,omitempty"`

	// Do not display application icon on mobile app
	// +kubebuilder:validation:Optional
	HideIos *bool `json:"hideIos,omitempty" tf:"hide_ios,omitempty"`

	// Do not display application icon to users
	// +kubebuilder:validation:Optional
	HideWeb *bool `json:"hideWeb,omitempty" tf:"hide_web,omitempty"`

	// Pretty name of app.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Local path to logo of the application.
	// +kubebuilder:validation:Optional
	Logo *string `json:"logo,omitempty" tf:"logo,omitempty"`

	// Login password field CSS selector
	// +kubebuilder:validation:Optional
	PasswordSelector *string `json:"passwordSelector,omitempty" tf:"password_selector,omitempty"`

	// Allow user to reveal password
	// +kubebuilder:validation:Optional
	RevealPassword *bool `json:"revealPassword,omitempty" tf:"reveal_password,omitempty"`

	// Shared password, required for certain schemes.
	// +kubebuilder:validation:Optional
	SharedPassword *string `json:"sharedPassword,omitempty" tf:"shared_password,omitempty"`

	// Shared username, required for certain schemes.
	// +kubebuilder:validation:Optional
	SharedUsername *string `json:"sharedUsername,omitempty" tf:"shared_username,omitempty"`

	// Status of application.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Login URL
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// A regex that further restricts URL to the specified regex
	// +kubebuilder:validation:Optional
	URLRegex *string `json:"urlRegex,omitempty" tf:"url_regex,omitempty"`

	// Username template
	// +kubebuilder:validation:Optional
	UserNameTemplate *string `json:"userNameTemplate,omitempty" tf:"user_name_template,omitempty"`

	// Push username on update
	// +kubebuilder:validation:Optional
	UserNameTemplatePushStatus *string `json:"userNameTemplatePushStatus,omitempty" tf:"user_name_template_push_status,omitempty"`

	// Username template suffix
	// +kubebuilder:validation:Optional
	UserNameTemplateSuffix *string `json:"userNameTemplateSuffix,omitempty" tf:"user_name_template_suffix,omitempty"`

	// Username template type
	// +kubebuilder:validation:Optional
	UserNameTemplateType *string `json:"userNameTemplateType,omitempty" tf:"user_name_template_type,omitempty"`

	// Login username field CSS selector
	// +kubebuilder:validation:Optional
	UsernameSelector *string `json:"usernameSelector,omitempty" tf:"username_selector,omitempty"`
}

// ThreeFieldSpec defines the desired state of ThreeField
type ThreeFieldSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThreeFieldParameters `json:"forProvider"`
}

// ThreeFieldStatus defines the observed state of ThreeField.
type ThreeFieldStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThreeFieldObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ThreeField is the Schema for the ThreeFields API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type ThreeField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.buttonSelector)",message="buttonSelector is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.extraFieldSelector)",message="extraFieldSelector is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.extraFieldValue)",message="extraFieldValue is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.label)",message="label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.passwordSelector)",message="passwordSelector is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.url)",message="url is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.usernameSelector)",message="usernameSelector is a required parameter"
	Spec   ThreeFieldSpec   `json:"spec"`
	Status ThreeFieldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThreeFieldList contains a list of ThreeFields
type ThreeFieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThreeField `json:"items"`
}

// Repository type metadata.
var (
	ThreeField_Kind             = "ThreeField"
	ThreeField_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThreeField_Kind}.String()
	ThreeField_KindAPIVersion   = ThreeField_Kind + "." + CRDGroupVersion.String()
	ThreeField_GroupVersionKind = CRDGroupVersion.WithKind(ThreeField_Kind)
)

func init() {
	SchemeBuilder.Register(&ThreeField{}, &ThreeFieldList{})
}
