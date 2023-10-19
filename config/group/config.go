package group

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_group", func(r *config.Resource) {
		r.ShortGroup = "group"
	})
}
