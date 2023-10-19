/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/healthcarecom/provider-okta/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal okta credentials as JSON"
	keyOrgName              = "org_name"
	keyBaseURL              = "base_url"
	keyHTTPProxy            = "http_proxy"
	keyAccessToken          = "access_token"
	keyAPIToken             = "api_token"
	keyClientID             = "client_id"
	keyScopes               = "scopes"
	keyPrivateKey           = "private_key"
	keyPrivateKeyID         = "private_key_id"
	keyBackoff              = "backoff"
	keyMinWaitSeconds       = "min_wait_seconds"
	keyMaxWaitSeconds       = "max_wait_seconds"
	keyMaxRetries           = "max_retries"
	keyRequestTimeout       = "request_timeout"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		ps.Configuration = map[string]any{}
		// Set credentials in Terraform provider configuration.
		if v, ok := creds[keyOrgName]; ok {
			ps.Configuration[keyOrgName] = v
		}
		if v, ok := creds[keyBaseURL]; ok {
			ps.Configuration[keyBaseURL] = v
		}
		if v, ok := creds[keyHTTPProxy]; ok {
			ps.Configuration[keyHTTPProxy] = v
		}
		if v, ok := creds[keyAccessToken]; ok {
			ps.Configuration[keyAccessToken] = v
		}
		if v, ok := creds[keyAPIToken]; ok {
			ps.Configuration[keyAPIToken] = v
		}
		if v, ok := creds[keyClientID]; ok {
			ps.Configuration[keyClientID] = v
		}
		if v, ok := creds[keyScopes]; ok {
			ps.Configuration[keyScopes] = v
		}
		if v, ok := creds[keyPrivateKey]; ok {
			ps.Configuration[keyPrivateKey] = v
		}
		if v, ok := creds[keyPrivateKeyID]; ok {
			ps.Configuration[keyPrivateKeyID] = v
		}
		if v, ok := creds[keyBackoff]; ok {
			ps.Configuration[keyBackoff] = v
		}
		if v, ok := creds[keyMinWaitSeconds]; ok {
			ps.Configuration[keyMinWaitSeconds] = v
		}
		if v, ok := creds[keyMaxWaitSeconds]; ok {
			ps.Configuration[keyMaxWaitSeconds] = v
		}
		if v, ok := creds[keyMaxRetries]; ok {
			ps.Configuration[keyMaxRetries] = v
		}
		if v, ok := creds[keyRequestTimeout]; ok {
			ps.Configuration[keyRequestTimeout] = v
		}
		return ps, nil
	}
}
