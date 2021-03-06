// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 * 
 * A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.
 * 
 * ## Example Usage
 * 
 * ### Internal Group
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const internal = new vault.identity.Group("internal", {
 *     metadata: {
 *         version: "2",
 *     },
 *     policies: [
 *         "dev",
 *         "test",
 *     ],
 *     type: "internal",
 * });
 * ```
 * 
 * ### External Group
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const group = new vault.identity.Group("group", {
 *     metadata: {
 *         version: "1",
 *     },
 *     policies: ["test"],
 *     type: "external",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group.html.markdown.
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
     */
    public readonly externalPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
     */
    public readonly memberEntityIds!: pulumi.Output<string[]>;
    /**
     * A list of Group IDs to be assigned as group members.
     */
    public readonly memberGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * A Map of additional metadata to associate with the group.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the identity group to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of policies to apply to the group.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * Type of the group, internal or external. Defaults to `internal`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["externalPolicies"] = state ? state.externalPolicies : undefined;
            inputs["memberEntityIds"] = state ? state.memberEntityIds : undefined;
            inputs["memberGroupIds"] = state ? state.memberGroupIds : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            inputs["externalPolicies"] = args ? args.externalPolicies : undefined;
            inputs["memberEntityIds"] = args ? args.memberEntityIds : undefined;
            inputs["memberGroupIds"] = args ? args.memberGroupIds : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
     */
    readonly externalPolicies?: pulumi.Input<boolean>;
    /**
     * A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
     */
    readonly memberEntityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Group IDs to be assigned as group members.
     */
    readonly memberGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Map of additional metadata to associate with the group.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the identity group to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of policies to apply to the group.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the group, internal or external. Defaults to `internal`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
     */
    readonly externalPolicies?: pulumi.Input<boolean>;
    /**
     * A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
     */
    readonly memberEntityIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Group IDs to be assigned as group members.
     */
    readonly memberGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Map of additional metadata to associate with the group.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the identity group to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of policies to apply to the group.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the group, internal or external. Defaults to `internal`.
     */
    readonly type?: pulumi.Input<string>;
}
