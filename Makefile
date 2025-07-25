.PHONY: fmt lint

BIND_ADDRESS ?= 0.0.0.0
PORT ?= 3000

# Build the server
build:
	docker build -t lunchmoney-mcp-server:dev .

# Run the server
run: build
	docker run -it --rm \
		-p $(PORT):$(PORT) \
		-e LUNCHMONEY_TOKEN=$(LUNCHMONEY_TOKEN) \
		-e BIND_ADDRESS=$(BIND_ADDRESS) \
		-e PORT=$(PORT) \
		lunchmoney-mcp-server:dev

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./...


# Help
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the dev image"
	@echo "  run     - Run the dev image"
	@echo "  fmt     - Format the code"
	@echo "  lint    - Lint the code"
