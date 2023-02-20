package kv

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for kv group.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_kv_secret_v2", func(r *config.Resource) {})
}