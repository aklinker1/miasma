#!/bin/bash
GEN_PATH="src/utils/api-gen/"
openapi-generator-cli generate \
    -i ../api/swagger.yml \
    -g typescript-axios \
    -o "$GEN_PATH"

# Cleanup
pushd "$GEN_PATH"
rm -rf \
    .openapi-generator \
    .openapi-generator-ignore \
    .npmignore \
    git_push.sh \
    .gitignore
popd

yarn prettier --write "$GEN_PATH"
