PWD=$(shell pwd)
GO_BIN=$(shell echo $$GOPATH)/bin
PROJECT_NAME=asher
GO_FILES=$(wildcard *.go) $(wildcard */*.go)
MODULE_NAME=$(shell "head go.mod -n1|awk '{print $2}'")

install:
	@echo "installing"
	go env -w GOBIN=$(PWD)/bin
	go install $(MODULE_NAME)
	go env -u GOBIN