prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/aaronland/go-artisanal-integers; then rm -rf src/github.com/aaronland/go-artisanal-integers; fi
	mkdir -p src/github.com/aaronland/go-artisanal-integers/
	cp *.go src/github.com/aaronland/go-artisanal-integers/
	cp -r application src/github.com/aaronland/go-artisanal-integers/
	cp -r client src/github.com/aaronland/go-artisanal-integers/
	cp -r engine src/github.com/aaronland/go-artisanal-integers/
	cp -r http src/github.com/aaronland/go-artisanal-integers/
	cp -r server src/github.com/aaronland/go-artisanal-integers/
	cp -r service src/github.com/aaronland/go-artisanal-integers/
	cp -r utils src/github.com/aaronland/go-artisanal-integers/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/algnhsa"

vendor-deps: rmdeps deps
	if test ! -d src; then mkdir src; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src
fmt:
	go fmt *.go
	go fmt application/*.go
	go fmt client/*.go
	go fmt cmd/*.go
	go fmt engine/*.go
	go fmt http/*.go
	go fmt server/*.go
	go fmt service/*.go
	go fmt utils/*.go

bin:    self
	if test ! -d bin; then mkdir bin; fi
	@GOPATH=$(shell pwd) go build -o bin/int cmd/int.go
	@GOPATH=$(shell pwd) go build -o bin/intd-client cmd/intd-client.go
	@GOPATH=$(shell pwd) go build -o bin/intd-server cmd/intd-server.go