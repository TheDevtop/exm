help:
	@echo 'EXM Makefile'
	@echo 'Targets: build/clean/docker/help'

clean:
	@go clean

build:
	@go fmt
	@go build -ldflags "-s -w"

docker:clean
	@docker build -t thedevtop/exm:build .
