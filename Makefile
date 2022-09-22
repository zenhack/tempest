all: build
configure build run clean nuke install:
	go run make.go $@

.PHONY: all configure build run clean nuke install
