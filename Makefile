.PHONY: build start version

APP_NAME := WiemanImages
VERSION := 1.1.2

LDFLAGS := -ldflags "-X main.appVersion=$(VERSION) -X main.appName=$(APP_NAME)"

build:
	go build $(LDFLAGS) -o ./build

build\:start: build start

build\:linux:
	CGO_ENABLED=0 GOOS=linux go build $(LDFLAGS) -o ./build

start:
	./build

test:
	go test ./...

version:
	@echo $(VERSION)
