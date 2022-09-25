all: build
configure build run clean nuke install:
	@# Just shell out to make.go.
	@#
	@# Note: The sandstorm-sandbox-agent binary *MUST* be linked statically,
	@# since we may not have access to a suitable libc inside the sandbox.
	@#
	@# For some reason setting this environment variable on the invocation
	@# of go build inside make.go causes an error, so we do it here. This
	@# has the effect of also statically linking all other go binaries, but
	@# that's fine.
	CGO_ENABLED=0 go run make.go $@

.PHONY: all configure build run clean nuke install
