#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..

echo "Ensuring Bin Directory (${PROJECT_DIRECTORY}/bin) Exists..."
mkdir -p "${PROJECT_DIRECTORY}/bin"
BIN_DIRECTORY=$(cd "${PROJECT_DIRECTORY}/bin" && pwd)
export PATH="${BIN_DIRECTORY}:${PATH}"

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${BIN_DIRECTORY} v1.27.0

echo
echo "running unit tests ..."
echo "golangci-lint run"
golangci-lint run
