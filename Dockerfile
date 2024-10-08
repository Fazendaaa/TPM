FROM golang:1.23.1-alpine3.20 AS base

WORKDIR /usr/src/

RUN apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/ mkcert=1.4.4-r14

COPY img/ img/
COPY go.sum .
COPY go.mod .

# RUN go mod download

COPY main.go .

COPY internal/ internal/
COPY server/ server/

#RUN go test -v
RUN go build -o tpm

#FROM scratch AS runner
FROM alpine:3.20 AS runner

WORKDIR /app/

COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /usr/bin/mkcert /usr/bin/mkcert
COPY --from=base /usr/src/img/ /app/img/
COPY --from=base /usr/src/server/static /app/server/static
COPY --from=base /usr/src/server/templates /app/server/templates
COPY --from=base /usr/src/tpm /usr/local/bin/tpm

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --retries=3 --start-interval=10s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1
ENTRYPOINT [ "./entrypoint.sh" ]
#ENTRYPOINT [ "/usr/local/bin/tpm" ]
