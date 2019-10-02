// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_intermediate_cert_request.html.markdown.
 */
export class SecretBackendIntermediateCertRequest extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendIntermediateCertRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendIntermediateCertRequestState, opts?: pulumi.CustomResourceOptions): SecretBackendIntermediateCertRequest {
        return new SecretBackendIntermediateCertRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pki/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest';

    /**
     * Returns true if the given object is an instance of SecretBackendIntermediateCertRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendIntermediateCertRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendIntermediateCertRequest.__pulumiType;
    }

    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * CN of intermediate to create
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * The country
     */
    public readonly country!: pulumi.Output<string | undefined>;
    /**
     * The CSR
     */
    public /*out*/ readonly csr!: pulumi.Output<string>;
    /**
     * Flag to exclude CN from SANs
     */
    public readonly excludeCnFromSans!: pulumi.Output<boolean | undefined>;
    /**
     * The format of data
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * List of alternative IPs
     */
    public readonly ipSans!: pulumi.Output<string[] | undefined>;
    /**
     * The number of bits to use
     */
    public readonly keyBits!: pulumi.Output<number | undefined>;
    /**
     * The desired key type
     */
    public readonly keyType!: pulumi.Output<string | undefined>;
    /**
     * The locality
     */
    public readonly locality!: pulumi.Output<string | undefined>;
    /**
     * The organization
     */
    public readonly organization!: pulumi.Output<string | undefined>;
    /**
     * List of other SANs
     */
    public readonly otherSans!: pulumi.Output<string[] | undefined>;
    /**
     * The organization unit
     */
    public readonly ou!: pulumi.Output<string | undefined>;
    /**
     * The postal code
     */
    public readonly postalCode!: pulumi.Output<string | undefined>;
    /**
     * The private key
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * The private key format
     */
    public readonly privateKeyFormat!: pulumi.Output<string | undefined>;
    /**
     * The private key type
     */
    public /*out*/ readonly privateKeyType!: pulumi.Output<string>;
    /**
     * The province
     */
    public readonly province!: pulumi.Output<string | undefined>;
    /**
     * The street address
     */
    public readonly streetAddress!: pulumi.Output<string | undefined>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * List of alternative URIs
     */
    public readonly uriSans!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendIntermediateCertRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendIntermediateCertRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendIntermediateCertRequestArgs | SecretBackendIntermediateCertRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendIntermediateCertRequestState | undefined;
            inputs["altNames"] = state ? state.altNames : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["commonName"] = state ? state.commonName : undefined;
            inputs["country"] = state ? state.country : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["ipSans"] = state ? state.ipSans : undefined;
            inputs["keyBits"] = state ? state.keyBits : undefined;
            inputs["keyType"] = state ? state.keyType : undefined;
            inputs["locality"] = state ? state.locality : undefined;
            inputs["organization"] = state ? state.organization : undefined;
            inputs["otherSans"] = state ? state.otherSans : undefined;
            inputs["ou"] = state ? state.ou : undefined;
            inputs["postalCode"] = state ? state.postalCode : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["privateKeyFormat"] = state ? state.privateKeyFormat : undefined;
            inputs["privateKeyType"] = state ? state.privateKeyType : undefined;
            inputs["province"] = state ? state.province : undefined;
            inputs["streetAddress"] = state ? state.streetAddress : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uriSans"] = state ? state.uriSans : undefined;
        } else {
            const args = argsOrState as SecretBackendIntermediateCertRequestArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.commonName === undefined) {
                throw new Error("Missing required property 'commonName'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["altNames"] = args ? args.altNames : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["commonName"] = args ? args.commonName : undefined;
            inputs["country"] = args ? args.country : undefined;
            inputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["ipSans"] = args ? args.ipSans : undefined;
            inputs["keyBits"] = args ? args.keyBits : undefined;
            inputs["keyType"] = args ? args.keyType : undefined;
            inputs["locality"] = args ? args.locality : undefined;
            inputs["organization"] = args ? args.organization : undefined;
            inputs["otherSans"] = args ? args.otherSans : undefined;
            inputs["ou"] = args ? args.ou : undefined;
            inputs["postalCode"] = args ? args.postalCode : undefined;
            inputs["privateKeyFormat"] = args ? args.privateKeyFormat : undefined;
            inputs["province"] = args ? args.province : undefined;
            inputs["streetAddress"] = args ? args.streetAddress : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uriSans"] = args ? args.uriSans : undefined;
            inputs["csr"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["privateKeyType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendIntermediateCertRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendIntermediateCertRequest resources.
 */
export interface SecretBackendIntermediateCertRequestState {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    readonly commonName?: pulumi.Input<string>;
    /**
     * The country
     */
    readonly country?: pulumi.Input<string>;
    /**
     * The CSR
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    readonly excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    readonly format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    readonly ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of bits to use
     */
    readonly keyBits?: pulumi.Input<number>;
    /**
     * The desired key type
     */
    readonly keyType?: pulumi.Input<string>;
    /**
     * The locality
     */
    readonly locality?: pulumi.Input<string>;
    /**
     * The organization
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    readonly ou?: pulumi.Input<string>;
    /**
     * The postal code
     */
    readonly postalCode?: pulumi.Input<string>;
    /**
     * The private key
     */
    readonly privateKey?: pulumi.Input<string>;
    /**
     * The private key format
     */
    readonly privateKeyFormat?: pulumi.Input<string>;
    /**
     * The private key type
     */
    readonly privateKeyType?: pulumi.Input<string>;
    /**
     * The province
     */
    readonly province?: pulumi.Input<string>;
    /**
     * The street address
     */
    readonly streetAddress?: pulumi.Input<string>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     */
    readonly type?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    readonly uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendIntermediateCertRequest resource.
 */
export interface SecretBackendIntermediateCertRequestArgs {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    readonly commonName: pulumi.Input<string>;
    /**
     * The country
     */
    readonly country?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    readonly excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    readonly format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    readonly ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of bits to use
     */
    readonly keyBits?: pulumi.Input<number>;
    /**
     * The desired key type
     */
    readonly keyType?: pulumi.Input<string>;
    /**
     * The locality
     */
    readonly locality?: pulumi.Input<string>;
    /**
     * The organization
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    readonly ou?: pulumi.Input<string>;
    /**
     * The postal code
     */
    readonly postalCode?: pulumi.Input<string>;
    /**
     * The private key format
     */
    readonly privateKeyFormat?: pulumi.Input<string>;
    /**
     * The province
     */
    readonly province?: pulumi.Input<string>;
    /**
     * The street address
     */
    readonly streetAddress?: pulumi.Input<string>;
    /**
     * Type of intermediate to create. Must be either \"exported\" or \"internal\"
     */
    readonly type: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    readonly uriSans?: pulumi.Input<pulumi.Input<string>[]>;
}
