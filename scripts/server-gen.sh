#!/bin/bash
set -e

if [[ "$(which swagger)" == "" ]]; then 
    echo -e "Install swagger code gen tool: https://goswagger.io/install.html\n"
    exit 1
fi

mkdir -p internal/server/gen package/client package/models
rm -rf internal/server/gen/restapi/operations package/client/operations
swagger generate -q server \
    --name miasma \
    --spec ./api/swagger.yml \
    --target internal/server/gen \
    --struct-tags json,gorm \
    --model-package ../../../package/models \
    --exclude-main
swagger generate -q client \
    --name miasma \
    --spec ./api/swagger.yml \
    --target package \
    --existing-models github.com/aklinker1/miasma/package/models
