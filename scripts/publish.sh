#!/bin/bash
set -e

docker login
docker buildx build . \
    -f docker/Dockerfile.prod \
    --push \
    --platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
    --tag aklinker1/miasma:nightly \
