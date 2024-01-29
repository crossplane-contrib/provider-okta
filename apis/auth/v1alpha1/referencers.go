package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

// AuthServerID returns an extractor that returns the ID for a Server
func AuthServerID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := mg.(*Server)
		if !ok {
			return ""
		}
		return s.GetID()
	}
}

// AuthServerPolicyID returns an extractor that returns the ID for a ServerPolicy
func AuthServerPolicyID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		s, ok := mg.(*ServerPolicy)
		if !ok {
			return ""
		}
		return s.GetID()
	}
}
