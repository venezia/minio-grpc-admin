package apiserver

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/interceptors/authentication"
	"github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/interceptors/authorization"
	"github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/interceptors/recovery"
	mav1pb "github.com/venezia/minio-grpc-admin/pkg/generated/api/minioadmin/v1"
	"github.com/venezia/minio-grpc-admin/internal/flags"
)

var (
	chainedUnaryInterceptors  []grpc.UnaryServerInterceptor
	chainedStreamInterceptors []grpc.StreamServerInterceptor
)

func addGRPCServer(tcpMux cmux.CMux) {
	grpcServer := newGRPCInstance()

	grpcListener := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	// Start servers
	go func() {
		logger.Infof("Starting gRPC Server")
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Criticalf("Unable to start external gRPC server")
		}
	}()
}

func newGRPCInstance() *grpc.Server {
	// Order matters...
	addPanicRecoveryInterceptors()
	addAuthenticationInterceptors()
	addAuthorizationInterceptors()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(chainedUnaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(chainedStreamInterceptors...)),
	}
	grpcServer := grpc.NewServer(opts...)
	mav1pb.RegisterMinioadminServer(grpcServer, newgRPCServiceServer())
	enableGRPCReflection(grpcServer)

	return grpcServer
}

func addPanicRecoveryInterceptors() {
	if viper.GetBool(flags.GRPCRecovery) {
		logger.Infof("enabling grpc recovery")
		chainedUnaryInterceptors = append(chainedUnaryInterceptors, recovery.RecoveryUnaryFunction())
		chainedStreamInterceptors = append(chainedStreamInterceptors, recovery.RecoveryStreamFunction())
	} else {
		logger.Infof("disabling grpc recovery")
	}
}

func addAuthenticationInterceptors() {
	if viper.GetBool(flags.GRPCRecovery) {
		logger.Infof("enabling grpc authentication")
		chainedUnaryInterceptors = append(chainedUnaryInterceptors, authentication.UnaryAuthenticationInterceptor)
	} else {
		logger.Infof("disabling grpc authentication")
	}
}

func addAuthorizationInterceptors() {
	if viper.GetBool(flags.GRPCRecovery) {
		logger.Infof("enabling grpc authorization")
		chainedUnaryInterceptors = append(chainedUnaryInterceptors, authorization.UnaryAuthorizationInterceptor)
	} else {
		logger.Infof("disabling grpc authorization")
	}
}

func enableGRPCReflection(server *grpc.Server) {
	reflection.Register(server)
}
