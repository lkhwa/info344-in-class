#!/usr/bin/env bash
docker rm -f zipsvr
docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/leannehwa/Desktop/code/go/src/github.com/lkhwa/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
>>>>>>> ee296d65393eb82cb7090e69ee848d296c254143
