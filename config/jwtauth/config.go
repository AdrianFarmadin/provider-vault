package jwtauth

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for jwtauth group.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_jwt_auth_backend", func(r *config.Resource) {
        r.ShortGroup = "jwtauth"
    })

    p.AddResourceConfigurator("vault_jwt_auth_backend_role", func(r *config.Resource) {
        r.ShortGroup = "jwtauth"
        r.References["backend"] = config.Reference{
			Type: "AuthBackend",
		}
    })
}