build:
	@go build -o bin/somesome

run: build
	@./bin/somesome

test:
	@go test -v ./...
