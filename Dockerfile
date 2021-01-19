# Build the dashboard
FROM node:14-alpine as vue-builder
RUN mkdir /build
WORKDIR /build
COPY web/dashboard .
# yarn build

# Build the server
FROM golang:1.14-alpine as go-builder
RUN apk update
RUN apk add git jq
RUN mkdir /build
WORKDIR /build
# Cache layer for dependencies
COPY go.mod go.sum ./
RUN go mod download
# Cached layer for source code
COPY .git ./.git
COPY cmd/server ./cmd/server
COPY internal/server ./internal/server
COPY meta.json .
RUN \
  VERSION="$(jq -r .version meta.json)-$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')" ;\
  go build \
    -ldflags "-X github.com/aklinker1/miasma/internal/server/utils/constants.VERSION=$VERSION" \
    -o bin/server \
    cmd/server/main.go

# Make the final image with just the go binary and dashboard dist
FROM alpine
RUN adduser -S -D -H -h /app appuser
RUN mkdir -p /data/miasma
RUN chown -R appuser /data/miasma
USER appuser
WORKDIR /app
COPY --from=go-builder /build/bin/server .
COPY --from=vue-builder /build/dist dashboard
EXPOSE 3000
CMD ["./server"]
