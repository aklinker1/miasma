#!/bin/bash
set -e
source scripts/build-config.sh

./scripts/build.sh
docker run \
    -i \
    --env-file .env \
    -p 3000:3000 \
    -v "$(pwd)/data":/data/miasma \
    -v /var/run/docker.sock:/var/run/docker.sock \
    aklinker1/miasma:local
