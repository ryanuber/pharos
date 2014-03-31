GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "$(shell git status --porcelain)" && echo "+CHANGES" || :)

all:
	@go build \
		-ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
		-v \
		-o ${GOPATH}/bin/pharos

test:
	go test -v ./...
