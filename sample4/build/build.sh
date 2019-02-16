#!/bin/bash
GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o restapp github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication