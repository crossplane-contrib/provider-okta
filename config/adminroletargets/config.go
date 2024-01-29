package adminroletargets

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_admin_role_targets", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.User",
			Extractor: "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.UserID()",
		}
	})
}
