help:
	@echo 'EXM Makefile'
	@echo 'Targets: build/clean/docker/help'

clean:
	@go clean

build:
	@go fmt
	@go build

docker:clean
	@docker build -t thedevtop/exm:build .
