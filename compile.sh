#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o hamlib-rest-api-linux64
GOOS=linux GOARCH=arm64 go build -o hamlib-rest-api-rpi64

#GOOS=windows GOARCH=amd64 go build -o hamlib-rest-api-windows-amd64.exe
#GOOS=windows GOARCH=arm64 go build -o hamlib-rest-api-windows
#arm64.exe
#GOOS=darwin GOARCH=amd64 go build -o hamlib-rest-api-m
#acos-amd64
#GOOS=darwin GOARCH=arm64 go build -o hamlib-rest-api-m
#acos-arm64
