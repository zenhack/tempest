package main

import (
	"golang.org/x/net/context"
	"net/http"
	"zenhack.net/go/sandstorm/sandstormhttpbridge"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	bridge, err := sandstormhttpbridge.Dial(context.Background())
	if err != nil {
		panic(err)
	}
	sandstormhttpbridge.ListenAndServe(bridge, ":8080", nil)
}
