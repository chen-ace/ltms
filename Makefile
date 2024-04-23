# Makefile for building server and client executables and packaging them into a tarball
export GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
LDFLAGS := -s -w

# Binary names
SERVER_BINARY := ltmsd
CLIENT_BINARY := ltmsclient
SERVER_CONFIG_FILE := server_config.yaml

# Source files
SERVER_SOURCE := ./cmd/ltmsd/server.go
CLIENT_SOURCE := ./cmd/ltms/client.go
ENV_SOURCE := ./.env

# Build directory
BUILD_DIR := ./build

# OS and architecture information
OS := $(shell uname -s | tr A-Z a-z)
ARCH := $(shell uname -m)

# Tarball name with OS and architecture information
TARBALL := ltms-$(OS)-$(ARCH).tar.gz

# Default target
all: build_server build_client

# Build server
build_server:
	@echo "Building server..."
	@go build -trimpath -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(SERVER_BINARY) $(SERVER_SOURCE)
	@cp $(SERVER_CONFIG_FILE) $(BUILD_DIR)/$(SERVER_CONFIG_FILE)
	@chmod +x $(BUILD_DIR)/$(SERVER_BINARY)

# Build client
build_client:
	@echo "Building client..."
	@go build -trimpath -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(CLIENT_BINARY) $(CLIENT_SOURCE)
	@chmod +x $(BUILD_DIR)/$(CLIENT_BINARY)

# Package executables into a tarball
package:
	@echo "Packaging executables..."
	@cp $(ENV_SOURCE) $(BUILD_DIR)/example.env
	@tar -czvf $(BUILD_DIR)/$(TARBALL) -C $(BUILD_DIR) $(SERVER_BINARY) $(CLIENT_BINARY) $(BUILD_DIR)/example.env

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)/$(SERVER_BINARY) $(BUILD_DIR)/$(CLIENT_BINARY) $(BUILD_DIR)/$(TARBALL)

.PHONY: all build_server build_client package clean