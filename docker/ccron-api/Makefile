#!/usr/bin/make -f
SHELL=/bin/bash

VERSION=`jq -r '.PackageVersion' ../../.goxc.json`

.PHONY: all deps build release

all: deps build

deps:
	tar xzf ${GOPATH}/bin/ccron-xc/${VERSION}/ccron_${VERSION}_linux_amd64.tar.gz --strip-components=1 ccron_${VERSION}_linux_amd64/ccron-api

build: deps
	docker build -t goodeggs/ccron-api:latest .

release: build
	docker tag goodeggs/ccron-api:latest goodeggs/ccron-api:${VERSION} && \
	docker push goodeggs/ccron-api:${VERSION} && \
	docker push goodeggs/ccron-api:latest

