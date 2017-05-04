.PHONY: run

SHELL     := /bin/bash
#PKG_NAME  = github.com/ricardo-ch/auto-product-import
IMAGENAME = go-postgresql-crud
CONTAINERNAME = go-postgresql-crud
BINARY    = go-postgresql-crud
FILES     = $(shell go list ./... | grep -v /vendor/)

${BINARY}:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix cgo -o $(BINARY) .

run:
	go run main.go