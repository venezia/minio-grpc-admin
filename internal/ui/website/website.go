package website

import (
	"github.com/juju/loggo"
	"mime"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/venezia/minio-grpc-admin/pkg/generated/ui/data/homepage"
	"github.com/venezia/minio-grpc-admin/pkg/generated/ui/data/protobuf"
	"github.com/venezia/minio-grpc-admin/pkg/generated/ui/data/swagger"
	"github.com/venezia/minio-grpc-admin/pkg/generated/ui/data/swaggerjson"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "internal.ui.website"
)

var (
	logger loggo.Logger
)

func init() {
	logger = log.GetModuleLogger(loggerModuleName)
}

func AddWebsiteHandles(mux *http.ServeMux) {
	serveSwagger(mux)
	serveSwaggerJSON(mux)
	serveProtoBuf(mux)
	serveHomepage(mux)
}

func serveSwagger(mux *http.ServeMux) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		logger.Criticalf("could not add svg extension in file server")
	}

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveSwaggerJSON(mux *http.ServeMux) {
	err := mime.AddExtensionType(".json", "application/json")
	if err != nil {
		logger.Criticalf("could not add .json extension in file server")
	}

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swaggerjson.Asset,
		AssetDir: swaggerjson.AssetDir,
	})
	prefix := "/swagger/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveProtoBuf(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    protobuf.Asset,
		AssetDir: protobuf.AssetDir,
	})
	prefix := "/protobuf/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveHomepage(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    homepage.Asset,
		AssetDir: homepage.AssetDir,
	})
	prefix := "/"
	mux.Handle(prefix, http.StripPrefix("/", fileServer))
}
