#!/bin/bash
go mod init hamlib_rest_api
GOOS=linux GOARCH=amd64 go build -o build/hamlib_rest_api 
GOOS=linux GOARCH=arm64 go build -o build/hamlib_rest_api

#GOOS=windows GOARCH=amd64 go build -o build/hamlib_rest_api-windows-amd64.exe
#GOOS=windows GOARCH=arm64 go build -o build/hamlib_rest_api-windows
#arm64.exe
#GOOS=darwin GOARCH=amd64 go build -o hamlib_rest_api-macos-amd64
#GOOS=darwin GOARCH=arm64 go build -o hamlib_rest_api-macos-arm64

