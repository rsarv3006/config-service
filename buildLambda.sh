#!/bin/bash

GOOS=linux GOARCH=arm64 go build -o bootstrap main.go
zip function.zip bootstrap
