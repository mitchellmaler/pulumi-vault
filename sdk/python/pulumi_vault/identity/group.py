# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Group(pulumi.CustomResource):
    external_policies: pulumi.Output[bool]
    """
    `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
    """
    member_entity_ids: pulumi.Output[list]
    """
    A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
    """
    member_group_ids: pulumi.Output[list]
    """
    A list of Group IDs to be assigned as group members.
    """
    metadata: pulumi.Output[dict]
    """
    A Map of additional metadata to associate with the group.
    """
    name: pulumi.Output[str]
    """
    Name of the identity group to create.
    """
    policies: pulumi.Output[list]
    """
    A list of policies to apply to the group.
    """
    type: pulumi.Output[str]
    """
    Type of the group, internal or external. Defaults to `internal`.
    """
    def __init__(__self__, resource_name, opts=None, external_policies=None, member_entity_ids=None, member_group_ids=None, metadata=None, name=None, policies=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
        
        A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] external_policies: `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
        :param pulumi.Input[list] member_entity_ids: A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        :param pulumi.Input[list] member_group_ids: A list of Group IDs to be assigned as group members.
        :param pulumi.Input[dict] metadata: A Map of additional metadata to associate with the group.
        :param pulumi.Input[str] name: Name of the identity group to create.
        :param pulumi.Input[list] policies: A list of policies to apply to the group.
        :param pulumi.Input[str] type: Type of the group, internal or external. Defaults to `internal`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['external_policies'] = external_policies
            __props__['member_entity_ids'] = member_entity_ids
            __props__['member_group_ids'] = member_group_ids
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['policies'] = policies
            __props__['type'] = type
        super(Group, __self__).__init__(
            'vault:identity/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, external_policies=None, member_entity_ids=None, member_group_ids=None, metadata=None, name=None, policies=None, type=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] external_policies: `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
        :param pulumi.Input[list] member_entity_ids: A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        :param pulumi.Input[list] member_group_ids: A list of Group IDs to be assigned as group members.
        :param pulumi.Input[dict] metadata: A Map of additional metadata to associate with the group.
        :param pulumi.Input[str] name: Name of the identity group to create.
        :param pulumi.Input[list] policies: A list of policies to apply to the group.
        :param pulumi.Input[str] type: Type of the group, internal or external. Defaults to `internal`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["external_policies"] = external_policies
        __props__["member_entity_ids"] = member_entity_ids
        __props__["member_group_ids"] = member_group_ids
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["policies"] = policies
        __props__["type"] = type
        return Group(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

