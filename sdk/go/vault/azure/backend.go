// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package azure

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/azure_secret_backend.html.markdown.
type Backend struct {
	pulumi.CustomResourceState

	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
	// required.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// Path to mount the backend at.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The tenant id for the Azure Active Directory organization.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil || args.SubscriptionId == nil {
		return nil, errors.New("missing required argument 'SubscriptionId'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &BackendArgs{}
	}
	var resource Backend
	err := ctx.RegisterResource("vault:azure/backend:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("vault:azure/backend:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
	// required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment *string `pulumi:"environment"`
	// Path to mount the backend at.
	Path *string `pulumi:"path"`
	// The subscription id for the Azure Active Directory.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The tenant id for the Azure Active Directory organization.
	TenantId *string `pulumi:"tenantId"`
}

type BackendState struct {
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
	// required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the Azure APIs
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment pulumi.StringPtrInput
	// Path to mount the backend at.
	Path pulumi.StringPtrInput
	// The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringPtrInput
	// The tenant id for the Azure Active Directory organization.
	TenantId pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
	// required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the Azure APIs
	ClientSecret *string `pulumi:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	Description *string `pulumi:"description"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment *string `pulumi:"environment"`
	// Path to mount the backend at.
	Path *string `pulumi:"path"`
	// The subscription id for the Azure Active Directory.
	SubscriptionId string `pulumi:"subscriptionId"`
	// The tenant id for the Azure Active Directory organization.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
	// required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the Azure APIs
	ClientSecret pulumi.StringPtrInput
	// Human-friendly description of the mount for the backend.
	Description pulumi.StringPtrInput
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	Environment pulumi.StringPtrInput
	// Path to mount the backend at.
	Path pulumi.StringPtrInput
	// The subscription id for the Azure Active Directory.
	SubscriptionId pulumi.StringInput
	// The tenant id for the Azure Active Directory organization.
	TenantId pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

