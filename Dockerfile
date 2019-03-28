FROM golang:1.11.1 AS build-env

ADD . $GOPATH/src/pure-media
WORKDIR $GOPATH/src/pure-media

ENV CGO_ENABLED=0
ENV GO111MODULE=on
#ENV GOPROXY="http://172.16.0.18:10081"
#ENV ALL_PROXY="http://172.16.0.18:10081"
ARG GOPROXY=""

RUN export GOPROXY=$GOPROXY \
    && mkdir /output \
    && go build -o /output/pure-media \
    && curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl \
    && chmod a+rx /usr/local/bin/youtube-dl

FROM alpine:latest
MAINTAINER lizhiyuan "lizhiyuan@apkpure.net"
WORKDIR /app
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime
COPY --from=build-env /output/pure-media /app/
CMD ["/app/pure-media"]

