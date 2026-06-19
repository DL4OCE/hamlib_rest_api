#!/bin/bash
go mod init hamlib_rest_api
GOOS=linux GOARCH=amd64 go build -o build/hamlib-rest-api
GOOS=linux GOARCH=arm64 go build -o build/hamlib-rest-api

#GOOS=windows GOARCH=amd64 go build -o build/hamlib-rest-api-windows-amd64.exe
#GOOS=windows GOARCH=arm64 go build -o build/hamlib-rest-api-windows
#arm64.exe
#GOOS=darwin GOARCH=amd64 go build -o hamlib-rest-api-macos-amd64
#GOOS=darwin GOARCH=arm64 go build -o hamlib-rest-api-macos-arm64

