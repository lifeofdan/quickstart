# Makefile for Go project

all: build

build:
	@echo "Build go app..."
	@templ generate
	@go build -o tmp/main cmd/main.go
	
# Live reload
dev:
	air
