// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_client.html.markdown.
 */
export class AuthBackendClient extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendClientState, opts?: pulumi.CustomResourceOptions): AuthBackendClient {
        return new AuthBackendClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendClient:AuthBackendClient';

    /**
     * Returns true if the given object is an instance of AuthBackendClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendClient.__pulumiType;
    }

    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    public readonly ec2Endpoint!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    public readonly iamEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    public readonly iamServerIdHeaderValue!: pulumi.Output<string | undefined>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    public readonly stsEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendClientArgs | AuthBackendClientState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendClientState | undefined;
            inputs["accessKey"] = state ? state.accessKey : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["ec2Endpoint"] = state ? state.ec2Endpoint : undefined;
            inputs["iamEndpoint"] = state ? state.iamEndpoint : undefined;
            inputs["iamServerIdHeaderValue"] = state ? state.iamServerIdHeaderValue : undefined;
            inputs["secretKey"] = state ? state.secretKey : undefined;
            inputs["stsEndpoint"] = state ? state.stsEndpoint : undefined;
        } else {
            const args = argsOrState as AuthBackendClientArgs | undefined;
            inputs["accessKey"] = args ? args.accessKey : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["ec2Endpoint"] = args ? args.ec2Endpoint : undefined;
            inputs["iamEndpoint"] = args ? args.iamEndpoint : undefined;
            inputs["iamServerIdHeaderValue"] = args ? args.iamServerIdHeaderValue : undefined;
            inputs["secretKey"] = args ? args.secretKey : undefined;
            inputs["stsEndpoint"] = args ? args.stsEndpoint : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendClient.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendClient resources.
 */
export interface AuthBackendClientState {
    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    readonly accessKey?: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    readonly ec2Endpoint?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    readonly iamEndpoint?: pulumi.Input<string>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    readonly iamServerIdHeaderValue?: pulumi.Input<string>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    readonly stsEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendClient resource.
 */
export interface AuthBackendClientArgs {
    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     */
    readonly accessKey?: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     */
    readonly ec2Endpoint?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     */
    readonly iamEndpoint?: pulumi.Input<string>;
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     */
    readonly iamServerIdHeaderValue?: pulumi.Input<string>;
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     */
    readonly stsEndpoint?: pulumi.Input<string>;
}
