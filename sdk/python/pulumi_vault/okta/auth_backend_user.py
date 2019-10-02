# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendUser(pulumi.CustomResource):
    groups: pulumi.Output[list]
    """
    List of Okta groups to associate with this user
    """
    path: pulumi.Output[str]
    """
    The path where the Okta auth backend is mounted
    """
    policies: pulumi.Output[list]
    """
    List of Vault policies to associate with this user
    """
    username: pulumi.Output[str]
    """
    Name of the user within Okta
    """
    def __init__(__self__, resource_name, opts=None, groups=None, path=None, policies=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a user in an
        [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] groups: List of Okta groups to associate with this user
        :param pulumi.Input[str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[list] policies: List of Vault policies to associate with this user
        :param pulumi.Input[str] username: Name of the user within Okta

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_user.html.markdown.
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

            __props__['groups'] = groups
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['policies'] = policies
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(AuthBackendUser, __self__).__init__(
            'vault:okta/authBackendUser:AuthBackendUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, groups=None, path=None, policies=None, username=None):
        """
        Get an existing AuthBackendUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] groups: List of Okta groups to associate with this user
        :param pulumi.Input[str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[list] policies: List of Vault policies to associate with this user
        :param pulumi.Input[str] username: Name of the user within Okta

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_user.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["groups"] = groups
        __props__["path"] = path
        __props__["policies"] = policies
        __props__["username"] = username
        return AuthBackendUser(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

