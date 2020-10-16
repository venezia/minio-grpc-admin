GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install

# Variables for binary version information
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# Associating version information variables with data
VERSION_PACKAGE=gitlab.com/mvenezia/version-info/pkg/version
LD_DATA_GIT_COMMIT=-X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT}
LD_DATA_GIT_TAG=-X ${VERSION_PACKAGE}.gitVersion=${GIT_TAG}
LD_DATA_GIT_DIRTY=-X ${VERSION_PACKAGE}.gitTreeState=${GIT_DIRTY}
LD_DATA_BUILD_DATE=-X ${VERSION_PACKAGE}.buildDate=${BUILD_DATE}
LD_DATA=${LD_DATA_GIT_COMMIT} ${LD_DATA_GIT_TAG} ${LD_DATA_GIT_DIRTY} ${LD_DATA_BUILD_DATE}

# Creating static binary
STATIC_FLAGS=-w -extldflags "-static"

# Generating ldflag string
LDFLAGS = '${STATIC_FLAGS} ${LD_DATA}'

# Location of
MGA_BINARY_NAME=minio-grpc-admin
MGA_SOURCE_FILE=cmd/minio-grpc-admin/main.go

# build binary command
BUILD_MGA_BINARY_CMD=${GOBUILD} -a -tags netgo -ldflags ${LDFLAGS} -o ${MGA_BINARY_NAME} ${MGA_SOURCE_FILE}

# By default we will most likely want to build a linux build of the application
all: mga-bin-linux-amd64

container-build: mga-bin-linux-amd64

# Clean will clean up dep and remove any binaries
clean: remove-binary

# Build a darwin (mac) binary
mga-bin-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${BUILD_MGA_BINARY_CMD}

# Build a linux amd64 binary
mga-bin-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${BUILD_MGA_BINARY_CMD}

# Delete an existing binary / cleanup
remove-binary:
	rm ${MGA_BINARY_NAME}

regenerate-api:
	./build/generators/api.sh
	./build/generators/bindata.sh

# Run tests over all the Go code and display code coverage
.PHONY: test-cover
test-cover:
	./build/scripts/goUnitTest.sh

# Run a linting over all the Go code in the repository
.PHONY: lint
lint:
	./build/scripts/goCiLint.sh