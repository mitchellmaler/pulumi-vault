// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ssh

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage roles in an SSH secret backend
// [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ssh_secret_backend_role.html.markdown.
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrOutput `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrOutput `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrOutput `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrOutput `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrOutput `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrOutput `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrOutput `pulumi:"allowedDomains"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrOutput `pulumi:"allowedExtensions"`
	// Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
	AllowedUserKeyLengths pulumi.MapOutput `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrOutput `pulumi:"allowedUsers"`
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrOutput `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapOutput `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapOutput `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrOutput `pulumi:"defaultUser"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrOutput `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringOutput `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Time To Live value.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.KeyType == nil {
		return nil, errors.New("missing required argument 'KeyType'")
	}
	if args == nil {
		args = &SecretBackendRoleArgs{}
	}
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:ssh/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:ssh/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates *bool `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates *bool `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds *bool `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions *string `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains *string `pulumi:"allowedDomains"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions *string `pulumi:"allowedExtensions"`
	// Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
	AllowedUserKeyLengths map[string]interface{} `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers *string `pulumi:"allowedUsers"`
	// The path where the SSH secret backend is mounted.
	Backend *string `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList *string `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions map[string]interface{} `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions map[string]interface{} `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser *string `pulumi:"defaultUser"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat *string `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType *string `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl *string `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name *string `pulumi:"name"`
	// Specifies the Time To Live value.
	Ttl *string `pulumi:"ttl"`
}

type SecretBackendRoleState struct {
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrInput
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrInput
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrInput
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrInput
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrInput
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrInput
	// Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
	AllowedUserKeyLengths pulumi.MapInput
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrInput
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringPtrInput
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrInput
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapInput
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapInput
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrInput
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrInput
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringPtrInput
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringPtrInput
	// Specifies the name of the role to create.
	Name pulumi.StringPtrInput
	// Specifies the Time To Live value.
	Ttl pulumi.StringPtrInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates *bool `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates *bool `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds *bool `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions *string `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains *string `pulumi:"allowedDomains"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions *string `pulumi:"allowedExtensions"`
	// Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
	AllowedUserKeyLengths map[string]interface{} `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers *string `pulumi:"allowedUsers"`
	// The path where the SSH secret backend is mounted.
	Backend string `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList *string `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions map[string]interface{} `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions map[string]interface{} `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser *string `pulumi:"defaultUser"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat *string `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType string `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl *string `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name *string `pulumi:"name"`
	// Specifies the Time To Live value.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrInput
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrInput
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrInput
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrInput
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrInput
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrInput
	// Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
	AllowedUserKeyLengths pulumi.MapInput
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrInput
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringInput
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrInput
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapInput
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapInput
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrInput
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrInput
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringInput
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringPtrInput
	// Specifies the name of the role to create.
	Name pulumi.StringPtrInput
	// Specifies the Time To Live value.
	Ttl pulumi.StringPtrInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}

