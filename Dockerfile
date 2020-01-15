FROM golang:1.13.1-alpine3.10 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && make all

FROM alpine:3.10

COPY --from=builder /go/src/ex-sync /usr/local/bin

CMD irita-sync
