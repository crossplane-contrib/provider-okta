package appsecurepasswordstore

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_secure_password_store", func(r *config.Resource) {
	})
}
