// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vault

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `.Policy` resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/policy_document.html.markdown.
func GetPolicyDocument(ctx *pulumi.Context, args *GetPolicyDocumentArgs, opts ...pulumi.InvokeOption) (*GetPolicyDocumentResult, error) {
	var rv GetPolicyDocumentResult
	err := ctx.Invoke("vault:index/getPolicyDocument:getPolicyDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentArgs struct {
	Rules []GetPolicyDocumentRule `pulumi:"rules"`
}


// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResult struct {
	// The above arguments serialized as a standard Vault HCL policy document.
	Hcl string `pulumi:"hcl"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Rules []GetPolicyDocumentRule `pulumi:"rules"`
}

