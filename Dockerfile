FROM golang:1.14-alpine as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && make all

FROM alpine:3.10

COPY --from=builder /go/src/irita-sync /usr/local/bin

CMD irita-sync
