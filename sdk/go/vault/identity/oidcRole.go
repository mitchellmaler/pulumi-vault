// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_oidc_role.html.markdown.
type OidcRole struct {
	s *pulumi.ResourceState
}

// NewOidcRole registers a new resource with the given unique name, arguments, and options.
func NewOidcRole(ctx *pulumi.Context,
	name string, args *OidcRoleArgs, opts ...pulumi.ResourceOpt) (*OidcRole, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["key"] = nil
		inputs["name"] = nil
		inputs["template"] = nil
		inputs["ttl"] = nil
	} else {
		inputs["key"] = args.Key
		inputs["name"] = args.Name
		inputs["template"] = args.Template
		inputs["ttl"] = args.Ttl
	}
	inputs["clientId"] = nil
	s, err := ctx.RegisterResource("vault:identity/oidcRole:OidcRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OidcRole{s: s}, nil
}

// GetOidcRole gets an existing OidcRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OidcRoleState, opts ...pulumi.ResourceOpt) (*OidcRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clientId"] = state.ClientId
		inputs["key"] = state.Key
		inputs["name"] = state.Name
		inputs["template"] = state.Template
		inputs["ttl"] = state.Ttl
	}
	s, err := ctx.ReadResource("vault:identity/oidcRole:OidcRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OidcRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OidcRole) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OidcRole) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The value that will be included in the `aud` field of all the OIDC identity
// tokens issued by this role
func (r *OidcRole) ClientId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clientId"])
}

// A configured named key, the key must already exist
// before tokens can be issued.
func (r *OidcRole) Key() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["key"])
}

// Name of the OIDC Role to create.
func (r *OidcRole) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The template string to use for generating tokens. This may be in
// string-ified JSON or base64 format. See the
// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
// for the template format.
func (r *OidcRole) Template() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["template"])
}

// TTL of the tokens generated against the role in number of seconds.
func (r *OidcRole) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// Input properties used for looking up and filtering OidcRole resources.
type OidcRoleState struct {
	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId interface{}
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key interface{}
	// Name of the OIDC Role to create.
	Name interface{}
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template interface{}
	// TTL of the tokens generated against the role in number of seconds.
	Ttl interface{}
}

// The set of arguments for constructing a OidcRole resource.
type OidcRoleArgs struct {
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key interface{}
	// Name of the OIDC Role to create.
	Name interface{}
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template interface{}
	// TTL of the tokens generated against the role in number of seconds.
	Ttl interface{}
}
