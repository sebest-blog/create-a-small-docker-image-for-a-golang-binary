#!/bin/bash
tar cfz zoneinfo.tar.gz /usr/share/zoneinfo
curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
docker build -t demo .
docker run demo
