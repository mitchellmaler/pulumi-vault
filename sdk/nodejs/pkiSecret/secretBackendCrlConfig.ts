// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const pki = new vault.Mount("pki", {
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 *     path: "%s",
 *     type: "pki",
 * });
 * const crlConfig = new vault.pkiSecret.SecretBackendCrlConfig("crlConfig", {
 *     backend: pki.path,
 *     disable: false,
 *     expiry: "72h",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_crl_config.html.markdown.
 */
export class SecretBackendCrlConfig extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendCrlConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendCrlConfigState, opts?: pulumi.CustomResourceOptions): SecretBackendCrlConfig {
        return new SecretBackendCrlConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig';

    /**
     * Returns true if the given object is an instance of SecretBackendCrlConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendCrlConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendCrlConfig.__pulumiType;
    }

    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Disables or enables CRL building.
     */
    public readonly disable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the time until expiration.
     */
    public readonly expiry!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackendCrlConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendCrlConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendCrlConfigArgs | SecretBackendCrlConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendCrlConfigState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["expiry"] = state ? state.expiry : undefined;
        } else {
            const args = argsOrState as SecretBackendCrlConfigArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["expiry"] = args ? args.expiry : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendCrlConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendCrlConfig resources.
 */
export interface SecretBackendCrlConfigState {
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Disables or enables CRL building.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Specifies the time until expiration.
     */
    readonly expiry?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendCrlConfig resource.
 */
export interface SecretBackendCrlConfigArgs {
    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Disables or enables CRL building.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Specifies the time until expiration.
     */
    readonly expiry?: pulumi.Input<string>;
}
