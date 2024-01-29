package authserverclaimdefault

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_auth_server_claim_default", func(r *config.Resource) {
		r.References["auth_server_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.Server",
			Extractor: "github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.AuthServerID()",
		}
	})
}
