/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1"
	common "github.com/AdrianFarmadin/provider-vault/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AuthBackendRole.
func (mg *AuthBackendRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AuthBackendRoleSecretID.
func (mg *AuthBackendRoleSecretID) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Backend),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackendRef,
		Selector:     mg.Spec.ForProvider.BackendSelector,
		To: reference.To{
			List:    &v1alpha1.BackendList{},
			Managed: &v1alpha1.Backend{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Backend")
	}
	mg.Spec.ForProvider.Backend = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackendRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleName),
		Extract:      common.RoleNameExtractor(),
		Reference:    mg.Spec.ForProvider.RoleNameRef,
		Selector:     mg.Spec.ForProvider.RoleNameSelector,
		To: reference.To{
			List:    &AuthBackendRoleList{},
			Managed: &AuthBackendRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleName")
	}
	mg.Spec.ForProvider.RoleName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleNameRef = rsp.ResolvedReference

	return nil
}
