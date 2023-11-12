all: fmt vet test solved

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test -race -v ./...

solved:
	ls -ld -- */ | wc -l