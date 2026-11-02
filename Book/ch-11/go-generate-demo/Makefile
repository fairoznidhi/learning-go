.PHONY: generate build test

generate:
	go generate ./...

build: generate
	go build ./...

test: generate
	go test ./...