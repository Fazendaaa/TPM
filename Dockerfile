FROM golang:1.22-alpine AS base


FROM alpine:1.22 AS runner

RUN apk install mkcert

