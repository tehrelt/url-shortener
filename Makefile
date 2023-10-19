.PHONY: build
build:
	go build -v ./cmd/url-shortener


.DEFAULT_GOAL := build