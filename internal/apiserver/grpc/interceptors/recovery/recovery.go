package recovery

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

//const (
//	loggerModuleName = "internal.apiserver.grpc.interceptors.recovery"
//)
//
//var (
//	logger loggo.Logger
//)
//
//func init() {
//	logger = log.GetModuleLogger(loggerModuleName)
//}

func RecoveryUnaryFunction() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor()
}

func RecoveryStreamFunction() grpc.StreamServerInterceptor {
	return grpc_recovery.StreamServerInterceptor()
}
