.SUFFIXES:

.SUFFIXES: .go

VERSION = $(shell git describe --abbrev=0)

.PHONY: default
default: debug

.PHONY: debug
debug: debug/moe-symlink

debug/moe-symlink: src/main.go
	go build -ldflags "-X main.version=$(VERSION)" -o $@ $<

.PHONY: release
release: release/moe-symlink.exe
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ $<
