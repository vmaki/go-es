FROM golang:1.19-alpine as builder

WORKDIR /build/go-es

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o main .

FROM alpine:latest as prod

WORKDIR /root

COPY --from=0 /build/go-es/main .
COPY --from=0 /build/go-es/config/settings.docker.yml ./config/settings.docker.yml

EXPOSE 7003
ENTRYPOINT ./main --env=docker
