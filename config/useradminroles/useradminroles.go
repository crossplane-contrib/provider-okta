package useradminroles

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_user_admin_roles", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.References["user_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.User",
			Extractor: "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.UserID()",
		}
	})
}
