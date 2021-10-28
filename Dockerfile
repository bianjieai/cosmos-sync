FROM golang:1.14-alpine as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@gitlab.bianjie.ai/cschain/cschain".insteadOf "https://gitlab.bianjie.ai/cschain/cschain" && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@github.com/bianjieai/iritamod".insteadOf "https://github.com/bianjieai/iritamod" && make all

FROM alpine:3.10

COPY --from=builder /go/src/cosmos-sync /usr/local/bin

CMD cosmos-sync
