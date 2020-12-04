FROM golang:1.13.1-alpine3.10 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@gitlab.bianjie.ai/cschain/cschain".insteadOf "https://gitlab.bianjie.ai/cschain/cschain" && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@github.com/bianjieai/iritamod".insteadOf "https://github.com/bianjieai/iritamod" && make all

FROM alpine:3.10

COPY --from=builder /go/src/irita-sync /usr/local/bin

CMD irita-sync
