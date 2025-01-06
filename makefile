# Makefile

# Load environment variables from the .env file
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Variables
TEST_FLAGS ?= ./... -v

# Default target
.PHONY: test
test:
	@echo "Running tests with environment variables from .env..."
	@go test $(TEST_FLAGS)

