#!/bin/bash

set -e

export GOPATH=$GOPATH

CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -installsuffix cgo -o kendo main.go