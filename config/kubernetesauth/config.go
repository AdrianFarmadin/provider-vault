package kubernetesauth

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_kubernetes_auth_backend_config", func(r *config.Resource) {
        r.ShortGroup = "kubernetesauth"
    })
    p.AddResourceConfigurator("vault_kubernetes_auth_backend_role", func(r *config.Resource) {
        r.ShortGroup = "kubernetesauth"
    })
}