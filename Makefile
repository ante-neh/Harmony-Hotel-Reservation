build:
	@go build -o bin/binary cmd/api/main.go

run:build
	@go run ./bin/binary

test:
	@go test -v ./...