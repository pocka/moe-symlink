.SUFFIXES:

.SUFFIXES: .go

VERSION = $(shell git describe --abbrev=0 --tags)

.PHONY: default
default: debug

.PHONY: debug
debug: debug/moe-symlink

debug/moe-symlink: src/main.go
	go build -ldflags "-X main.version=$(VERSION)" -o $@ $<

.PHONY: release
release: release/moe-symlink.exe

release/moe-symlink.exe: src/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ $<
