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

// ResolveReferences of this IpRestrictions.
func (mg *IpRestrictions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KubeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KubeIDRef,
		Selector:     mg.Spec.ForProvider.KubeIDSelector,
		To: reference.To{
			List:    &KubeList{},
			Managed: &Kube{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KubeID")
	}
	mg.Spec.ForProvider.KubeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KubeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodePool.
func (mg *NodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KubeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KubeIDRef,
		Selector:     mg.Spec.ForProvider.KubeIDSelector,
		To: reference.To{
			List:    &KubeList{},
			Managed: &Kube{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KubeID")
	}
	mg.Spec.ForProvider.KubeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KubeIDRef = rsp.ResolvedReference

	return nil
}
