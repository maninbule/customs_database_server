# syntax=docker/dockerfile:1

FROM golang:1.21
WORKDIR /app

ENV GOPROXY=https://goproxy.io,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
# 入口脚本
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh
EXPOSE 8082

ENTRYPOINT ["entrypoint.sh"]