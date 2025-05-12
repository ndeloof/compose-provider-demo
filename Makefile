all: build

.PHONY: build ## Build the compose cli-plugin
build:
	go build -o demo main.go

.PHONY: install ## Build the compose cli-plugin
install: build
	cp demo ~/bin/demo
