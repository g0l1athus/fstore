build:
	@go build -o bin/fstore

run: build
	@./bin/fstore

test:
	@go test ./... -v
