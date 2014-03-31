all:
	go build -v -o ${GOPATH}/bin/pharos

test:
	go test -v ./...
