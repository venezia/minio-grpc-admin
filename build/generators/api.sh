#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..
DESTINATION_LIBRARY=pkg/generated/api
SWAGGER_DESTINATION=assets/generated/swagger

echo "Ensuring Destination Directory ( ${DESTINATION_LIBRARY} ) Exists..."
mkdir -p ${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/minioadmin/v1
echo "Ensuring Swagger Asset Directory ( ${SWAGGER_DESTINATION} ) Exists..."
mkdir -p ${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}/minioadmin/v1

echo
echo "generating minioadmin/v1 api stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/minioadmin/v1/ -I${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger"
protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/minioadmin/v1/ -I${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger

echo
echo "generating minioadmin/v1 REST gateway stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/minioadmin/v1/"
protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/minioadmin/v1/

echo
echo "generating minioadmin/v1 swagger docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1 -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}/minioadmin/v1"
protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1 -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}/minioadmin/v1

echo
echo "generating minioadmin/v1 api docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1 -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md"
protoc ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1 -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party -I ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/third_party/protoc-gen-swagger ${PROJECT_DIRECTORY}/api/minio-grpc-admin-spec/api/minioadmin/v1/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md
