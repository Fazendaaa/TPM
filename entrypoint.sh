#!/bin/sh
mkcert -install
mkcert -cert-file /etc/certs/local-cert.pem -key-file /etc/certs/local-key.pem ${DOMAIN} "*.${DOMAIN}" ${IPSERVER}
/usr/local/bin/tpm
