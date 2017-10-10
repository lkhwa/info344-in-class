#!/usr/bin/env bash
set -e #if error, stop everything
GOOS=linux go build
docker build -t lkhwa/testserver .
docker push lkhwa/testserver