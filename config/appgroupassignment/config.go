package appgroupassignment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_group_assignment", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			Type: "github.com/healthcarecom/provider-okta/apis/group/v1alpha1.Group",
		}
	})
}
