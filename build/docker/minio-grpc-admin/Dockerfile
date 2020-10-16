# Initial preparations
FROM golang:1.14-alpine
RUN apk update && apk add git make curl

# Let's add the code in
ADD . /go/src/github.com/venezia/minio-grpc-admin

# Build package in container
WORKDIR /go/src/github.com/venezia/minio-grpc-admin
RUN make -f build/Makefile container-build

# Now that we have build a static binary, we can stuff it in a bare-bones alpine image
FROM alpine
COPY --from=0 /go/src/github.com/venezia/minio-grpc-admin /minio-grpc-admin

ENTRYPOINT ["/minio-grpc-admin"]
