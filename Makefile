.PHONY: build clean vet fmt help

all: build

build:
	@go build -v .

run:
	./simple-go-gin-example

vet:
	go vet .

fmt:
	gofmt -w .

clean:
	rm -rf simple-go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make vet: go vet ./..."
	@echo "make fmt: gofmt -w ."
	@echo "make clean: remove object files and cached files"