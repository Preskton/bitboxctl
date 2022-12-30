

.PHONY: build
build:
	go build -o dist/bin/bitboxctl .

.PHONY: dist
dist:
	GOOS=linux GOARCH=amd64 go build -o dist/bin/bitboxctl-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/bin/bitboxctl-linux-arm64 .
	GOOS=linux GOARCH=arm go build -o dist/bin/bitboxctl-linux-arm .

.PHONY: run
run:
	@go run .

.PHONY: docs
docs: build
	dist/bin/bitboxctl codegen docs
