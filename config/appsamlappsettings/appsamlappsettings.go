package appsamlappsettings

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_saml_app_settings", func(r *config.Resource) {
		r.References["app_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/app/v1alpha1.SAML",
			Extractor: "github.com/healthcarecom/provider-okta/apis/app/v1alpha1.SamlID()",
		}
	})
}
