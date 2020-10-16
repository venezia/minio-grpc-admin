package authentication

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/venezia/minio-grpc-admin/internal/apiserver/grpc/constants"
	"github.com/venezia/minio-grpc-admin/internal/flags"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "internal.apiserver.grpc.interceptors.authentication"
)

var (
	logger loggo.Logger
)

func init() {
	logger = log.GetModuleLogger(loggerModuleName)
}

func UnaryAuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error

	// If we're not enabled, we aren't enabled... bail!
	if !viper.GetBool(flags.GRPCAuthentication) {
		logger.Debugf("authentication is not enabled, skipping!")
		return handler(ctx, req)
	}

	// If we're being used, for security, we should reset the user and groups information
	resetUserAndGroups(ctx)

	// Let's challenge that token
	err = challengeToken(ctx)
	if err != nil {
		return nil, err
	}

	logger.Debugf("User is -->%s<--, and is part of groups -->%s<--", metautils.ExtractIncoming(ctx).Get(constants.AuthorizedUserHeader), metautils.ExtractIncoming(ctx).Get(constants.AuthorizedGroupsHeader))
	h, err := handler(ctx, req)
	return h, err
}

func challengeToken(ctx context.Context) error {
	authorizationArray := strings.Split(metautils.ExtractIncoming(ctx).Get(constants.AuthorizationHeader), " ")
	if len(authorizationArray) != 2 {
		return status.Error(codes.Unauthenticated, "not really allowed")
	}
	token := authorizationArray[1]
	logger.Debugf(token)

	user, groups, err := verifyOIDCToken(ctx, token)
	if err != nil {
		return status.Error(codes.Unauthenticated, "not allowed")
	}

	setUserAndGroups(ctx, user, groups)
	return nil
}

func resetUserAndGroups(ctx context.Context) {
	metautils.ExtractIncoming(ctx).Set(constants.AuthorizedUserHeader, "")
	metautils.ExtractIncoming(ctx).Set(constants.AuthorizedGroupsHeader, "")
}

func setUserAndGroups(ctx context.Context, user string, groups []string) {
	metautils.ExtractIncoming(ctx).Set(constants.AuthorizedUserHeader, user)
	metautils.ExtractIncoming(ctx).Set(constants.AuthorizedGroupsHeader, strings.Join(groups, ","))
}
