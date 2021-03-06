// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Reads role tag information from an AWS auth backend in Vault. 
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_role_tag.html.markdown.
type AuthBackendRoleTag struct {
	pulumi.CustomResourceState

	// If set, allows migration of the underlying instances where the client resides. Use with caution.
	AllowInstanceMigration pulumi.BoolPtrOutput `pulumi:"allowInstanceMigration"`
	// The path to the AWS auth backend to
	// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set, only allows a single token to be granted per instance ID.
	DisallowReauthentication pulumi.BoolPtrOutput `pulumi:"disallowReauthentication"`
	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// The maximum TTL of the tokens issued using this role.
	MaxTtl pulumi.StringPtrOutput `pulumi:"maxTtl"`
	// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The name of the AWS auth backend role to read
	// role tags from, with no leading or trailing `/`s.
	Role pulumi.StringOutput `pulumi:"role"`
	// The key of the role tag.
	TagKey pulumi.StringOutput `pulumi:"tagKey"`
	// The value to set the role key.
	TagValue pulumi.StringOutput `pulumi:"tagValue"`
}

// NewAuthBackendRoleTag registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoleTag(ctx *pulumi.Context,
	name string, args *AuthBackendRoleTagArgs, opts ...pulumi.ResourceOption) (*AuthBackendRoleTag, error) {
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &AuthBackendRoleTagArgs{}
	}
	var resource AuthBackendRoleTag
	err := ctx.RegisterResource("vault:aws/authBackendRoleTag:AuthBackendRoleTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRoleTag gets an existing AuthBackendRoleTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoleTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoleTagState, opts ...pulumi.ResourceOption) (*AuthBackendRoleTag, error) {
	var resource AuthBackendRoleTag
	err := ctx.ReadResource("vault:aws/authBackendRoleTag:AuthBackendRoleTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRoleTag resources.
type authBackendRoleTagState struct {
	// If set, allows migration of the underlying instances where the client resides. Use with caution.
	AllowInstanceMigration *bool `pulumi:"allowInstanceMigration"`
	// The path to the AWS auth backend to
	// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
	Backend *string `pulumi:"backend"`
	// If set, only allows a single token to be granted per instance ID.
	DisallowReauthentication *bool `pulumi:"disallowReauthentication"`
	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId *string `pulumi:"instanceId"`
	// The maximum TTL of the tokens issued using this role.
	MaxTtl *string `pulumi:"maxTtl"`
	// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
	Policies []string `pulumi:"policies"`
	// The name of the AWS auth backend role to read
	// role tags from, with no leading or trailing `/`s.
	Role *string `pulumi:"role"`
	// The key of the role tag.
	TagKey *string `pulumi:"tagKey"`
	// The value to set the role key.
	TagValue *string `pulumi:"tagValue"`
}

type AuthBackendRoleTagState struct {
	// If set, allows migration of the underlying instances where the client resides. Use with caution.
	AllowInstanceMigration pulumi.BoolPtrInput
	// The path to the AWS auth backend to
	// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
	Backend pulumi.StringPtrInput
	// If set, only allows a single token to be granted per instance ID.
	DisallowReauthentication pulumi.BoolPtrInput
	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId pulumi.StringPtrInput
	// The maximum TTL of the tokens issued using this role.
	MaxTtl pulumi.StringPtrInput
	// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
	Policies pulumi.StringArrayInput
	// The name of the AWS auth backend role to read
	// role tags from, with no leading or trailing `/`s.
	Role pulumi.StringPtrInput
	// The key of the role tag.
	TagKey pulumi.StringPtrInput
	// The value to set the role key.
	TagValue pulumi.StringPtrInput
}

func (AuthBackendRoleTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleTagState)(nil)).Elem()
}

type authBackendRoleTagArgs struct {
	// If set, allows migration of the underlying instances where the client resides. Use with caution.
	AllowInstanceMigration *bool `pulumi:"allowInstanceMigration"`
	// The path to the AWS auth backend to
	// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
	Backend *string `pulumi:"backend"`
	// If set, only allows a single token to be granted per instance ID.
	DisallowReauthentication *bool `pulumi:"disallowReauthentication"`
	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId *string `pulumi:"instanceId"`
	// The maximum TTL of the tokens issued using this role.
	MaxTtl *string `pulumi:"maxTtl"`
	// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
	Policies []string `pulumi:"policies"`
	// The name of the AWS auth backend role to read
	// role tags from, with no leading or trailing `/`s.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AuthBackendRoleTag resource.
type AuthBackendRoleTagArgs struct {
	// If set, allows migration of the underlying instances where the client resides. Use with caution.
	AllowInstanceMigration pulumi.BoolPtrInput
	// The path to the AWS auth backend to
	// read role tags from, with no leading or trailing `/`s. Defaults to "aws".
	Backend pulumi.StringPtrInput
	// If set, only allows a single token to be granted per instance ID.
	DisallowReauthentication pulumi.BoolPtrInput
	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId pulumi.StringPtrInput
	// The maximum TTL of the tokens issued using this role.
	MaxTtl pulumi.StringPtrInput
	// The policies to be associated with the tag. Must be a subset of the policies associated with the role.
	Policies pulumi.StringArrayInput
	// The name of the AWS auth backend role to read
	// role tags from, with no leading or trailing `/`s.
	Role pulumi.StringInput
}

func (AuthBackendRoleTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleTagArgs)(nil)).Elem()
}

