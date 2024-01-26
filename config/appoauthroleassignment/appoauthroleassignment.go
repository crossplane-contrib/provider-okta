package appoauthroleassignment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_oauth_role_assignment", func(r *config.Resource) {
		r.References["client_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/app/v1alpha1.Oauth",
			Extractor: "github.com/healthcarecom/provider-okta/apis/app/v1alpha1.OauthClientID()",
		}
	})
}
