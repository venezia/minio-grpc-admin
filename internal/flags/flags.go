package flags

import "github.com/juju/loggo"

const (
	EnvironmentPrefix = "miniogrpc"

	GRPCAuthentication = "grpc-authentication"
	GRPCAuthorization  = "grpc-authorization"
	GRPCUI             = "grpcui"
	GRPCUIPort         = "grpcui-port"
	GRPCRecovery       = "grpc-recovery"
	GRPCReflection     = "grpc-reflection"
	LogLevel           = "log-level"
	OIDCAudience       = "oidc-audience"
	OIDCEndpoint       = "oidc-endpoint"
	Port               = "port"

	GRPCAuthenticationDefault = false
	GRPCAuthorizationDefault  = true
	GRPCUIDefault             = false
	GRPCUIPortDefault         = 19050
	GRPCRecoveryDefault       = true
	GRPCReflectionDefault     = true
	LogLevelDefault           = "INFO"
	OIDCAudienceDefault       = "minio-grpc-admin"
	OIDCEndpointDefault       = ""
	PortDefault               = 9050

	GRPCAuthenticationDescription = "Should grpc validate the identity of the client"
	GRPCAuthorizationDescription  = "Should grpc validate if the client can request the rpc"
	GRPCUIDescription             = "Should gRPC-UI be enabled"
	GRPCUIPortDescription         = "Port to have grpcui listen on"
	GRPCRecoveryDescription       = "Should grpc recover from panics"
	GRPCReflectionDescription     = "Should grpc server have reflection enabled"
	LogLevelDescription           = "Log verbosity, values are: CRITICAL, ERROR, WARNING, INFO, DEBUG, TRACE"
	OIDCAudienceDescription       = "What Audience to require for use with OIDC identity tokens"
	OIDCEndpointDescription       = "What OIDC Provider to use for authentication"
	PortDescription               = "Port to listen on"
)

var (
	RealLogLevel loggo.Level
)
