#!/usr/bin/env bash
docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/leannehwa/Desktop/code/go/src/github.com/lkhwa/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
lkhwa/zipserver