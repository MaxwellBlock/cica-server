FROM golang:1.16

LABEL maintainer="lawga2020@protonmail.com"
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
WORKDIR /app
ADD . /app
RUN cd /app && go build