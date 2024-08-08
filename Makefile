build:
	@go build -o bin/binary cmd/api/main.go

run:build
	@go run ./cmd/api/main.go

test:
	@go test -v ./...