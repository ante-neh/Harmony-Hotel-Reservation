build:
	@echo "Building..."
	@go build -o bin/cmd/api
run:build
	@./bin/cmd/api

test:
	@go test -v ./..