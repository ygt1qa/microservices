#!/bin/bash

go clean --cache && go test -v -cover github.com/ygt1qa/microservices/authentication/...
go build -o authentication/authsvc authentication/main.go