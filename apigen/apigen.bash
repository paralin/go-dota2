#!/bin/bash
set -eo pipefail

echo "This script should be run in GOPATH due to go-mod compatibility issues."
GO111MODULE=on go build -o apigen ./
GO111MODULE=on go mod vendor
GO111MODULE=off ./apigen generate-api
