#!/bin/bash

go clean --cache && go test -v -cover github.com/ygt1qa/microservices/...
go build -o authentication/authsvc authentication/main.go
go build -o api/apisvc api/main.go
