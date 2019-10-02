// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_intermediate_set_signed.html.markdown.
type BackendIntermediateSetSigned struct {
	s *pulumi.ResourceState
}

// NewBackendIntermediateSetSigned registers a new resource with the given unique name, arguments, and options.
func NewBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, args *BackendIntermediateSetSignedArgs, opts ...pulumi.ResourceOpt) (*BackendIntermediateSetSigned, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["certificate"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["certificate"] = args.Certificate
	}
	s, err := ctx.RegisterResource("vault:pkiSecret/backendIntermediateSetSigned:BackendIntermediateSetSigned", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendIntermediateSetSigned{s: s}, nil
}

// GetBackendIntermediateSetSigned gets an existing BackendIntermediateSetSigned resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BackendIntermediateSetSignedState, opts ...pulumi.ResourceOpt) (*BackendIntermediateSetSigned, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["certificate"] = state.Certificate
	}
	s, err := ctx.ReadResource("vault:pkiSecret/backendIntermediateSetSigned:BackendIntermediateSetSigned", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendIntermediateSetSigned{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BackendIntermediateSetSigned) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BackendIntermediateSetSigned) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The PKI secret backend the resource belongs to.
func (r *BackendIntermediateSetSigned) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// The certificate
func (r *BackendIntermediateSetSigned) Certificate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["certificate"])
}

// Input properties used for looking up and filtering BackendIntermediateSetSigned resources.
type BackendIntermediateSetSignedState struct {
	// The PKI secret backend the resource belongs to.
	Backend interface{}
	// The certificate
	Certificate interface{}
}

// The set of arguments for constructing a BackendIntermediateSetSigned resource.
type BackendIntermediateSetSignedArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend interface{}
	// The certificate
	Certificate interface{}
}
