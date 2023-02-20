package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	SelfPackagePath = "github.com/AdrianFarmadin/provider-vault/config/common"
	PathRoleNameExtractor = SelfPackagePath + ".RoleNameExtractor()"

)

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

