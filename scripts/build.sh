#!/bin/bash
set -e
source scripts/build-config.sh

docker build . -f docker/Dockerfile.prod \
    -t aklinker1/miasma:dev \
    --build-arg VERSION="$VERSION" \
    --build-arg BUILD="$BUILD" \
    --build-arg BUILD_HASH="$BUILD_HASH" \
    --build-arg BUILD_DATE="$BUILD_DATE" \
    --build-arg BUILD_VAR_PATH="$BUILD_VAR_PATH"
