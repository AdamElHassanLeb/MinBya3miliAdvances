# Set default values, allowing overrides from the environment or command line
GOROOT ?= $(shell go env GOROOT)
GOPATH ?= $(shell go env GOPATH)
OUTPUT_DIR ?= ./build
OUTPUT_BIN ?= $(OUTPUT_DIR)/api
PACKAGE ?= ./cmd/api

# Ensure the output directory exists
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Build the application
build: $(OUTPUT_DIR)
	$(GOROOT)/bin/go build -o $(OUTPUT_BIN) $(PACKAGE)

# Run the application
run: build
	$(OUTPUT_BIN)

# Clean the build directory
clean:
	rm -rf $(OUTPUT_DIR)

# Run all targets
all: clean build run
