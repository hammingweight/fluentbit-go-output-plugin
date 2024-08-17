.DEFAULT_GOAL := build

fmt:
	go fmt .
.PHONY:fmt

lint: fmt
	golint .
.PHONY:lint

vet: fmt
	go vet .
.PHONY:vet

build: vet
	go build -buildmode=c-shared -o out_noop_plugin.so
.PHONY:build
