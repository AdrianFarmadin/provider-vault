/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/AdrianFarmadin/provider-vault/config/approleauth"
	"github.com/AdrianFarmadin/provider-vault/config/auth"
	"github.com/AdrianFarmadin/provider-vault/config/awssecret"
	"github.com/AdrianFarmadin/provider-vault/config/identity"
	"github.com/AdrianFarmadin/provider-vault/config/jwtauth"
	"github.com/AdrianFarmadin/provider-vault/config/kubernetesauth"
	"github.com/AdrianFarmadin/provider-vault/config/kv"
	"github.com/AdrianFarmadin/provider-vault/config/vault"
)

const (
	resourcePrefix = "vault"
	modulePath     = "github.com/AdrianFarmadin/provider-vault"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		approleauth.Configure,
		auth.Configure,
		awssecret.Configure,
		identity.Configure,
		jwtauth.Configure,
		kubernetesauth.Configure,
		kv.Configure,
		vault.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
