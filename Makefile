.PHONY: build generate-module

# Path to the CLI binary or go run command
ANVIL_CMD = go run cmd/anvil/main.go

build:
	@go build -o bin/anvil cmd/anvil/main.go
	@echo "Anvil built successfully in bin/anvil"

generate-module:
	@if [ -z "$(entity)" ]; then \
		echo "Error: You must provide an entity name."; \
		echo "Usage: make generate-module entity=Vehicle"; \
		exit 1; \
	fi
	@$(ANVIL_CMD) generate $(entity)

init:
	@$(ANVIL_CMD) init $(path)
