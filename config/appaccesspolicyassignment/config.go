package appaccesspolicyassignment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_access_policy_assignment", func(r *config.Resource) {
	})
}
