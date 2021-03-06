# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GroupPolicies(pulumi.CustomResource):
    exclusive: pulumi.Output[bool]
    """
    Defaults to `true`.
    """
    group_id: pulumi.Output[str]
    """
    Group ID to assign policies to.
    """
    group_name: pulumi.Output[str]
    """
    The name of the group that are assigned the policies.
    """
    policies: pulumi.Output[list]
    """
    List of policies to assign to the group
    """
    def __init__(__self__, resource_name, opts=None, exclusive=None, group_id=None, policies=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclusive: Defaults to `true`.
        :param pulumi.Input[str] group_id: Group ID to assign policies to.
        :param pulumi.Input[list] policies: List of policies to assign to the group

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group_policies.html.markdown.
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

            __props__['exclusive'] = exclusive
            if group_id is None:
                raise TypeError("Missing required property 'group_id'")
            __props__['group_id'] = group_id
            if policies is None:
                raise TypeError("Missing required property 'policies'")
            __props__['policies'] = policies
            __props__['group_name'] = None
        super(GroupPolicies, __self__).__init__(
            'vault:identity/groupPolicies:GroupPolicies',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, exclusive=None, group_id=None, group_name=None, policies=None):
        """
        Get an existing GroupPolicies resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] exclusive: Defaults to `true`.
        :param pulumi.Input[str] group_id: Group ID to assign policies to.
        :param pulumi.Input[str] group_name: The name of the group that are assigned the policies.
        :param pulumi.Input[list] policies: List of policies to assign to the group

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group_policies.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["exclusive"] = exclusive
        __props__["group_id"] = group_id
        __props__["group_name"] = group_name
        __props__["policies"] = policies
        return GroupPolicies(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

