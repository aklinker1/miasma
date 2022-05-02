#!/bin/bash
set -e
source scripts/build-config.sh

go build \
    -ldflags "-X $BUILD_VAR_PATH.VERSION=$VERSION -X $BUILD_VAR_PATH.BUILD=$BUILD -X $BUILD_VAR_PATH.BUILD_HASH=$BUILD_HASH -X $BUILD_VAR_PATH.BUILD_DATE=$BUILD_DATE" \
    -o bin/miasma \
    cmd/cli/main.go
cp bin/miasma "$GOPATH/bin"
