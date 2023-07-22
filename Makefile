.PHONY: build clean vet fmt help

all: build

build:
	go build -ldflags "-w -s" -o simple-go-gin-example main.go

run:
	./simple-go-gin-example -c conf/config.yaml

vet:
	go vet ./...; true

fmt:
	gofmt -w .

clean:
	rm -rf simple-go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make run: ./simple-go-gin-example"
	@echo "make vet: go vet ./..."
	@echo "make fmt: gofmt -w ."
	@echo "make clean: remove object files and cached files"