// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_oidc_key.html.markdown.
type OidcKey struct {
	s *pulumi.ResourceState
}

// NewOidcKey registers a new resource with the given unique name, arguments, and options.
func NewOidcKey(ctx *pulumi.Context,
	name string, args *OidcKeyArgs, opts ...pulumi.ResourceOpt) (*OidcKey, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["algorithm"] = nil
		inputs["allowedClientIds"] = nil
		inputs["name"] = nil
		inputs["rotationPeriod"] = nil
		inputs["verificationTtl"] = nil
	} else {
		inputs["algorithm"] = args.Algorithm
		inputs["allowedClientIds"] = args.AllowedClientIds
		inputs["name"] = args.Name
		inputs["rotationPeriod"] = args.RotationPeriod
		inputs["verificationTtl"] = args.VerificationTtl
	}
	s, err := ctx.RegisterResource("vault:identity/oidcKey:OidcKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OidcKey{s: s}, nil
}

// GetOidcKey gets an existing OidcKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OidcKeyState, opts ...pulumi.ResourceOpt) (*OidcKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["algorithm"] = state.Algorithm
		inputs["allowedClientIds"] = state.AllowedClientIds
		inputs["name"] = state.Name
		inputs["rotationPeriod"] = state.RotationPeriod
		inputs["verificationTtl"] = state.VerificationTtl
	}
	s, err := ctx.ReadResource("vault:identity/oidcKey:OidcKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OidcKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OidcKey) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OidcKey) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Signing algorithm to use. Signing algorithm to use.
// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
func (r *OidcKey) Algorithm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["algorithm"])
}

// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
// allowed.
func (r *OidcKey) AllowedClientIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["allowedClientIds"])
}

// Name of the OIDC Key to create.
func (r *OidcKey) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// How often to generate a new signing key in number of seconds
func (r *OidcKey) RotationPeriod() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["rotationPeriod"])
}

// "Controls how long the public portion of a signing key will be
// available for verification after being rotated in seconds.
func (r *OidcKey) VerificationTtl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["verificationTtl"])
}

// Input properties used for looking up and filtering OidcKey resources.
type OidcKeyState struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm interface{}
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds interface{}
	// Name of the OIDC Key to create.
	Name interface{}
	// How often to generate a new signing key in number of seconds
	RotationPeriod interface{}
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl interface{}
}

// The set of arguments for constructing a OidcKey resource.
type OidcKeyArgs struct {
	// Signing algorithm to use. Signing algorithm to use.
	// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	Algorithm interface{}
	// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
	// allowed.
	AllowedClientIds interface{}
	// Name of the OIDC Key to create.
	Name interface{}
	// How often to generate a new signing key in number of seconds
	RotationPeriod interface{}
	// "Controls how long the public portion of a signing key will be
	// available for verification after being rotated in seconds.
	VerificationTtl interface{}
}
