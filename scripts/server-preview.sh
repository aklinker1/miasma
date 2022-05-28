#!/bin/bash
set -e
source scripts/build-config.sh

./scripts/server-build.sh
docker run \
    -i \
    --rm \
    -p 3000:3000 \
    -v "$(pwd)/data":/data/miasma \
    -v /var/run/docker.sock:/var/run/docker.sock \
    aklinker1/miasma:local
