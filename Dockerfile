FROM golang:1.22-alpine AS base

WORKDIR /usr/src/


FROM alpine:1.22 AS runner

WORKDIR /app/

RUN apk install mkcert

