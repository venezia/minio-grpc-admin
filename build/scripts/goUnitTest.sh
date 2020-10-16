#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

DIRS="./internal/...
./cmd/..."

echo
echo "running unit tests ..."
echo 'echo ${DIRS} | while read -r line; do GO111MODULE=on go test -cover $line; done'
echo ${DIRS} | while read -r line; do GO111MODULE=on go test -cover $line; done
