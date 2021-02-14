VERSION := $(shell jq -r .version meta.json)
BUILD := $(shell TZ=UTC git --no-pager show --quiet --abbrev=40 --format='%h')
BUILD_8 := $(shell TZ=UTC git --no-pager show --quiet --abbrev=8 --format='%h')
BUILD_DATE := $(shell TZ=UTC git --no-pager show --quiet --date='format-local:%Y%m%d%H%M%S' --format='%cd')
BUILD_VAR_PATH := github.com/aklinker1/miasma/internal/shared/constants

# Server

build:
	docker build . -f docker/Dockerfile.server \
		-t aklinker1/miasma:dev \
		--build-arg VERSION="$(VERSION)" \
		--build-arg BUILD="$(BUILD)" \
		--build-arg BUILD_8="$(BUILD_8)" \
		--build-arg BUILD_DATE="$(BUILD_DATE)"
run: build
	@echo ""
	@echo "---"
	@echo ""
	docker run -i --env-file .env -p 3000:3000 -v "$(shell pwd)"/data:/data/miasma -v /var/run/docker.sock:/var/run/docker.sock aklinker1/miasma:dev
watch:
	@modd
swagger:
	swagger generate server -t internal/server/gen -f ./api/swagger.yml --exclude-main -A miasma
	swagger generate client -t package -f ./api/swagger.yml -A miasma
publish:
	docker login
	docker buildx build \
		-f docker/Dockerfile.server \
		--push \
		--platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
		--tag aklinker1/miasma:nightly .

# CLI

cli:
	go build \
		-ldflags "-X ${BUILD_VAR_PATH}.VERSION=${VERSION} -X ${BUILD_VAR_PATH}.BUILD=${BUILD} -X ${BUILD_VAR_PATH}.BUILD_8=${BUILD_8} -X ${BUILD_VAR_PATH}.BUILD_DATE=${BUILD_DATE}" \
		-o bin/miasma \
		cmd/cli/main.go
	cp bin/miasma ${GOPATH}/bin

# Docs

start-docs:
	yarn --cwd web/docs start --port 8888
build-cli-docs:
	@echo 'Generate CLI Docs...'
	go build -o bin/cli-docs cmd/cli-docs/main.go && ./bin/cli-docs > web/docs/docs/cli/Usage.md
build-server-docs:
build-docs-site:
	@echo 'Building docs site...'
	yarn --cwd web/docs build
build-docs: build-cli-docs, build-server-docs, build-docs-site
publish-docs: build-docs
	yarn --cwd web/docs deploy

#  Aliases

b: build
r: run
w: watch
s: swagger
p: publish

c: cli
d: docs
