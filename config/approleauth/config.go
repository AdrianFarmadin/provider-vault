package approleauth

import (
    "github.com/upbound/upjet/pkg/config"
    "github.com/AdrianFarmadin/provider-vault/config/common"
)

// Configure adds configurations for approleauth group.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_approle_auth_backend_role", func(r *config.Resource) {
        r.ShortGroup = "approleauth"
        r.References["backend"] = config.Reference{
			Type: "github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend",
		}
    })

    p.AddResourceConfigurator("vault_approle_auth_backend_role_secret_id", func(r *config.Resource) {
        r.ShortGroup = "approleauth"
        r.References["backend"] = config.Reference{
			Type: "github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend",
		}
        r.References["role_name"] = config.Reference{
			Type: "AuthBackendRole",
			Extractor: common.PathRoleNameExtractor,
		}
    })
}