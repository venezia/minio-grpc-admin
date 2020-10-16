package apiserver

import (
	"github.com/juju/loggo"
	"github.com/soheilhy/cmux"

	service "github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/miniogrpcadmin"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "internal.apiserver"
)

var (
	logger loggo.Logger
)

func init() {
	logger = log.GetModuleLogger(loggerModuleName)
}

type ServerOptions struct {
	PortNumber int
}

func AddServersToMux(tcpMux cmux.CMux, options *ServerOptions) {
	addGRPCServer(tcpMux)
	addRestAndWebsite(tcpMux, options.PortNumber)
}

func newgRPCServiceServer() *service.Server {
	return new(service.Server)
}
