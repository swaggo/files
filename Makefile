all: build

.PHONY: build
build:
	rm -rf dist
	cp -r swagger-ui/dist .
