# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Token(pulumi.CustomResource):
    client_token: pulumi.Output[str]
    """
    String containing the client token if stored in present file
    """
    display_name: pulumi.Output[str]
    """
    String containing the token display name
    """
    explicit_max_ttl: pulumi.Output[str]
    """
    The explicit max TTL of this token
    """
    lease_duration: pulumi.Output[float]
    """
    String containing the token lease duration if present in state file
    """
    lease_started: pulumi.Output[str]
    """
    String containing the token lease started time if present in state file
    """
    no_default_policy: pulumi.Output[bool]
    """
    Flag to not attach the default policy to this token
    """
    no_parent: pulumi.Output[bool]
    """
    Flag to create a token without parent
    """
    num_uses: pulumi.Output[float]
    """
    The number of allowed uses of this token
    """
    period: pulumi.Output[str]
    """
    The period of this token
    """
    policies: pulumi.Output[list]
    """
    List of policies to attach to this token
    """
    renew_increment: pulumi.Output[float]
    """
    The renew increment
    """
    renew_min_lease: pulumi.Output[float]
    """
    The minimal lease to renew this token
    """
    renewable: pulumi.Output[bool]
    """
    Flag to allow to renew this token
    """
    role_name: pulumi.Output[str]
    """
    The token role name
    """
    ttl: pulumi.Output[str]
    """
    The TTL period of this token
    """
    wrapped_token: pulumi.Output[str]
    wrapping_accessor: pulumi.Output[str]
    wrapping_ttl: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, display_name=None, explicit_max_ttl=None, no_default_policy=None, no_parent=None, num_uses=None, period=None, policies=None, renew_increment=None, renew_min_lease=None, renewable=None, role_name=None, ttl=None, wrapping_ttl=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Token resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: String containing the token display name
        :param pulumi.Input[str] explicit_max_ttl: The explicit max TTL of this token
        :param pulumi.Input[bool] no_default_policy: Flag to not attach the default policy to this token
        :param pulumi.Input[bool] no_parent: Flag to create a token without parent
        :param pulumi.Input[float] num_uses: The number of allowed uses of this token
        :param pulumi.Input[str] period: The period of this token
        :param pulumi.Input[list] policies: List of policies to attach to this token
        :param pulumi.Input[float] renew_increment: The renew increment
        :param pulumi.Input[float] renew_min_lease: The minimal lease to renew this token
        :param pulumi.Input[bool] renewable: Flag to allow to renew this token
        :param pulumi.Input[str] role_name: The token role name
        :param pulumi.Input[str] ttl: The TTL period of this token

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token.html.markdown.
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

            __props__['display_name'] = display_name
            __props__['explicit_max_ttl'] = explicit_max_ttl
            __props__['no_default_policy'] = no_default_policy
            __props__['no_parent'] = no_parent
            __props__['num_uses'] = num_uses
            __props__['period'] = period
            __props__['policies'] = policies
            __props__['renew_increment'] = renew_increment
            __props__['renew_min_lease'] = renew_min_lease
            __props__['renewable'] = renewable
            __props__['role_name'] = role_name
            __props__['ttl'] = ttl
            __props__['wrapping_ttl'] = wrapping_ttl
            __props__['client_token'] = None
            __props__['lease_duration'] = None
            __props__['lease_started'] = None
            __props__['wrapped_token'] = None
            __props__['wrapping_accessor'] = None
        super(Token, __self__).__init__(
            'vault:index/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_token=None, display_name=None, explicit_max_ttl=None, lease_duration=None, lease_started=None, no_default_policy=None, no_parent=None, num_uses=None, period=None, policies=None, renew_increment=None, renew_min_lease=None, renewable=None, role_name=None, ttl=None, wrapped_token=None, wrapping_accessor=None, wrapping_ttl=None):
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_token: String containing the client token if stored in present file
        :param pulumi.Input[str] display_name: String containing the token display name
        :param pulumi.Input[str] explicit_max_ttl: The explicit max TTL of this token
        :param pulumi.Input[float] lease_duration: String containing the token lease duration if present in state file
        :param pulumi.Input[str] lease_started: String containing the token lease started time if present in state file
        :param pulumi.Input[bool] no_default_policy: Flag to not attach the default policy to this token
        :param pulumi.Input[bool] no_parent: Flag to create a token without parent
        :param pulumi.Input[float] num_uses: The number of allowed uses of this token
        :param pulumi.Input[str] period: The period of this token
        :param pulumi.Input[list] policies: List of policies to attach to this token
        :param pulumi.Input[float] renew_increment: The renew increment
        :param pulumi.Input[float] renew_min_lease: The minimal lease to renew this token
        :param pulumi.Input[bool] renewable: Flag to allow to renew this token
        :param pulumi.Input[str] role_name: The token role name
        :param pulumi.Input[str] ttl: The TTL period of this token

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["client_token"] = client_token
        __props__["display_name"] = display_name
        __props__["explicit_max_ttl"] = explicit_max_ttl
        __props__["lease_duration"] = lease_duration
        __props__["lease_started"] = lease_started
        __props__["no_default_policy"] = no_default_policy
        __props__["no_parent"] = no_parent
        __props__["num_uses"] = num_uses
        __props__["period"] = period
        __props__["policies"] = policies
        __props__["renew_increment"] = renew_increment
        __props__["renew_min_lease"] = renew_min_lease
        __props__["renewable"] = renewable
        __props__["role_name"] = role_name
        __props__["ttl"] = ttl
        __props__["wrapped_token"] = wrapped_token
        __props__["wrapping_accessor"] = wrapping_accessor
        __props__["wrapping_ttl"] = wrapping_ttl
        return Token(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

