/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	authbackendconfig "github.com/AdrianFarmadin/provider-vault/internal/controller/kubernetesauth/authbackendconfig"
authbackendrole "github.com/AdrianFarmadin/provider-vault/internal/controller/kubernetesauth/authbackendrole"
providerconfig "github.com/AdrianFarmadin/provider-vault/internal/controller/providerconfig"

)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authbackendconfig.Setup,
		authbackendrole.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
