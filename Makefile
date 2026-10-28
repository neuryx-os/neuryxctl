BINARY  := neuryxctl
VERSION ?= dev
LDFLAGS := -X github.com/neuryx-os/neuryxctl/internal/version.Version=$(VERSION)

.PHONY: all build test lint clean

all: lint test build

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/neuryxctl

test:
	go test ./...

lint:
	test -z "$$(gofmt -l .)"
	go vet ./...

clean:
	rm -rf bin dist
