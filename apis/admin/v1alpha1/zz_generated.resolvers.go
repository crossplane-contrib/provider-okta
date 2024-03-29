/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/healthcarecom/provider-okta/apis/user/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this RoleTargets.
func (mg *RoleTargets) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserID),
		Extract:      v1alpha1.UserID(),
		Reference:    mg.Spec.ForProvider.UserIDRef,
		Selector:     mg.Spec.ForProvider.UserIDSelector,
		To: reference.To{
			List:    &v1alpha1.UserList{},
			Managed: &v1alpha1.User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserID")
	}
	mg.Spec.ForProvider.UserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserIDRef = rsp.ResolvedReference

	return nil
}
