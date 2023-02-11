BINARY_NAME := extract_bitwarden_totp
BIN_DIR := bin

all: build

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/cli

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/*
