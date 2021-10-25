#!/bin/bash

ls -la
cd /app/cmd/lbc-test
go build -o lbc-test main.go
./lbc-test