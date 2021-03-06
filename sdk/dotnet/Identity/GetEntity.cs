// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    public static partial class Invokes
    {
        /// <summary>
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/identity_entity.html.markdown.
        /// </summary>
        public static Task<GetEntityResult> GetEntity(GetEntityArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntityResult>("vault:identity/getEntity:getEntity", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetEntityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the alias.
        /// </summary>
        [Input("aliasId")]
        public string? AliasId { get; set; }

        /// <summary>
        /// Accessor of the mount to which the alias belongs to.
        /// This should be supplied in conjunction with `alias_name`.
        /// </summary>
        [Input("aliasMountAccessor")]
        public string? AliasMountAccessor { get; set; }

        /// <summary>
        /// Name of the alias. This should be supplied in conjunction with
        /// `alias_mount_accessor`.
        /// </summary>
        [Input("aliasName")]
        public string? AliasName { get; set; }

        /// <summary>
        /// ID of the entity.
        /// </summary>
        [Input("entityId")]
        public string? EntityId { get; set; }

        /// <summary>
        /// Name of the entity.
        /// </summary>
        [Input("entityName")]
        public string? EntityName { get; set; }

        public GetEntityArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetEntityResult
    {
        public readonly string AliasId;
        public readonly string AliasMountAccessor;
        public readonly string AliasName;
        /// <summary>
        /// A list of entity alias. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEntityAliasesResult> Aliases;
        /// <summary>
        /// Creation time of the Alias
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// A string containing the full data payload retrieved from
        /// Vault, serialized in JSON format.
        /// </summary>
        public readonly string DataJson;
        /// <summary>
        /// List of Group IDs of which the entity is directly a member of
        /// </summary>
        public readonly ImmutableArray<string> DirectGroupIds;
        /// <summary>
        /// Whether the entity is disabled
        /// </summary>
        public readonly bool Disabled;
        public readonly string EntityId;
        public readonly string EntityName;
        /// <summary>
        /// List of all Group IDs of which the entity is a member of
        /// </summary>
        public readonly ImmutableArray<string> GroupIds;
        /// <summary>
        /// List of all Group IDs of which the entity is a member of transitively
        /// </summary>
        public readonly ImmutableArray<string> InheritedGroupIds;
        /// <summary>
        /// Last update time of the alias
        /// </summary>
        public readonly string LastUpdateTime;
        /// <summary>
        /// Other entity IDs which is merged with this entity
        /// </summary>
        public readonly ImmutableArray<string> MergedEntityIds;
        /// <summary>
        /// Arbitrary metadata
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// Namespace of which the entity is part of
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// List of policies attached to the entity
        /// </summary>
        public readonly ImmutableArray<string> Policies;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetEntityResult(
            string aliasId,
            string aliasMountAccessor,
            string aliasName,
            ImmutableArray<Outputs.GetEntityAliasesResult> aliases,
            string creationTime,
            string dataJson,
            ImmutableArray<string> directGroupIds,
            bool disabled,
            string entityId,
            string entityName,
            ImmutableArray<string> groupIds,
            ImmutableArray<string> inheritedGroupIds,
            string lastUpdateTime,
            ImmutableArray<string> mergedEntityIds,
            ImmutableDictionary<string, object> metadata,
            string namespaceId,
            ImmutableArray<string> policies,
            string id)
        {
            AliasId = aliasId;
            AliasMountAccessor = aliasMountAccessor;
            AliasName = aliasName;
            Aliases = aliases;
            CreationTime = creationTime;
            DataJson = dataJson;
            DirectGroupIds = directGroupIds;
            Disabled = disabled;
            EntityId = entityId;
            EntityName = entityName;
            GroupIds = groupIds;
            InheritedGroupIds = inheritedGroupIds;
            LastUpdateTime = lastUpdateTime;
            MergedEntityIds = mergedEntityIds;
            Metadata = metadata;
            NamespaceId = namespaceId;
            Policies = policies;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetEntityAliasesResult
    {
        /// <summary>
        /// Canonical ID of the Alias
        /// </summary>
        public readonly string CanonicalId;
        /// <summary>
        /// Creation time of the Alias
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// ID of the alias
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Last update time of the alias
        /// </summary>
        public readonly string LastUpdateTime;
        /// <summary>
        /// List of canonical IDs merged with this alias
        /// </summary>
        public readonly ImmutableArray<string> MergedFromCanonicalIds;
        /// <summary>
        /// Arbitrary metadata
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// Authentication mount acccessor which this alias belongs to
        /// </summary>
        public readonly string MountAccessor;
        /// <summary>
        /// Authentication mount path which this alias belongs to
        /// </summary>
        public readonly string MountPath;
        /// <summary>
        /// Authentication mount type which this alias belongs to
        /// </summary>
        public readonly string MountType;
        /// <summary>
        /// Name of the alias
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetEntityAliasesResult(
            string canonicalId,
            string creationTime,
            string id,
            string lastUpdateTime,
            ImmutableArray<string> mergedFromCanonicalIds,
            ImmutableDictionary<string, object> metadata,
            string mountAccessor,
            string mountPath,
            string mountType,
            string name)
        {
            CanonicalId = canonicalId;
            CreationTime = creationTime;
            Id = id;
            LastUpdateTime = lastUpdateTime;
            MergedFromCanonicalIds = mergedFromCanonicalIds;
            Metadata = metadata;
            MountAccessor = mountAccessor;
            MountPath = mountPath;
            MountType = mountType;
            Name = name;
        }
    }
    }
}
