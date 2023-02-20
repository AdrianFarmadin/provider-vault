/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	authbackendrole "github.com/AdrianFarmadin/provider-vault/internal/controller/approleauth/authbackendrole"
	authbackendrolesecretid "github.com/AdrianFarmadin/provider-vault/internal/controller/approleauth/authbackendrolesecretid"
	backend "github.com/AdrianFarmadin/provider-vault/internal/controller/auth/backend"
	secretbackend "github.com/AdrianFarmadin/provider-vault/internal/controller/awssecret/secretbackend"
	secretbackendrole "github.com/AdrianFarmadin/provider-vault/internal/controller/awssecret/secretbackendrole"
	group "github.com/AdrianFarmadin/provider-vault/internal/controller/identity/group"
	groupalias "github.com/AdrianFarmadin/provider-vault/internal/controller/identity/groupalias"
	authbackend "github.com/AdrianFarmadin/provider-vault/internal/controller/jwtauth/authbackend"
	authbackendrolejwtauth "github.com/AdrianFarmadin/provider-vault/internal/controller/jwtauth/authbackendrole"
	authbackendconfig "github.com/AdrianFarmadin/provider-vault/internal/controller/kubernetesauth/authbackendconfig"
	authbackendrolekubernetesauth "github.com/AdrianFarmadin/provider-vault/internal/controller/kubernetesauth/authbackendrole"
	secretv2 "github.com/AdrianFarmadin/provider-vault/internal/controller/kv/secretv2"
	providerconfig "github.com/AdrianFarmadin/provider-vault/internal/controller/providerconfig"
	audit "github.com/AdrianFarmadin/provider-vault/internal/controller/vault/audit"
	mount "github.com/AdrianFarmadin/provider-vault/internal/controller/vault/mount"
	policy "github.com/AdrianFarmadin/provider-vault/internal/controller/vault/policy"
	token "github.com/AdrianFarmadin/provider-vault/internal/controller/vault/token"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authbackendrole.Setup,
		authbackendrolesecretid.Setup,
		backend.Setup,
		secretbackend.Setup,
		secretbackendrole.Setup,
		group.Setup,
		groupalias.Setup,
		authbackend.Setup,
		authbackendrolejwtauth.Setup,
		authbackendconfig.Setup,
		authbackendrolekubernetesauth.Setup,
		secretv2.Setup,
		providerconfig.Setup,
		audit.Setup,
		mount.Setup,
		policy.Setup,
		token.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
