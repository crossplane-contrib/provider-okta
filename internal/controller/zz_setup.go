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
	group "github.com/healthcarecom/provider-okta/internal/controller/group/group"
	memberships "github.com/healthcarecom/provider-okta/internal/controller/groupmembership/memberships"
	providerconfig "github.com/healthcarecom/provider-okta/internal/controller/providerconfig"
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
		group.Setup,
		memberships.Setup,
		providerconfig.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
