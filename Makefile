#!/usr/bin/make -f

export GO111MODULE = on
export GOPROXY = https://goproxy.io

install:
	@echo "installing bsn-irita-adapter..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/bsn-irita-adapter
