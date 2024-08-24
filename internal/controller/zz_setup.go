/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	roletargets "github.com/healthcarecom/provider-okta/internal/controller/admin/roletargets"
	accesspolicyassignment "github.com/healthcarecom/provider-okta/internal/controller/app/accesspolicyassignment"
	autologin "github.com/healthcarecom/provider-okta/internal/controller/app/autologin"
	basicauth "github.com/healthcarecom/provider-okta/internal/controller/app/basicauth"
	bookmark "github.com/healthcarecom/provider-okta/internal/controller/app/bookmark"
	groupassignment "github.com/healthcarecom/provider-okta/internal/controller/app/groupassignment"
	oauth "github.com/healthcarecom/provider-okta/internal/controller/app/oauth"
	oauthapiscope "github.com/healthcarecom/provider-okta/internal/controller/app/oauthapiscope"
	oauthroleassignment "github.com/healthcarecom/provider-okta/internal/controller/app/oauthroleassignment"
	saml "github.com/healthcarecom/provider-okta/internal/controller/app/saml"
	samlappsettings "github.com/healthcarecom/provider-okta/internal/controller/app/samlappsettings"
	securepasswordstore "github.com/healthcarecom/provider-okta/internal/controller/app/securepasswordstore"
	sharedcredentials "github.com/healthcarecom/provider-okta/internal/controller/app/sharedcredentials"
	swa "github.com/healthcarecom/provider-okta/internal/controller/app/swa"
	threefield "github.com/healthcarecom/provider-okta/internal/controller/app/threefield"
	server "github.com/healthcarecom/provider-okta/internal/controller/auth/server"
	serverclaim "github.com/healthcarecom/provider-okta/internal/controller/auth/serverclaim"
	serverclaimdefault "github.com/healthcarecom/provider-okta/internal/controller/auth/serverclaimdefault"
	serverdefault "github.com/healthcarecom/provider-okta/internal/controller/auth/serverdefault"
	serverpolicy "github.com/healthcarecom/provider-okta/internal/controller/auth/serverpolicy"
	serverpolicyrule "github.com/healthcarecom/provider-okta/internal/controller/auth/serverpolicyrule"
	serverscope "github.com/healthcarecom/provider-okta/internal/controller/auth/serverscope"
	group "github.com/healthcarecom/provider-okta/internal/controller/group/group"
	role "github.com/healthcarecom/provider-okta/internal/controller/group/role"
	rule "github.com/healthcarecom/provider-okta/internal/controller/group/rule"
	memberships "github.com/healthcarecom/provider-okta/internal/controller/groupmembership/memberships"
	oidc "github.com/healthcarecom/provider-okta/internal/controller/idp/oidc"
	samlidp "github.com/healthcarecom/provider-okta/internal/controller/idp/saml"
	behavior "github.com/healthcarecom/provider-okta/internal/controller/okta/behavior"
	providerconfig "github.com/healthcarecom/provider-okta/internal/controller/providerconfig"
	origin "github.com/healthcarecom/provider-okta/internal/controller/trustedorigin/origin"
	adminroles "github.com/healthcarecom/provider-okta/internal/controller/user/adminroles"
	baseschemaproperty "github.com/healthcarecom/provider-okta/internal/controller/user/baseschemaproperty"
	user "github.com/healthcarecom/provider-okta/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		roletargets.Setup,
		accesspolicyassignment.Setup,
		autologin.Setup,
		basicauth.Setup,
		bookmark.Setup,
		groupassignment.Setup,
		oauth.Setup,
		oauthapiscope.Setup,
		oauthroleassignment.Setup,
		saml.Setup,
		samlappsettings.Setup,
		securepasswordstore.Setup,
		sharedcredentials.Setup,
		swa.Setup,
		threefield.Setup,
		server.Setup,
		serverclaim.Setup,
		serverclaimdefault.Setup,
		serverdefault.Setup,
		serverpolicy.Setup,
		serverpolicyrule.Setup,
		serverscope.Setup,
		group.Setup,
		role.Setup,
		rule.Setup,
		memberships.Setup,
		oidc.Setup,
		samlidp.Setup,
		behavior.Setup,
		providerconfig.Setup,
		origin.Setup,
		adminroles.Setup,
		baseschemaproperty.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
