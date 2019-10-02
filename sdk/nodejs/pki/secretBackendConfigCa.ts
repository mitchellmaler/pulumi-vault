// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_config_ca.html.markdown.
 */
export class SecretBackendConfigCa extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendConfigCa resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendConfigCaState, opts?: pulumi.CustomResourceOptions): SecretBackendConfigCa {
        return new SecretBackendConfigCa(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pki/secretBackendConfigCa:SecretBackendConfigCa';

    /**
     * Returns true if the given object is an instance of SecretBackendConfigCa.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendConfigCa {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendConfigCa.__pulumiType;
    }

    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The key and certificate PEM bundle
     */
    public readonly pemBundle!: pulumi.Output<string>;

    /**
     * Create a SecretBackendConfigCa resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendConfigCaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendConfigCaArgs | SecretBackendConfigCaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendConfigCaState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["pemBundle"] = state ? state.pemBundle : undefined;
        } else {
            const args = argsOrState as SecretBackendConfigCaArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.pemBundle === undefined) {
                throw new Error("Missing required property 'pemBundle'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["pemBundle"] = args ? args.pemBundle : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendConfigCa.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendConfigCa resources.
 */
export interface SecretBackendConfigCaState {
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The key and certificate PEM bundle
     */
    readonly pemBundle?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendConfigCa resource.
 */
export interface SecretBackendConfigCaArgs {
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * The key and certificate PEM bundle
     */
    readonly pemBundle: pulumi.Input<string>;
}
