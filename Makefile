all: build
build install dev test-app export-import:
	@# Just shell out to make.go.
	go run internal/make/make.go $@
update-deps:
	# Update the versions of these in go.mod:
	go get capnproto.org/go/capnp/v3
	go get zenhack.net/go/util
	go get zenhack.net/go/vdom
	go get zenhack.net/go/websocket-capnp
	# and clean up:
	go mod tidy
clean:
	cd c && $(MAKE) clean
	rm -rf _build
	rm -f \
		go/internal/server/embed/*.wasm \
		c/config.h  \
		go/internal/config/config.go
	find * -type f -name '*.capnp.go' -delete
	find * -type f -name '*.cgr' -delete
	find * -type d -empty -delete
check: all
	go test \
		./internal/server/... \
		./pkg/...
nuke: clean
	rm -f config.json

.PHONY: all configure build run check clean nuke install
