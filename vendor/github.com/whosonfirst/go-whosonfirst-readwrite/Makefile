CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-readwrite
	cp -r http src/github.com/whosonfirst/go-whosonfirst-readwrite/
	cp -r pruner src/github.com/whosonfirst/go-whosonfirst-readwrite/
	cp -r reader src/github.com/whosonfirst/go-whosonfirst-readwrite/
	cp -r utils src/github.com/whosonfirst/go-whosonfirst-readwrite/
	cp -r writer src/github.com/whosonfirst/go-whosonfirst-readwrite/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

docker-build:
	docker build -t wof-readwrited .

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-cache"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-cli"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"       
	@GOPATH=$(GOPATH) go get -u "github.com/facebookgo/atomicfile"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt http/*.go
	go fmt pruner/*.go
	go fmt reader/*.go
	go fmt utils/*.go
	go fmt writer/*.go

bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/wof-fs-reader cmd/wof-fs-reader.go
