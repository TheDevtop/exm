help:
	@echo 'EXM Makefile'
	@echo 'Targets: build/clean/docker/help'

clean:
	@go clean

build:
	@go fmt
	@go build -ldflags "-s -w"

docker:clean
	@docker build -t ghcr.io/thedevtop/exm:v0.2.0 .
