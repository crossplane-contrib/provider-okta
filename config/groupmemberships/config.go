package groupmemberships

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_group_memberships", func(r *config.Resource) {
		r.ShortGroup = "groupMembership"
		r.References["users"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.User",
			Extractor: "github.com/healthcarecom/provider-okta/apis/user/v1alpha1.UserID()",
		}
		r.References["group_id"] = config.Reference{
			Type:      "github.com/healthcarecom/provider-okta/apis/group/v1alpha1.Group",
			Extractor: "github.com/healthcarecom/provider-okta/apis/group/v1alpha1.GroupID()",
		}
	})
}
