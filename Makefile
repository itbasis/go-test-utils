go-dependencies:
	# https://asdf-vm.com/
	asdf install golang

	#
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	#
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

	asdf reshim golang || :
	#
	go get -u -t -v ./... || :

go-generate: go-dependencies
	go generate ./...

go-lint:
	golangci-lint run

go-test:
	go vet -vettool=$(which shadow) ./...
	gosec ./...

go-all: go-dependencies go-generate go-lint go-test
	go mod tidy || :
