all: build
build install:
	@# Just shell out to make.go.
	go run go/internal/make/make.go $@
clean:
	cd c && $(MAKE) clean
	rm -rf _build
	rm -f \
		go/internal/server/embed/*.wasm \
		c/config.h  \
		go/internal/config/config.go
	find capnp -type f -name '*.capnp.go' -delete
	find capnp -type f -name '*.cgr' -delete
	find capnp -type d -empty -delete
nuke: clean
	rm -f config.json

.PHONY: all configure build run clean nuke install
