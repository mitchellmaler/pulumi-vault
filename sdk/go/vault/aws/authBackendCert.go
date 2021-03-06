// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_cert.html.markdown.
type AuthBackendCert struct {
	pulumi.CustomResourceState

	// The  Base64 encoded AWS Public key required to
	// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
	// the [AWS
	// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
	AwsPublicCert pulumi.StringOutput `pulumi:"awsPublicCert"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The name of the certificate.
	CertName pulumi.StringOutput `pulumi:"certName"`
	// Either "pkcs7" or "identity", indicating the type of
	// document which can be verified using the given certificate. Defaults to
	// "pkcs7".
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAuthBackendCert registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendCert(ctx *pulumi.Context,
	name string, args *AuthBackendCertArgs, opts ...pulumi.ResourceOption) (*AuthBackendCert, error) {
	if args == nil || args.AwsPublicCert == nil {
		return nil, errors.New("missing required argument 'AwsPublicCert'")
	}
	if args == nil || args.CertName == nil {
		return nil, errors.New("missing required argument 'CertName'")
	}
	if args == nil {
		args = &AuthBackendCertArgs{}
	}
	var resource AuthBackendCert
	err := ctx.RegisterResource("vault:aws/authBackendCert:AuthBackendCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendCert gets an existing AuthBackendCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendCertState, opts ...pulumi.ResourceOption) (*AuthBackendCert, error) {
	var resource AuthBackendCert
	err := ctx.ReadResource("vault:aws/authBackendCert:AuthBackendCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendCert resources.
type authBackendCertState struct {
	// The  Base64 encoded AWS Public key required to
	// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
	// the [AWS
	// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
	AwsPublicCert *string `pulumi:"awsPublicCert"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// The name of the certificate.
	CertName *string `pulumi:"certName"`
	// Either "pkcs7" or "identity", indicating the type of
	// document which can be verified using the given certificate. Defaults to
	// "pkcs7".
	Type *string `pulumi:"type"`
}

type AuthBackendCertState struct {
	// The  Base64 encoded AWS Public key required to
	// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
	// the [AWS
	// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
	AwsPublicCert pulumi.StringPtrInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// The name of the certificate.
	CertName pulumi.StringPtrInput
	// Either "pkcs7" or "identity", indicating the type of
	// document which can be verified using the given certificate. Defaults to
	// "pkcs7".
	Type pulumi.StringPtrInput
}

func (AuthBackendCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendCertState)(nil)).Elem()
}

type authBackendCertArgs struct {
	// The  Base64 encoded AWS Public key required to
	// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
	// the [AWS
	// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
	AwsPublicCert string `pulumi:"awsPublicCert"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// The name of the certificate.
	CertName string `pulumi:"certName"`
	// Either "pkcs7" or "identity", indicating the type of
	// document which can be verified using the given certificate. Defaults to
	// "pkcs7".
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AuthBackendCert resource.
type AuthBackendCertArgs struct {
	// The  Base64 encoded AWS Public key required to
	// verify PKCS7 signature of the EC2 instance metadata. You can find this key in
	// the [AWS
	// documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).
	AwsPublicCert pulumi.StringInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// The name of the certificate.
	CertName pulumi.StringInput
	// Either "pkcs7" or "identity", indicating the type of
	// document which can be verified using the given certificate. Defaults to
	// "pkcs7".
	Type pulumi.StringPtrInput
}

func (AuthBackendCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendCertArgs)(nil)).Elem()
}

