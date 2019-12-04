// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure the cache for the Transit Secret Backend in Vault.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const transit = new vault.Mount("transit", {
 *     defaultLeaseTtlSeconds: 3600,
 *     description: "Example description",
 *     maxLeaseTtlSeconds: 86400,
 *     path: "transit",
 *     type: "transit",
 * });
 * const cfg = new vault.TransitSecretBackendCacheConfig("cfg", {
 *     backend: transit.path,
 *     size: 500,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/transit_secret_cache_config.html.markdown.
 */
export class SecretCacheConfig extends pulumi.CustomResource {
    /**
     * Get an existing SecretCacheConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretCacheConfigState, opts?: pulumi.CustomResourceOptions): SecretCacheConfig {
        return new SecretCacheConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:transit/secretCacheConfig:SecretCacheConfig';

    /**
     * Returns true if the given object is an instance of SecretCacheConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretCacheConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretCacheConfig.__pulumiType;
    }

    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    public readonly size!: pulumi.Output<number>;

    /**
     * Create a SecretCacheConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretCacheConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretCacheConfigArgs | SecretCacheConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretCacheConfigState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["size"] = state ? state.size : undefined;
        } else {
            const args = argsOrState as SecretCacheConfigArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["size"] = args ? args.size : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretCacheConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretCacheConfig resources.
 */
export interface SecretCacheConfigState {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    readonly size?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretCacheConfig resource.
 */
export interface SecretCacheConfigArgs {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    readonly size: pulumi.Input<number>;
}