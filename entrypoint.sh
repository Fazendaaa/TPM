#!/bin/sh
mkcert -install
mkcert -cert-file /etc/certs/local-cert.pem -key-file /etc/certs/local-key.pem ${DOMAIN} "*.${DOMAIN}"
/usr/local/bin/tpm
