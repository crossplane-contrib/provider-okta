package appoauthapiscope

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_oauth_api_scope", func(r *config.Resource) {
	})
}
