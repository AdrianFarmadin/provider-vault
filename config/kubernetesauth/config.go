package kubernetesauth

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for kubernetesauth group.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_kubernetes_auth_backend_config", func(r *config.Resource) {
        r.ShortGroup = "kubernetesauth"
        r.References["backend"] = config.Reference{
			Type: "github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend",
		}
    })
    p.AddResourceConfigurator("vault_kubernetes_auth_backend_role", func(r *config.Resource) {
        r.ShortGroup = "kubernetesauth"
        r.References["backend"] = config.Reference{
			Type: "github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend",
		}
    })
}