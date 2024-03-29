/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Role.
func (mg *Role) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      GroupID(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.GroupAssignments),
		Extract:       GroupID(),
		References:    mg.Spec.ForProvider.GroupAssignmentsRefs,
		Selector:      mg.Spec.ForProvider.GroupAssignmentsSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupAssignments")
	}
	mg.Spec.ForProvider.GroupAssignments = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupAssignmentsRefs = mrsp.ResolvedReferences

	return nil
}
