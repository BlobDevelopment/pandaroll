FROM golang:1.19-alpine

ENV MIGRATIONS_DIRECTORY=/migrations

WORKDIR /usr/src/app

COPY .git ./.git
COPY cmd ./cmd
COPY internal ./internal
COPY vendor ./vendor
COPY main.go go.mod go.sum Makefile ./

# Build
RUN apk add make git && \
    make build && \
    # Cleanup
    rm -rf cmd internal vendor

ENTRYPOINT [ "./bin/pandaroll" ]
CMD [ "migrate" ]
