// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_crl_config.html.markdown.
type SecretBackendCrlConfig struct {
	s *pulumi.ResourceState
}

// NewSecretBackendCrlConfig registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCrlConfig(ctx *pulumi.Context,
	name string, args *SecretBackendCrlConfigArgs, opts ...pulumi.ResourceOpt) (*SecretBackendCrlConfig, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["disable"] = nil
		inputs["expiry"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["disable"] = args.Disable
		inputs["expiry"] = args.Expiry
	}
	s, err := ctx.RegisterResource("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendCrlConfig{s: s}, nil
}

// GetSecretBackendCrlConfig gets an existing SecretBackendCrlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCrlConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretBackendCrlConfigState, opts ...pulumi.ResourceOpt) (*SecretBackendCrlConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["disable"] = state.Disable
		inputs["expiry"] = state.Expiry
	}
	s, err := ctx.ReadResource("vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendCrlConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretBackendCrlConfig) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretBackendCrlConfig) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
func (r *SecretBackendCrlConfig) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// Disables or enables CRL building.
func (r *SecretBackendCrlConfig) Disable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disable"])
}

// Specifies the time until expiration.
func (r *SecretBackendCrlConfig) Expiry() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expiry"])
}

// Input properties used for looking up and filtering SecretBackendCrlConfig resources.
type SecretBackendCrlConfigState struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// Disables or enables CRL building.
	Disable interface{}
	// Specifies the time until expiration.
	Expiry interface{}
}

// The set of arguments for constructing a SecretBackendCrlConfig resource.
type SecretBackendCrlConfigArgs struct {
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// Disables or enables CRL building.
	Disable interface{}
	// Specifies the time until expiration.
	Expiry interface{}
}
