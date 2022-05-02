#!/bin/bash
set -e
source scripts/build-config.sh

docker build -q . -f Dockerfile.dev \
    -t aklinker1/miasma:dev \
    --build-arg VERSION="$VERSION" \
    --build-arg BUILD="$BUILD" \
    --build-arg BUILD_HASH="$BUILD_HASH" \
    --build-arg BUILD_DATE="$BUILD_DATE" \
    --build-arg BUILD_VAR_PATH="$BUILD_VAR_PATH"

docker run \
    -i \
    --rm \
    --env-file .env \
    -p 3001:3001 \
    -v "$(pwd)/data":/data/miasma \
    -v /var/run/docker.sock:/var/run/docker.sock \
    --name miasma-dev \
    aklinker1/miasma:dev
