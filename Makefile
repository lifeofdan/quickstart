# Makefile for Go project

# Build app
all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o tmp/main cmd/api/main.go

# Live reload
dev:
	air
