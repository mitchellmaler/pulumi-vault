// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_cert.html.markdown.
 */
export class BackendCert extends pulumi.CustomResource {
    /**
     * Get an existing BackendCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendCertState, opts?: pulumi.CustomResourceOptions): BackendCert {
        return new BackendCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/backendCert:BackendCert';

    /**
     * Returns true if the given object is an instance of BackendCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendCert.__pulumiType;
    }

    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The CA chain
     */
    public /*out*/ readonly caChain!: pulumi.Output<string>;
    /**
     * The certificate
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * CN of certificate to create
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * Flag to exclude CN from SANs
     */
    public readonly excludeCnFromSans!: pulumi.Output<boolean | undefined>;
    /**
     * The expiration date of the certificate in unix epoch format
     */
    public /*out*/ readonly expiration!: pulumi.Output<number>;
    /**
     * The format of data
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * List of alternative IPs
     */
    public readonly ipSans!: pulumi.Output<string[] | undefined>;
    /**
     * The issuing CA
     */
    public /*out*/ readonly issuingCa!: pulumi.Output<string>;
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    public readonly minSecondsRemaining!: pulumi.Output<number | undefined>;
    /**
     * Name of the role to create the certificate against
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of other SANs
     */
    public readonly otherSans!: pulumi.Output<string[] | undefined>;
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
     * The serial number
     */
    public /*out*/ readonly serialNumber!: pulumi.Output<string>;
    /**
     * Time to live
     */
    public readonly ttl!: pulumi.Output<string | undefined>;

    /**
     * Create a BackendCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendCertArgs | BackendCertState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendCertState | undefined;
            inputs["altNames"] = state ? state.altNames : undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["caChain"] = state ? state.caChain : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["commonName"] = state ? state.commonName : undefined;
            inputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            inputs["expiration"] = state ? state.expiration : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["ipSans"] = state ? state.ipSans : undefined;
            inputs["issuingCa"] = state ? state.issuingCa : undefined;
            inputs["minSecondsRemaining"] = state ? state.minSecondsRemaining : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["otherSans"] = state ? state.otherSans : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["privateKeyFormat"] = state ? state.privateKeyFormat : undefined;
            inputs["privateKeyType"] = state ? state.privateKeyType : undefined;
            inputs["serialNumber"] = state ? state.serialNumber : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as BackendCertArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.commonName === undefined) {
                throw new Error("Missing required property 'commonName'");
            }
            inputs["altNames"] = args ? args.altNames : undefined;
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["commonName"] = args ? args.commonName : undefined;
            inputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["ipSans"] = args ? args.ipSans : undefined;
            inputs["minSecondsRemaining"] = args ? args.minSecondsRemaining : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["otherSans"] = args ? args.otherSans : undefined;
            inputs["privateKeyFormat"] = args ? args.privateKeyFormat : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["caChain"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["expiration"] = undefined /*out*/;
            inputs["issuingCa"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
            inputs["privateKeyType"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BackendCert.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendCert resources.
 */
export interface BackendCertState {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The CA chain
     */
    readonly caChain?: pulumi.Input<string>;
    /**
     * The certificate
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * CN of certificate to create
     */
    readonly commonName?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    readonly excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The expiration date of the certificate in unix epoch format
     */
    readonly expiration?: pulumi.Input<number>;
    /**
     * The format of data
     */
    readonly format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    readonly ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The issuing CA
     */
    readonly issuingCa?: pulumi.Input<string>;
    /**
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    readonly minSecondsRemaining?: pulumi.Input<number>;
    /**
     * Name of the role to create the certificate against
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
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
     * The serial number
     */
    readonly serialNumber?: pulumi.Input<string>;
    /**
     * Time to live
     */
    readonly ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendCert resource.
 */
export interface BackendCertArgs {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * CN of certificate to create
     */
    readonly commonName: pulumi.Input<string>;
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
     * Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
     */
    readonly minSecondsRemaining?: pulumi.Input<number>;
    /**
     * Name of the role to create the certificate against
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The private key format
     */
    readonly privateKeyFormat?: pulumi.Input<string>;
    /**
     * Time to live
     */
    readonly ttl?: pulumi.Input<string>;
}
