package grouprole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_group_role", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.References["group_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/group/v1alpha1.Group",
			Extractor: "github.com/healthcarecom/provider-okta/apis/group/v1alpha1.GroupID()",
		}
	})
}
