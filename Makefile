.PHONY: bin linux

PWD := $(shell pwd)

default: bin

bin:
	go build -tags static,system_libgit2 -p 4 -o $(PWD)/_output/git-go github.com/git-roll/libgitgo/cmd/cli

linux:
	GOOS=linux GOARCH=amd64 $(MAKE)
