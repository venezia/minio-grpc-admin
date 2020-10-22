package minio

import (
	"github.com/juju/loggo"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "internal.minio"
)

var (
	logger loggo.Logger
)

func init() {
	logger = log.GetModuleLogger(loggerModuleName)
}
