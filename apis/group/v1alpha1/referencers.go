package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// GroupID returns an extractor that returns the Group ID.
func GroupID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		u, ok := mg.(*Group)
		if !ok {
			return ""
		}
		return u.GetID()
	}
}
