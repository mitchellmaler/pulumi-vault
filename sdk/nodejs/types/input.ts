// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";

export interface AuthBackendTune {
    allowedResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
    auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
    defaultLeaseTtl?: pulumi.Input<string>;
    /**
     * Speficies whether to show this mount in the UI-specific listing endpoint.
     */
    listingVisibility?: pulumi.Input<string>;
    maxLeaseTtl?: pulumi.Input<string>;
    passthroughRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    tokenType?: pulumi.Input<string>;
}

export interface GetPolicyDocumentRule {
    /**
     * Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
     */
    allowedParameters?: inputs.GetPolicyDocumentRuleAllowedParameter[];
    /**
     * A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
     */
    capabilities: string[];
    /**
     * Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
     */
    deniedParameters?: inputs.GetPolicyDocumentRuleDeniedParameter[];
    /**
     * Description of the rule. Will be added as a commend to rendered rule.
     */
    description?: string;
    /**
     * The maximum allowed TTL that clients can specify for a wrapped response.
     */
    maxWrappingTtl?: string;
    /**
     * The minimum allowed TTL that clients can specify for a wrapped response.
     */
    minWrappingTtl?: string;
    /**
     * A path in Vault that this rule applies to.
     */
    path: string;
    /**
     * A list of parameters that must be specified.
     */
    requiredParameters?: string[];
}

export interface GetPolicyDocumentRuleAllowedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export interface GetPolicyDocumentRuleDeniedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export interface ProviderAuthLogin {
    namespace?: pulumi.Input<string>;
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    path: pulumi.Input<string>;
}

export interface ProviderClientAuth {
    certFile: pulumi.Input<string>;
    keyFile: pulumi.Input<string>;
}

export namespace azure {
    export interface BackendRoleAzureRole {
        roleId?: pulumi.Input<string>;
        roleName: pulumi.Input<string>;
        scope: pulumi.Input<string>;
    }
}

export namespace database {
    export interface SecretBackendConnectionCassandra {
        /**
         * The number of seconds to use as a connection
         * timeout.
         */
        connectTimeout?: pulumi.Input<number>;
        /**
         * The hosts to connect to.
         */
        hosts?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Whether to skip verification of the server
         * certificate when using TLS.
         */
        insecureTls?: pulumi.Input<boolean>;
        /**
         * The password to authenticate with.
         */
        password?: pulumi.Input<string>;
        /**
         * Concatenated PEM blocks configuring the certificate
         * chain.
         */
        pemBundle?: pulumi.Input<string>;
        /**
         * A JSON structure configuring the certificate chain.
         */
        pemJson?: pulumi.Input<string>;
        /**
         * The default port to connect to if no port is specified as
         * part of the host.
         */
        port?: pulumi.Input<number>;
        /**
         * The CQL protocol version to use.
         */
        protocolVersion?: pulumi.Input<number>;
        /**
         * Whether to use TLS when connecting to Cassandra.
         */
        tls?: pulumi.Input<boolean>;
        /**
         * The username to authenticate with.
         */
        username?: pulumi.Input<string>;
    }

    export interface SecretBackendConnectionHana {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMongodb {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMssql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMysql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMysqlAurora {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMysqlLegacy {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionMysqlRds {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionOracle {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }

    export interface SecretBackendConnectionPostgresql {
        /**
         * A URL containing connection information. See
         * the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/oracle.html#sample-payload)
         * for an example.
         */
        connectionUrl?: pulumi.Input<string>;
        /**
         * The maximum number of seconds to keep
         * a connection alive for.
         */
        maxConnectionLifetime?: pulumi.Input<number>;
        /**
         * The maximum number of idle connections to
         * maintain.
         */
        maxIdleConnections?: pulumi.Input<number>;
        /**
         * The maximum number of open connections to
         * use.
         */
        maxOpenConnections?: pulumi.Input<number>;
    }
}

export namespace gcp {
    export interface SecretRolesetBinding {
        resource: pulumi.Input<string>;
        roles: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace github {
    export interface AuthBackendTune {
        allowedResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
        auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
        defaultLeaseTtl?: pulumi.Input<string>;
        listingVisibility?: pulumi.Input<string>;
        maxLeaseTtl?: pulumi.Input<string>;
        passthroughRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * (Optional) The type of token that should be generated. Can be `service`,
         * `batch`, or `default` to use the mount's tuned default (which unless changed will be
         * `service` tokens). For token store roles, there are two additional possibilities:
         * `default-service` and `default-batch` which specify the type to return unless the client
         * requests a different type at generation time.
         */
        tokenType?: pulumi.Input<string>;
    }
}

export namespace identity {
}

export namespace jwt {
    export interface AuthBackendTune {
        allowedResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        auditNonHmacRequestKeys?: pulumi.Input<pulumi.Input<string>[]>;
        auditNonHmacResponseKeys?: pulumi.Input<pulumi.Input<string>[]>;
        defaultLeaseTtl?: pulumi.Input<string>;
        listingVisibility?: pulumi.Input<string>;
        maxLeaseTtl?: pulumi.Input<string>;
        passthroughRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
        tokenType?: pulumi.Input<string>;
    }
}

export namespace okta {
    export interface AuthBackendGroup {
        /**
         * Name of the group within the Okta
         */
        groupName: pulumi.Input<string>;
        /**
         * List of Vault policies to associate with this user
         */
        policies: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface AuthBackendUser {
        /**
         * List of Okta groups to associate with this user
         */
        groups: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * List of Vault policies to associate with this user
         */
        policies?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Name of the user within Okta
         */
        username: pulumi.Input<string>;
    }
}

export namespace rabbitMq {
    export interface SecretBackendRoleVhost {
        configure: pulumi.Input<string>;
        host: pulumi.Input<string>;
        read: pulumi.Input<string>;
        write: pulumi.Input<string>;
    }
}
