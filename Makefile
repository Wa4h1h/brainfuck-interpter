.PHONY: fmt
fmt:
	@gofumpt -l -w .

build: test
	@go build -o bfi main.go
test:
	@go test -v ./...