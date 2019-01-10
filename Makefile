CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

self:   prep
	if test ! -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-export
	cp export.go src/github.com/whosonfirst/go-whosonfirst-export/
	cp -r exporter src/github.com/whosonfirst/go-whosonfirst-export/
	cp -r options src/github.com/whosonfirst/go-whosonfirst-export/
	cp -r properties src/github.com/whosonfirst/go-whosonfirst-export/
	cp -r uid src/github.com/whosonfirst/go-whosonfirst-export/
	cp -r vendor/* src/

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/tidwall/gjson"
	@GOPATH=$(GOPATH) go get -u "github.com/tidwall/pretty"
	@GOPATH=$(GOPATH) go get -u "github.com/tidwall/sjson"
	@GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-artisanal-integers"
	@GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-brooklynintegers-api"
	rm -rf src/github.com/aaronland/go-brooklynintegers-api/vendor/github.com/tidwall
	rm -rf src/github.com/aaronland/go-brooklynintegers-api/vendor/github.com/aaronland/go-artisanal-integers

vendor-deps: deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt *.go
	go fmt exporter/*.go
	go fmt options/*.go
	go fmt properties/*.go
	go fmt uid/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-export-feature cmd/wof-export-feature.go
