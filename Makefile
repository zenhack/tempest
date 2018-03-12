all: build
configure build run clean nuke:
	go run make.go $@

.PHONY: all configure build run clean nuke
