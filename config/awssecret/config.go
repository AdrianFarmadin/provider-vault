package awssecret

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_aws_secret_backend", func(r *config.Resource) {
        r.ShortGroup = "awssecret"
    })

    p.AddResourceConfigurator("vault_aws_secret_backend_role", func(r *config.Resource) {
        r.ShortGroup = "awssecret"
        r.References["backend"] = config.Reference{
			Type: "SecretBackend",
		}
    })
}