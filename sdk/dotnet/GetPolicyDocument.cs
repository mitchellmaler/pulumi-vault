// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static partial class Invokes
    {
        /// <summary>
        /// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `vault..Policy` resource.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/policy_document.html.markdown.
        /// </summary>
        public static Task<GetPolicyDocumentResult> GetPolicyDocument(GetPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDocumentResult>("vault:index/getPolicyDocument:getPolicyDocument", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetPolicyDocumentArgs : Pulumi.InvokeArgs
    {
        [Input("rules")]
        private List<Inputs.GetPolicyDocumentRulesArgs>? _rules;
        public List<Inputs.GetPolicyDocumentRulesArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.GetPolicyDocumentRulesArgs>());
            set => _rules = value;
        }

        public GetPolicyDocumentArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentResult
    {
        /// <summary>
        /// The above arguments serialized as a standard Vault HCL policy document.
        /// </summary>
        public readonly string Hcl;
        public readonly ImmutableArray<Outputs.GetPolicyDocumentRulesResult> Rules;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPolicyDocumentResult(
            string hcl,
            ImmutableArray<Outputs.GetPolicyDocumentRulesResult> rules,
            string id)
        {
            Hcl = hcl;
            Rules = rules;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetPolicyDocumentRulesAllowedParametersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// name of permitted or denied parameter.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// list of values what are permitted or denied by policy rule.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetPolicyDocumentRulesAllowedParametersArgs()
        {
        }
    }

    public sealed class GetPolicyDocumentRulesArgs : Pulumi.InvokeArgs
    {
        [Input("allowedParameters")]
        private List<GetPolicyDocumentRulesAllowedParametersArgs>? _allowedParameters;

        /// <summary>
        /// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        /// </summary>
        public List<GetPolicyDocumentRulesAllowedParametersArgs> AllowedParameters
        {
            get => _allowedParameters ?? (_allowedParameters = new List<GetPolicyDocumentRulesAllowedParametersArgs>());
            set => _allowedParameters = value;
        }

        [Input("capabilities", required: true)]
        private List<string>? _capabilities;

        /// <summary>
        /// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        /// </summary>
        public List<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new List<string>());
            set => _capabilities = value;
        }

        [Input("deniedParameters")]
        private List<GetPolicyDocumentRulesDeniedParametersArgs>? _deniedParameters;

        /// <summary>
        /// Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        /// </summary>
        public List<GetPolicyDocumentRulesDeniedParametersArgs> DeniedParameters
        {
            get => _deniedParameters ?? (_deniedParameters = new List<GetPolicyDocumentRulesDeniedParametersArgs>());
            set => _deniedParameters = value;
        }

        /// <summary>
        /// Description of the rule. Will be added as a commend to rendered rule.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The maximum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        [Input("maxWrappingTtl")]
        public string? MaxWrappingTtl { get; set; }

        /// <summary>
        /// The minimum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        [Input("minWrappingTtl")]
        public string? MinWrappingTtl { get; set; }

        /// <summary>
        /// A path in Vault that this rule applies to.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        [Input("requiredParameters")]
        private List<string>? _requiredParameters;

        /// <summary>
        /// A list of parameters that must be specified.
        /// </summary>
        public List<string> RequiredParameters
        {
            get => _requiredParameters ?? (_requiredParameters = new List<string>());
            set => _requiredParameters = value;
        }

        public GetPolicyDocumentRulesArgs()
        {
        }
    }

    public sealed class GetPolicyDocumentRulesDeniedParametersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// name of permitted or denied parameter.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// list of values what are permitted or denied by policy rule.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetPolicyDocumentRulesDeniedParametersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetPolicyDocumentRulesAllowedParametersResult
    {
        /// <summary>
        /// name of permitted or denied parameter.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// list of values what are permitted or denied by policy rule.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetPolicyDocumentRulesAllowedParametersResult(
            string key,
            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentRulesDeniedParametersResult
    {
        /// <summary>
        /// name of permitted or denied parameter.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// list of values what are permitted or denied by policy rule.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetPolicyDocumentRulesDeniedParametersResult(
            string key,
            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetPolicyDocumentRulesResult
    {
        /// <summary>
        /// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        /// </summary>
        public readonly ImmutableArray<GetPolicyDocumentRulesAllowedParametersResult> AllowedParameters;
        /// <summary>
        /// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        /// </summary>
        public readonly ImmutableArray<string> Capabilities;
        /// <summary>
        /// Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        /// </summary>
        public readonly ImmutableArray<GetPolicyDocumentRulesDeniedParametersResult> DeniedParameters;
        /// <summary>
        /// Description of the rule. Will be added as a commend to rendered rule.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The maximum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        public readonly string? MaxWrappingTtl;
        /// <summary>
        /// The minimum allowed TTL that clients can specify for a wrapped response.
        /// </summary>
        public readonly string? MinWrappingTtl;
        /// <summary>
        /// A path in Vault that this rule applies to.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// A list of parameters that must be specified.
        /// </summary>
        public readonly ImmutableArray<string> RequiredParameters;

        [OutputConstructor]
        private GetPolicyDocumentRulesResult(
            ImmutableArray<GetPolicyDocumentRulesAllowedParametersResult> allowedParameters,
            ImmutableArray<string> capabilities,
            ImmutableArray<GetPolicyDocumentRulesDeniedParametersResult> deniedParameters,
            string? description,
            string? maxWrappingTtl,
            string? minWrappingTtl,
            string path,
            ImmutableArray<string> requiredParameters)
        {
            AllowedParameters = allowedParameters;
            Capabilities = capabilities;
            DeniedParameters = deniedParameters;
            Description = description;
            MaxWrappingTtl = maxWrappingTtl;
            MinWrappingTtl = minWrappingTtl;
            Path = path;
            RequiredParameters = requiredParameters;
        }
    }
    }
}
