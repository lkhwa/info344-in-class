#!/usr/bin/env bash
set -e #if error, stop everything
echo "building linux executable"
GOOS=linux go build
docker build -t lkhwa/zipserver .
docker push lkhwa/zipserver
go clean
