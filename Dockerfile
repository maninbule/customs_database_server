# syntax=docker/dockerfile:1

FROM golang:1.21
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://goproxy.io,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
EXPOSE 8082
# 入口脚本
COPY enter.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/enter.sh
ENTRYPOINT ["enter.sh"]