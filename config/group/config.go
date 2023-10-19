package group

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_group", func(r *config.Resource) {
		r.ShortGroup = "group"
	})
}
