build:
	go build

test:
	go test

release:
	go run github.com/goreleaser/goreleaser --rm-dist
	./dist/serve_linux_amd64/serve -version
