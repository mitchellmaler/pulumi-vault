// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/identity_entity.html.markdown.
func LookupEntity(ctx *pulumi.Context, args *LookupEntityArgs, opts ...pulumi.InvokeOption) (*LookupEntityResult, error) {
	var rv LookupEntityResult
	err := ctx.Invoke("vault:identity/getEntity:getEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEntity.
type LookupEntityArgs struct {
	// ID of the alias.
	AliasId *string `pulumi:"aliasId"`
	// Accessor of the mount to which the alias belongs to.
	// This should be supplied in conjunction with `aliasName`.
	AliasMountAccessor *string `pulumi:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with
	// `aliasMountAccessor`.
	AliasName *string `pulumi:"aliasName"`
	// ID of the entity.
	EntityId *string `pulumi:"entityId"`
	// Name of the entity.
	EntityName *string `pulumi:"entityName"`
}


// A collection of values returned by getEntity.
type LookupEntityResult struct {
	AliasId string `pulumi:"aliasId"`
	AliasMountAccessor string `pulumi:"aliasMountAccessor"`
	AliasName string `pulumi:"aliasName"`
	// A list of entity alias. Structure is documented below.
	Aliases []GetEntityAliasType `pulumi:"aliases"`
	// Creation time of the Alias
	CreationTime string `pulumi:"creationTime"`
	// A string containing the full data payload retrieved from
	// Vault, serialized in JSON format.
	DataJson string `pulumi:"dataJson"`
	// List of Group IDs of which the entity is directly a member of
	DirectGroupIds []string `pulumi:"directGroupIds"`
	// Whether the entity is disabled
	Disabled bool `pulumi:"disabled"`
	EntityId string `pulumi:"entityId"`
	EntityName string `pulumi:"entityName"`
	// List of all Group IDs of which the entity is a member of
	GroupIds []string `pulumi:"groupIds"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all Group IDs of which the entity is a member of transitively
	InheritedGroupIds []string `pulumi:"inheritedGroupIds"`
	// Last update time of the alias
	LastUpdateTime string `pulumi:"lastUpdateTime"`
	// Other entity IDs which is merged with this entity
	MergedEntityIds []string `pulumi:"mergedEntityIds"`
	// Arbitrary metadata
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Namespace of which the entity is part of
	NamespaceId string `pulumi:"namespaceId"`
	// List of policies attached to the entity
	Policies []string `pulumi:"policies"`
}

