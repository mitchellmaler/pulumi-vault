// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AppRole
{
    public static partial class Invokes
    {
        /// <summary>
        /// Reads the Role ID of an AppRole from a Vault server.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/approle_auth_backend_role_id.html.markdown.
        /// </summary>
        public static Task<GetAuthBackendRoleIdResult> GetAuthBackendRoleId(GetAuthBackendRoleIdArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendRoleIdResult>("vault:appRole/getAuthBackendRoleId:getAuthBackendRoleId", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAuthBackendRoleIdArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name for the AppRole backend the role to
        /// retrieve a RoleID for resides in. Defaults to "approle".
        /// </summary>
        [Input("backend")]
        public string? Backend { get; set; }

        /// <summary>
        /// The name of the role to retrieve the Role ID for.
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        public GetAuthBackendRoleIdArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAuthBackendRoleIdResult
    {
        public readonly string? Backend;
        /// <summary>
        /// The RoleID of the role.
        /// </summary>
        public readonly string RoleId;
        public readonly string RoleName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAuthBackendRoleIdResult(
            string? backend,
            string roleId,
            string roleName,
            string id)
        {
            Backend = backend;
            RoleId = roleId;
            RoleName = roleName;
            Id = id;
        }
    }
}
