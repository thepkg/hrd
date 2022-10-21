APP_NAME := $(shell basename "$(CURDIR)")

# Build golang application binary.
.PHONY: build
build: GOOS ?= linux
build: GOARCH ?= amd64
build:
	@echo "==> Starting build"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME) ./main.go
	@echo "==> Done! Binary size: "
	@du -h $(APP_NAME)
