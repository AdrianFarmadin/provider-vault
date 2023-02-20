package vault

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_mount", func(r *config.Resource) {})

    p.AddResourceConfigurator("vault_policy", func(r *config.Resource) {})

    p.AddResourceConfigurator("vault_audit", func(r *config.Resource) {})

    p.AddResourceConfigurator("vault_token", func(r *config.Resource) {})
}