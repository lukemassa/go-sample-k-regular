
.PHONY: all run

run: all
	@go run main.go

ci: all
	go fmt ./...
	go vet ./...
	go test -cover ./...

test: all
	go test -cover ./...

all:
	@which go >/dev/null || (echo 'ERROR: go is not installed. Please visit: https://golang.org/doc/install' && false)
