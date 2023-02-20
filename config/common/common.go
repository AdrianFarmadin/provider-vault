package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/AdrianFarmadin/provider-vault/config/common"

	// PathRoleNameExtractor is the golang path to RoleName extractor
	// function in this package.
	PathRoleNameExtractor = SelfPackagePath + ".RoleNameExtractor()"

)

// RoleNameExtractor extracts role name of the resources from "spec.forProvider.roleName"
func RoleNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("spec.forProvider.roleName")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}

