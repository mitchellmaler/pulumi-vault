// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/database_secret_backend_connection.html.markdown.
type SecretBackendConnection struct {
	pulumi.CustomResourceState

	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrOutput `pulumi:"cassandra"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapOutput `pulumi:"data"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrOutput `pulumi:"hana"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrOutput `pulumi:"mongodb"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrOutput `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrOutput `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrOutput `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrOutput `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrOutput `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrOutput `pulumi:"oracle"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrOutput `pulumi:"postgresql"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayOutput `pulumi:"rootRotationStatements"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrOutput `pulumi:"verifyConnection"`
}

// NewSecretBackendConnection registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConnection(ctx *pulumi.Context,
	name string, args *SecretBackendConnectionArgs, opts ...pulumi.ResourceOption) (*SecretBackendConnection, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil {
		args = &SecretBackendConnectionArgs{}
	}
	var resource SecretBackendConnection
	err := ctx.RegisterResource("vault:database/secretBackendConnection:SecretBackendConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConnection gets an existing SecretBackendConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConnectionState, opts ...pulumi.ResourceOption) (*SecretBackendConnection, error) {
	var resource SecretBackendConnection
	err := ctx.ReadResource("vault:database/secretBackendConnection:SecretBackendConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConnection resources.
type secretBackendConnectionState struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend *string `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra *SecretBackendConnectionCassandra `pulumi:"cassandra"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data map[string]interface{} `pulumi:"data"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana *SecretBackendConnectionHana `pulumi:"hana"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb *SecretBackendConnectionMongodb `pulumi:"mongodb"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql *SecretBackendConnectionMssql `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql *SecretBackendConnectionMysql `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora *SecretBackendConnectionMysqlAurora `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy *SecretBackendConnectionMysqlLegacy `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds *SecretBackendConnectionMysqlRds `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name *string `pulumi:"name"`
	// A nested block containing configuration options for Oracle connections.
	Oracle *SecretBackendConnectionOracle `pulumi:"oracle"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql *SecretBackendConnectionPostgresql `pulumi:"postgresql"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements []string `pulumi:"rootRotationStatements"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

type SecretBackendConnectionState struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayInput
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringPtrInput
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrInput
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapInput
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrInput
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrInput
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrInput
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrInput
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrInput
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrInput
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrInput
	// A unique name to give the database connection.
	Name pulumi.StringPtrInput
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrInput
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrInput
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayInput
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConnectionState)(nil)).Elem()
}

type secretBackendConnectionArgs struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The unique name of the Vault mount to configure.
	Backend string `pulumi:"backend"`
	// A nested block containing configuration options for Cassandra connections.
	Cassandra *SecretBackendConnectionCassandra `pulumi:"cassandra"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data map[string]interface{} `pulumi:"data"`
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana *SecretBackendConnectionHana `pulumi:"hana"`
	// A nested block containing configuration options for MongoDB connections.
	Mongodb *SecretBackendConnectionMongodb `pulumi:"mongodb"`
	// A nested block containing configuration options for MSSQL connections.
	Mssql *SecretBackendConnectionMssql `pulumi:"mssql"`
	// A nested block containing configuration options for MySQL connections.
	Mysql *SecretBackendConnectionMysql `pulumi:"mysql"`
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora *SecretBackendConnectionMysqlAurora `pulumi:"mysqlAurora"`
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy *SecretBackendConnectionMysqlLegacy `pulumi:"mysqlLegacy"`
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds *SecretBackendConnectionMysqlRds `pulumi:"mysqlRds"`
	// A unique name to give the database connection.
	Name *string `pulumi:"name"`
	// A nested block containing configuration options for Oracle connections.
	Oracle *SecretBackendConnectionOracle `pulumi:"oracle"`
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql *SecretBackendConnectionPostgresql `pulumi:"postgresql"`
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements []string `pulumi:"rootRotationStatements"`
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

// The set of arguments for constructing a SecretBackendConnection resource.
type SecretBackendConnectionArgs struct {
	// A list of roles that are allowed to use this
	// connection.
	AllowedRoles pulumi.StringArrayInput
	// The unique name of the Vault mount to configure.
	Backend pulumi.StringInput
	// A nested block containing configuration options for Cassandra connections.
	Cassandra SecretBackendConnectionCassandraPtrInput
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	Data pulumi.MapInput
	// A nested block containing configuration options for SAP HanaDB connections.
	Hana SecretBackendConnectionHanaPtrInput
	// A nested block containing configuration options for MongoDB connections.
	Mongodb SecretBackendConnectionMongodbPtrInput
	// A nested block containing configuration options for MSSQL connections.
	Mssql SecretBackendConnectionMssqlPtrInput
	// A nested block containing configuration options for MySQL connections.
	Mysql SecretBackendConnectionMysqlPtrInput
	// A nested block containing configuration options for Aurora MySQL connections.
	MysqlAurora SecretBackendConnectionMysqlAuroraPtrInput
	// A nested block containing configuration options for legacy MySQL connections.
	MysqlLegacy SecretBackendConnectionMysqlLegacyPtrInput
	// A nested block containing configuration options for RDS MySQL connections.
	MysqlRds SecretBackendConnectionMysqlRdsPtrInput
	// A unique name to give the database connection.
	Name pulumi.StringPtrInput
	// A nested block containing configuration options for Oracle connections.
	Oracle SecretBackendConnectionOraclePtrInput
	// A nested block containing configuration options for PostgreSQL connections.
	Postgresql SecretBackendConnectionPostgresqlPtrInput
	// A list of database statements to be executed to rotate the root user's credentials.
	RootRotationStatements pulumi.StringArrayInput
	// Whether the connection should be verified on
	// initial configuration or not.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConnectionArgs)(nil)).Elem()
}

