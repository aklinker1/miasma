# Setup base images
FROM alpine as base-image
RUN mkdir -p /data/miasma
WORKDIR /app
EXPOSE 3000

FROM node:14-alpine as vue-builder-base
RUN mkdir /build
WORKDIR /build

FROM golang:1.14-alpine as go-builder-base
RUN apk update
RUN apk add git jq
RUN mkdir /build
WORKDIR /build


# Build the dashboard
FROM vue-builder-base as vue-builder
COPY web/dashboard .
# yarn build


# Build the server
FROM go-builder-base as go-builder
COPY go.mod go.sum ./
RUN go mod download
COPY meta.json .
COPY .git ./.git
COPY cmd/server ./cmd/server
COPY internal/shared ./internal/shared
COPY internal/server ./internal/server
RUN \
  VERSION="$(jq -r .version meta.json)-$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')" ;\
  go build \
    -ldflags "-X github.com/aklinker1/miasma/internal/server/utils/constants.VERSION=$VERSION" \
    -o bin/server \
    cmd/server/main.go


# Make the final image with just the docker cli, the server's go binary, and dashboard UI
FROM base-image
ENV \
  DOCKER_HOST="unix:///var/run/docker.sock"
COPY --from=go-builder /build/bin/server .
COPY --from=vue-builder /build/dist dashboard
CMD ["./server"]
