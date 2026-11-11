#!/bin/bash

cd ../build
#go mod init hamlib_rest_api
GOOS=linux GOARCH=amd64 go build -o binaries/hamlib_rest_api
#GOOS=linux GOARCH=arm64 go build -o binaries/hamlib_rest_api
killall hamlib_rest_api || true
./binaries/hamlib_rest_api

