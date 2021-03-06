// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_secret_backend.html.markdown.
    /// </summary>
    public partial class SecretBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// The GCP service account credentials in JSON format.
        /// </summary>
        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend. Defaults to '0'.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int?> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend. Defaults to '0'.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int?> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `gcp`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:gcp/secretBackend:SecretBackend", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:gcp/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackend Get(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackend(name, id, state, options);
        }
    }

    public sealed class SecretBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GCP service account credentials in JSON format.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend. Defaults to '0'.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend. Defaults to '0'.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `gcp`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public SecretBackendArgs()
        {
        }
    }

    public sealed class SecretBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GCP service account credentials in JSON format.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend. Defaults to '0'.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend. Defaults to '0'.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `gcp`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public SecretBackendState()
        {
        }
    }
}
