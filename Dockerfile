FROM golang:1.16-alpine as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash
ENV GO111MODULE  on

ARG GITUSER=bamboo
ARG GITPASS=FS_Q5LmxwExwK6hFN9Fs
ARG GOPRIVATE=gitlab.bianjie.ai
ARG GOPROXY=https://goproxy.cn,direct
ARG APKPROXY=http://mirrors.ustc.edu.cn/alpine

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache $PACKAGES && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@gitlab.bianjie.ai/cschain/cschain".insteadOf "https://gitlab.bianjie.ai/cschain/cschain" && git config --global url."https://bamboo:EAVxPiH3DiuYUaTWRzgz@github.com/bianjieai/iritamod".insteadOf "https://github.com/bianjieai/iritamod" && make all

FROM alpine:3.10

COPY --from=builder /go/src/cosmos-sync /usr/local/bin

CMD cosmos-sync
