
ci:
	-which apt-get && apt-get update -y && apt-get install -y libmagic-mgc || true
	$(MAKE) cover

fmt:
	gofumpt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

lint:
	golangci-lint run

test:
	go test ./...

.PHONY: cover
cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html
