all: static/index.html server
static/index.html: $(wildcard elm/*.elm) elm/elm-package.json
	cd elm ; elm-make Main.elm --output ../static/index.html
server:
	go build -v -i -o $@
clean:
	rm -f static/index.html elm/elm-stuff server
run: all
	./server

.PHONY: run all server
