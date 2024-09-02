#! /bin/bash
gofmt -w . && staticcheck && gotags ./*.go > tags && go test && go run . $1 $2 $3

