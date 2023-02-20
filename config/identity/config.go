package identity

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vault_identity_group", func(r *config.Resource) {})

    p.AddResourceConfigurator("vault_identity_group_alias", func(r *config.Resource) {
        r.References["canonical_id"] = config.Reference{
			Type: "Group",
		}
    })
}