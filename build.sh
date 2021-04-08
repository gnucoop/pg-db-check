#!/bin/sh

GOOS=linux go build -ldflags="-w -s" -o pg-db-check && upx pg-db-check
