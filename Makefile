all: install_dependencies build_binaries format lint test 

install_dependencies:
	go get ./...

build_binaries: install_dependencies
	go build -o ./httpclient main.go

run_binaries: build_binaries
	./httpclient

format:
	go fmt ./...
lint:
	sudo snap install golangci-lint --classic
	golangci-lint run ./...

test:
	go test ./... -v
