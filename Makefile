# Makefile for Go project

# Build app

build:
	@echo "Building..."
	@go build -o tmp/main cmd/api/main.go

# Live reload
dev:
	air
