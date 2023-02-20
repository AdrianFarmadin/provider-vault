/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"vault_identity_group": config.IdentifierFromProvider,
	"vault_identity_group_alias": config.IdentifierFromProvider,
	"vault_auth_backend": config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_config": config.IdentifierFromProvider,
	"vault_kubernetes_auth_backend_role": config.IdentifierFromProvider,
	"vault_mount": config.IdentifierFromProvider,
	"vault_policy": config.IdentifierFromProvider,
	"vault_audit": config.IdentifierFromProvider,
	"vault_token": config.IdentifierFromProvider,
	"vault_jwt_auth_backend": config.IdentifierFromProvider,
	"vault_jwt_auth_backend_role": config.IdentifierFromProvider,
	"vault_aws_secret_backend": config.IdentifierFromProvider,
	"vault_aws_secret_backend_role": config.IdentifierFromProvider,
	"vault_kv_secret_v2": config.IdentifierFromProvider,
	"vault_approle_auth_backend_role": config.IdentifierFromProvider,
	"vault_approle_auth_backend_role_secret_id": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
