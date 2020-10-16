package authorization

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/juju/loggo"
	"google.golang.org/grpc"

	"github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/constants"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "internal.apiserver.grpc.interceptors.authorization"
)

var (
	logger loggo.Logger
)

func init() {
	logger = log.GetModuleLogger(loggerModuleName)
}

func UnaryAuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error

	username := metautils.ExtractIncoming(ctx).Get(constants.AuthorizedUserHeader)
	groups := strings.Split(metautils.ExtractIncoming(ctx).Get(constants.AuthorizedGroupsHeader), ",")

	logger.Debugf(info.FullMethod)
	switch info.FullMethod {
	case constants.APIRootPath + "GetInstances":
		if username == "" || len(groups) < 1 {
			return nil, status.Error(codes.PermissionDenied, "not a valid user")
		}
	}

	h, err := handler(ctx, req)
	return h, err
}
