all: build
configure build run install:
	@# Just shell out to make.go.
	go run go/internal/make/make.go $@
clean:
	cd c && $(MAKE) clean
	rm -rf _build
	rm -f go/internal/server/embed/*.wasm
nuke: clean
	rm -f \
		c/config.h \
		go/internal/config/config.go \
		config.json

.PHONY: all configure build run clean nuke install
