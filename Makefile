all: build

.PHONY: init
init:
	git submodule update --init --recursive

.PHONY: update
update:
	git submodule update --remote

.PHONY: build
build:
	rm -rf dist/*
	cp -r swagger-ui/dist/* dist/
