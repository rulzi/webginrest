GO111MODULE=on

build:
	export GO111MODULE on; \
	go build ./...

build-generate:
	export GO111MODULE on; \
	go build -o webginrest webginrest/cmd

lint: build
	golint -set_exit_status ./...