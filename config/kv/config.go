package kv

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_kv_secret_v2", func(r *config.Resource) {})
}