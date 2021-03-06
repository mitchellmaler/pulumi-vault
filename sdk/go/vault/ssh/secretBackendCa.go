// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ssh

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage CA information in an SSH secret backend
// [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ssh_secret_backend_ca.html.markdown.
type SecretBackendCa struct {
	pulumi.CustomResourceState

	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrOutput `pulumi:"generateSigningKey"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewSecretBackendCa registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCa(ctx *pulumi.Context,
	name string, args *SecretBackendCaArgs, opts ...pulumi.ResourceOption) (*SecretBackendCa, error) {
	if args == nil {
		args = &SecretBackendCaArgs{}
	}
	var resource SecretBackendCa
	err := ctx.RegisterResource("vault:ssh/secretBackendCa:SecretBackendCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendCa gets an existing SecretBackendCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendCaState, opts ...pulumi.ResourceOption) (*SecretBackendCa, error) {
	var resource SecretBackendCa
	err := ctx.ReadResource("vault:ssh/secretBackendCa:SecretBackendCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendCa resources.
type secretBackendCaState struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend *string `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey *bool `pulumi:"generateSigningKey"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey *string `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey *string `pulumi:"publicKey"`
}

type SecretBackendCaState struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrInput
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrInput
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringPtrInput
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringPtrInput
}

func (SecretBackendCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCaState)(nil)).Elem()
}

type secretBackendCaArgs struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend *string `pulumi:"backend"`
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey *bool `pulumi:"generateSigningKey"`
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey *string `pulumi:"privateKey"`
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey *string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SecretBackendCa resource.
type SecretBackendCaArgs struct {
	// The path where the SSH secret backend is mounted. Defaults to 'ssh'
	Backend pulumi.StringPtrInput
	// Whether Vault should generate the signing key pair internally. Defaults to true
	GenerateSigningKey pulumi.BoolPtrInput
	// The private key part the SSH CA key pair; required if generateSigningKey is false.
	PrivateKey pulumi.StringPtrInput
	// The public key part the SSH CA key pair; required if generateSigningKey is false.
	PublicKey pulumi.StringPtrInput
}

func (SecretBackendCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCaArgs)(nil)).Elem()
}

