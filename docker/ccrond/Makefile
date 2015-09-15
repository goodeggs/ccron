#!/usr/bin/make -f
SHELL=/bin/bash

VERSION=`jq -r '.PackageVersion' ../../.goxc.json`

.PHONY: all deps build release

all: deps build

deps:
	tar xzf ${GOPATH}/bin/ccron-xc/${VERSION}/ccron_${VERSION}_linux_amd64.tar.gz --strip-components=1 ccron_${VERSION}_linux_amd64/ccron

build: deps
	docker build -t goodeggs/ccrond:latest .

release: build
	docker tag goodeggs/ccrond:latest goodeggs/ccrond:${VERSION} && \
	docker push goodeggs/ccrond:${VERSION} && \
	docker push goodeggs/ccrond:latest

