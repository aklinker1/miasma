#!/bin/bash
set -e

echo 'Generate CLI help docs...'
go run cmd/cli-docs/main.go > docs/docs/cli/Usage.md

echo 'Generate Server swagger docs...'
echo -e "---\nid: swagger\ntitle: Server API\nslug: /server\n---\n\nTODO" > docs/docs/server/Swagger.md
