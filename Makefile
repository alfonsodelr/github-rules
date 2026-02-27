.PHONY: build run swagger lint

build:
	go build -o bin/server .

run:
	go run main.go

lint:
	golangci-lint run ./...

swagger:
	swag init -g main.go --output docs/

all: swagger build

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build    - Build the server binary"
	@echo "  run      - Run the server"
	@echo "  lint     - Run golangci-lint"
	@echo "  swagger  - Generate Swagger documentation"
	@echo "  all      - Generate Swagger documentation and build the server"