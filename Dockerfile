FROM golang:1.13.1-alpine3.10 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@gitlab.bianjie.ai/irita/irita".insteadOf "https://gitlab.bianjie.ai/irita/irita" && make all

FROM alpine:3.10

COPY --from=builder /go/src/rb-sync-binance /usr/local/bin

CMD ex-sync
