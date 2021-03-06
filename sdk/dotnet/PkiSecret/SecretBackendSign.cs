// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_sign.html.markdown.
    /// </summary>
    public partial class SecretBackendSign : Pulumi.CustomResource
    {
        /// <summary>
        /// List of alternative names
        /// </summary>
        [Output("altNames")]
        public Output<ImmutableArray<string>> AltNames { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The CA chain
        /// </summary>
        [Output("caChains")]
        public Output<ImmutableArray<string>> CaChains { get; private set; } = null!;

        /// <summary>
        /// The certificate
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Output("commonName")]
        public Output<string> CommonName { get; private set; } = null!;

        /// <summary>
        /// The CSR
        /// </summary>
        [Output("csr")]
        public Output<string> Csr { get; private set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Output("excludeCnFromSans")]
        public Output<bool?> ExcludeCnFromSans { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the certificate in unix epoch format
        /// </summary>
        [Output("expiration")]
        public Output<int> Expiration { get; private set; } = null!;

        /// <summary>
        /// The format of data
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        [Output("ipSans")]
        public Output<ImmutableArray<string>> IpSans { get; private set; } = null!;

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Output("issuingCa")]
        public Output<string> IssuingCa { get; private set; } = null!;

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Output("minSecondsRemaining")]
        public Output<int?> MinSecondsRemaining { get; private set; } = null!;

        /// <summary>
        /// Name of the role to create the certificate against
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of other SANs
        /// </summary>
        [Output("otherSans")]
        public Output<ImmutableArray<string>> OtherSans { get; private set; } = null!;

        /// <summary>
        /// The serial
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// Time to live
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// List of alterative URIs
        /// </summary>
        [Output("uriSans")]
        public Output<ImmutableArray<string>> UriSans { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendSign resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendSign(string name, SecretBackendSignArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendSign(string name, Input<string> id, SecretBackendSignState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendSign resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendSign Get(string name, Input<string> id, SecretBackendSignState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendSign(name, id, state, options);
        }
    }

    public sealed class SecretBackendSignArgs : Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        /// <summary>
        /// The CSR
        /// </summary>
        [Input("csr", required: true)]
        public Input<string> Csr { get; set; } = null!;

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Input("minSecondsRemaining")]
        public Input<int>? MinSecondsRemaining { get; set; }

        /// <summary>
        /// Name of the role to create the certificate against
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alterative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        public SecretBackendSignArgs()
        {
        }
    }

    public sealed class SecretBackendSignState : Pulumi.ResourceArgs
    {
        [Input("altNames")]
        private InputList<string>? _altNames;

        /// <summary>
        /// List of alternative names
        /// </summary>
        public InputList<string> AltNames
        {
            get => _altNames ?? (_altNames = new InputList<string>());
            set => _altNames = value;
        }

        /// <summary>
        /// If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("caChains")]
        private InputList<string>? _caChains;

        /// <summary>
        /// The CA chain
        /// </summary>
        public InputList<string> CaChains
        {
            get => _caChains ?? (_caChains = new InputList<string>());
            set => _caChains = value;
        }

        /// <summary>
        /// The certificate
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// CN of certificate to create
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// The CSR
        /// </summary>
        [Input("csr")]
        public Input<string>? Csr { get; set; }

        /// <summary>
        /// Flag to exclude CN from SANs
        /// </summary>
        [Input("excludeCnFromSans")]
        public Input<bool>? ExcludeCnFromSans { get; set; }

        /// <summary>
        /// The expiration date of the certificate in unix epoch format
        /// </summary>
        [Input("expiration")]
        public Input<int>? Expiration { get; set; }

        /// <summary>
        /// The format of data
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        [Input("ipSans")]
        private InputList<string>? _ipSans;

        /// <summary>
        /// List of alternative IPs
        /// </summary>
        public InputList<string> IpSans
        {
            get => _ipSans ?? (_ipSans = new InputList<string>());
            set => _ipSans = value;
        }

        /// <summary>
        /// The issuing CA
        /// </summary>
        [Input("issuingCa")]
        public Input<string>? IssuingCa { get; set; }

        /// <summary>
        /// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        /// </summary>
        [Input("minSecondsRemaining")]
        public Input<int>? MinSecondsRemaining { get; set; }

        /// <summary>
        /// Name of the role to create the certificate against
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("otherSans")]
        private InputList<string>? _otherSans;

        /// <summary>
        /// List of other SANs
        /// </summary>
        public InputList<string> OtherSans
        {
            get => _otherSans ?? (_otherSans = new InputList<string>());
            set => _otherSans = value;
        }

        /// <summary>
        /// The serial
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// Time to live
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        [Input("uriSans")]
        private InputList<string>? _uriSans;

        /// <summary>
        /// List of alterative URIs
        /// </summary>
        public InputList<string> UriSans
        {
            get => _uriSans ?? (_uriSans = new InputList<string>());
            set => _uriSans = value;
        }

        public SecretBackendSignState()
        {
        }
    }
}
