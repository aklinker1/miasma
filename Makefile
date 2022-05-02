VERSION := $(shell jq -r .version package.json)
BUILD := $(shell TZ=UTC git --no-pager show --quiet --abbrev=40 --format='%h')
BUILD_HASH := $(shell TZ=UTC git --no-pager show --quiet --abbrev=8 --format='%h')
BUILD_DATE := $(shell TZ=UTC git --no-pager show --quiet --date='format-local:%Y%m%d%H%M%S' --format='%cd')
BUILD_VAR_PATH := github.com/aklinker1/miasma/internal/shared/constants

# Server

build:
	docker build . -f docker/Dockerfile.prod \
		-t aklinker1/miasma:dev \
		--build-arg VERSION="$(VERSION)" \
		--build-arg BUILD="$(BUILD)" \
		--build-arg BUILD_HASH="$(BUILD_HASH)" \
		--build-arg BUILD_DATE="$(BUILD_DATE)"
prep:
	@mkdir -p web/dist
	@echo "<html></html>" > web/dist/index.html

run: prep
	VERSION=$(VERSION) docker-compose up --build -V


swagger:
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
publish:
	docker login
	docker buildx build \
		-f docker/Dockerfile.prod \
		--push \
		--platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
		--tag aklinker1/miasma:nightly \
		.

# CLI

cli:
	go build \
		-ldflags "-X ${BUILD_VAR_PATH}.VERSION=${VERSION} -X ${BUILD_VAR_PATH}.BUILD=${BUILD} -X ${BUILD_VAR_PATH}.BUILD_HASH=${BUILD_HASH} -X ${BUILD_VAR_PATH}.BUILD_DATE=${BUILD_DATE}" \
		-o bin/miasma \
		cmd/cli/main.go
	cp bin/miasma ${GOPATH}/bin

# Docs

start-docs:
	yarn --cwd docs start --port 8888
build-cli-docs:
	@echo 'Generate CLI Docs...'
	go build -o bin/cli-docs cmd/cli-docs/main.go && ./bin/cli-docs > docs/docs/cli/Usage.md
build-server-docs:
	@echo 'Generate Server Swagger Docs...'
	@echo "---\nid: swagger\ntitle: Server API\nslug: /server\n---\n\nTODO" > docs/docs/server/Swagger.md
build-docs-site:
	@echo 'Building docs site...'
	yarn --cwd docs build
build-docs: build-cli-docs build-server-docs build-docs-site
publish-docs: build-docs
	yarn --cwd docs deploy

#  Aliases

b: build
r: run
w: watch
s: swagger
p: publish
sr: swagger run

c: cli
d: docs
