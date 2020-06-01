.PHONY: build

export GO111MODULE=on

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main ./cmd/app/main.go

run:
	go run ./cmd/app/main.go

test:
	go test -v ./... -cover
