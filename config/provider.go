/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/healthcarecom/provider-okta/config/appbasicauth"
	"github.com/healthcarecom/provider-okta/config/appbookmark"
	"github.com/healthcarecom/provider-okta/config/appoauth"
	"github.com/healthcarecom/provider-okta/config/appoauthapiscope"
	"github.com/healthcarecom/provider-okta/config/appoauthroleassignment"
	"github.com/healthcarecom/provider-okta/config/appsaml"
	"github.com/healthcarecom/provider-okta/config/appsamlappsettings"
	"github.com/healthcarecom/provider-okta/config/appsecurepasswordstore"
	"github.com/healthcarecom/provider-okta/config/appsharedcredentials"
	"github.com/healthcarecom/provider-okta/config/appswa"
	"github.com/healthcarecom/provider-okta/config/appthreefield"
	"github.com/healthcarecom/provider-okta/config/authserver"
	"github.com/healthcarecom/provider-okta/config/authserverclaim"
	"github.com/healthcarecom/provider-okta/config/authserverclaimdefault"
	"github.com/healthcarecom/provider-okta/config/authserverdefault"
	"github.com/healthcarecom/provider-okta/config/authserverpolicy"
	"github.com/healthcarecom/provider-okta/config/authserverpolicyrule"
	"github.com/healthcarecom/provider-okta/config/authserverscope"
	"github.com/healthcarecom/provider-okta/config/behavior"
	"github.com/healthcarecom/provider-okta/config/grouprole"
	"github.com/healthcarecom/provider-okta/config/grouprule"
	"github.com/healthcarecom/provider-okta/config/idpoidc"
	"github.com/healthcarecom/provider-okta/config/idpsaml"
	"github.com/healthcarecom/provider-okta/config/useradminroles"
	"github.com/healthcarecom/provider-okta/config/userbaseschemaproperty"
	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/healthcarecom/provider-okta/config/adminroletargets"
	"github.com/healthcarecom/provider-okta/config/appaccesspolicyassignment"
	"github.com/healthcarecom/provider-okta/config/appautologin"
	"github.com/healthcarecom/provider-okta/config/appgroupassignment"
	"github.com/healthcarecom/provider-okta/config/group"
	"github.com/healthcarecom/provider-okta/config/groupmemberships"
	"github.com/healthcarecom/provider-okta/config/trustedorigin"
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
		adminroletargets.Configure,
		appaccesspolicyassignment.Configure,
		appautologin.Configure,
		appbasicauth.Configure,
		appbookmark.Configure,
		appoauth.Configure,
		appoauthapiscope.Configure,
		appoauthroleassignment.Configure,
		appsaml.Configure,
		appsamlappsettings.Configure,
		appsecurepasswordstore.Configure,
		appsharedcredentials.Configure,
		appswa.Configure,
		appthreefield.Configure,
		authserver.Configure,
		authserverclaim.Configure,
		authserverclaimdefault.Configure,
		authserverdefault.Configure,
		authserverpolicy.Configure,
		authserverpolicyrule.Configure,
		authserverscope.Configure,
		behavior.Configure,
		grouprole.Configure,
		grouprule.Configure,
		idpoidc.Configure,
		idpsaml.Configure,
		useradminroles.Configure,
		userbaseschemaproperty.Configure,
		trustedorigin.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
