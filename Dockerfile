# syntax = docker/dockerfile:1-experimental
FROM golang:1.19-alpine3.16 AS base

RUN apk --update upgrade && apk --no-cache --update-cache --upgrade --latest add ca-certificates build-base gcc

WORKDIR /go/src/github.com/ory/kratos

ADD go.mod go.mod
ADD go.sum go.sum
ADD internal/httpclient/go.* internal/httpclient/
ADD internal/client-go/go.* internal/client-go/

ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV CGO_CPPFLAGS -DSQLITE_DEFAULT_FILE_PERMISSIONS=0600

RUN go mod download

ADD . .

ARG VERSION
ARG COMMIT
ARG BUILD_DATE

RUN --mount=type=cache,target=/root/.cache/go-build go build -tags sqlite \
    -ldflags="-X 'github.com/ory/kratos/driver/config.Version=${VERSION}' -X 'github.com/ory/kratos/driver/config.Date=${BUILD_DATE}' -X 'github.com/ory/kratos/driver/config.Commit=${COMMIT}'" \
    -o /usr/bin/kratos

FROM alpine:3.16

RUN addgroup -S ory; \
    adduser -S ory -G ory -D -u 10000 -h /home/ory -s /bin/nologin; \
    chown -R ory:ory /home/ory

COPY --from=base /usr/bin/kratos /usr/bin/kratos
COPY econt.jsonnet /usr/bin/econt.jsonnet

# By creating the sqlite folder as the ory user, the mounted volume will be owned by ory:ory, which
# is required for read/write of SQLite.
RUN mkdir -p /var/lib/sqlite
RUN chown ory:ory /var/lib/sqlite
VOLUME /var/lib/sqlite

# Exposing the ory home directory to simplify passing in Kratos configuration (e.g. if the file $HOME/.kratos.yaml
# exists, it will be automatically used as the configuration file).
VOLUME /home/ory

# Declare the standard ports used by Kratos (4433 for public service endpoint, 4434 for admin service endpoint)
EXPOSE 4433 4434

USER 10000

ENTRYPOINT ["kratos"]
CMD ["serve"]
