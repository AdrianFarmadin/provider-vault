package auth

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for auth group.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_auth_backend", func(r *config.Resource) {
        r.ShortGroup = "auth"
    })
}