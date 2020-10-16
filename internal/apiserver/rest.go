package apiserver

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	"github.com/venezia/minio-grpc-admin/internal/ui/website"
	mav1pb "github.com/venezia/minio-grpc-admin/pkg/generated/api/minioadmin/v1"
)

func addRestAndWebsite(tcpMux cmux.CMux, grpcPortNumber int) {
	httpListener := tcpMux.Match(cmux.HTTP1Fast())

	go func() {
		router := http.NewServeMux()
		website.AddWebsiteHandles(router)
		addgRPCRestGateway(router, grpcPortNumber)
		httpServer := http.Server{
			Handler: addCORSHeader(router),
		}
		logger.Infof("Starting HTTP/1 Server")
		err := httpServer.Serve(httpListener)
		if err != nil {
			logger.Criticalf("could not start HTTP/1 server, error was: %s", err.Error())
		}
	}()

}

func addgRPCRestGateway(router *http.ServeMux, grpcPortNumber int) {
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	err := mav1pb.RegisterMinioadminHandlerFromEndpoint(context.Background(), gwmux, "localhost:"+strconv.Itoa(grpcPortNumber), dopts)
	if err != nil {
		logger.Criticalf("could not register rest gateway - error was: %s", err.Error())
	}
	router.Handle("/api/", gwmux)
}

func addCORSHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
