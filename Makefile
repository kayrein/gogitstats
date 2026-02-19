BINARY_NAME=gogitstats
PACKAGE_PATH=github.com/kayrein/gogitstats

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -X main.Version=$(VERSION)

.PHONY: build install clean

build:
	mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY_NAME) cmd/gogitstats/*

install:
	go install -ldflags "$(LDFLAGS)" $(PACKAGE_PATH)

clean:
	rm -f $(BINARY_NAME)
