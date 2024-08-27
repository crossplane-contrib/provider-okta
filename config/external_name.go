/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"okta_group":                        config.IdentifierFromProvider,
	"okta_group_memberships":            config.IdentifierFromProvider,
	"okta_user":                         config.IdentifierFromProvider,
	"okta_app_group_assignment":         config.IdentifierFromProvider,
	"okta_admin_role_targets":           config.IdentifierFromProvider,
	"okta_app_access_policy_assignment": config.IdentifierFromProvider,
	"okta_app_auto_login":               config.IdentifierFromProvider,
	"okta_app_basic_auth":               config.IdentifierFromProvider,
	"okta_app_bookmark":                 config.IdentifierFromProvider,
	"okta_app_oauth":                    config.IdentifierFromProvider,
	"okta_app_oauth_api_scope":          config.IdentifierFromProvider,
	"okta_app_oauth_role_assignment":    config.IdentifierFromProvider,
	"okta_app_saml":                     config.IdentifierFromProvider,
	"okta_app_saml_app_settings":        config.IdentifierFromProvider,
	"okta_app_secure_password_store":    config.IdentifierFromProvider,
	"okta_app_shared_credentials":       config.IdentifierFromProvider,
	"okta_app_swa":                      config.IdentifierFromProvider,
	"okta_app_three_field":              config.IdentifierFromProvider,
	"okta_auth_server":                  config.IdentifierFromProvider,
	"okta_auth_server_claim":            config.IdentifierFromProvider,
	"okta_auth_server_claim_default":    config.IdentifierFromProvider,
	"okta_auth_server_default":          config.IdentifierFromProvider,
	"okta_auth_server_policy":           config.IdentifierFromProvider,
	"okta_auth_server_policy_rule":      config.IdentifierFromProvider,
	"okta_auth_server_scope":            config.IdentifierFromProvider,
	"okta_behavior":                     config.IdentifierFromProvider,
	"okta_group_role":                   config.IdentifierFromProvider,
	"okta_group_rule":                   config.IdentifierFromProvider,
	"okta_idp_oidc":                     config.IdentifierFromProvider,
	"okta_idp_saml":                     config.IdentifierFromProvider,
	"okta_user_admin_roles":             config.IdentifierFromProvider,
	"okta_user_base_schema_property":    config.IdentifierFromProvider,
	"okta_trusted_origin":               config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
