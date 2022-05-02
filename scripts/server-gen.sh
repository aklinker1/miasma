#!/bin/bash
set -e

if [[ "$(which swagger)" == "" ]]; then 
    echo -e "Install swagger code gen tool: https://goswagger.io/install.html\n"
    exit 1
fi

mkdir -p internal/miasma-server/swagger package/client

swagger generate -q server \
    --name miasma \
    --spec ./api/swagger.yml \
    --target ../../internal/miasma-server/swagger \
    --struct-tags json,gorm \
    --model-package github.com/aklinker1/miasma/internal \
    --exclude-main

# swagger generate -q client \
#     --name miasma \
#     --spec ./api/swagger.yml \
#     --target internal/miasma-client \
#     --existing-models github.com/aklinker1/miasma/internal
