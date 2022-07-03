# Setup base images
FROM alpine as base-image
RUN mkdir -p /data/miasma
WORKDIR /app

FROM node:16-alpine as web-builder-base
RUN apk add --update curl
RUN mkdir /build
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm@7
WORKDIR /build

FROM golang:1.18-alpine as api-builder-base
RUN apk add --update git jq build-base
RUN mkdir /build
WORKDIR /build


# Build the dashboard
FROM web-builder-base as web-builder
COPY web/package.json package.json
COPY web/pnpm-lock.yaml pnpm-lock.yaml
RUN pnpm install --frozen-lockfile
COPY web ./
RUN pnpm build

# Build the server
FROM api-builder-base as api-builder
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/mattn/go-sqlite3
COPY cmd/server ./cmd/server
COPY web/fs.go web/fs.go
COPY internal internal
COPY --from=web-builder /build/dist web/dist
ARG VERSION
ARG BUILD
ARG BUILD_HASH
ARG BUILD_DATE
ARG BUILD_VAR_PATH
RUN go build \
  -ldflags "-X $BUILD_VAR_PATH.VERSION=$VERSION -X $BUILD_VAR_PATH.BUILD=$BUILD -X $BUILD_VAR_PATH.BUILD_HASH=$BUILD_HASH -X $BUILD_VAR_PATH.BUILD_DATE=$BUILD_DATE" \
  -o bin/server \
  cmd/server/main.go

# Make the final image with just the docker cli, the server's go binary, and dashboard UI
FROM base-image
ENV DOCKER_HOST="unix:///var/run/docker.sock"
COPY --from=api-builder /build/bin/server .
ENTRYPOINT [ "./server" ]
