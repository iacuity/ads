.PHONY:	help build install clean

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

build: ## build application
	go build

install: ## build and install application
	go install -v ./...

fmt: ## code indentation
	go fmt ./...

clean:  ## go clean
	go clean
