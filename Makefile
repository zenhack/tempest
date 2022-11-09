all: build
configure build run clean nuke install:
	@# Just shell out to make.go.
	go run make.go $@

.PHONY: all configure build run clean nuke install
