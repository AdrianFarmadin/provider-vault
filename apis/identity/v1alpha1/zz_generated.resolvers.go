/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this GroupAlias.
func (mg *GroupAlias) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CanonicalID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CanonicalIDRef,
		Selector:     mg.Spec.ForProvider.CanonicalIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CanonicalID")
	}
	mg.Spec.ForProvider.CanonicalID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CanonicalIDRef = rsp.ResolvedReference

	return nil
}
