/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/healthcarecom/provider-okta/config/appgroupassignment"
	"github.com/healthcarecom/provider-okta/config/group"
	"github.com/healthcarecom/provider-okta/config/groupmemberships"
	"github.com/healthcarecom/provider-okta/config/user"
)

const (
	resourcePrefix = "okta"
	modulePath     = "github.com/healthcarecom/provider-okta"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		group.Configure,
		user.Configure,
		groupmemberships.Configure,
		appgroupassignment.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
