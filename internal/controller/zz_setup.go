/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	group "github.com/healthcarecom/provider-okta/internal/controller/group/group"
	memberships "github.com/healthcarecom/provider-okta/internal/controller/groupmembership/memberships"
	providerconfig "github.com/healthcarecom/provider-okta/internal/controller/providerconfig"
	user "github.com/healthcarecom/provider-okta/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
