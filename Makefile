
ci:
	-which apt-get && sudo apt-get update -y && sudo apt-get install -y libmagic-mgc
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
