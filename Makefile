build:
	@go build -o bin/littleecomgoapi cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go