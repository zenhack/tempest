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

	err := sandstormhttpbridge.ListenAndServe(context.Background(), nil, ":8000", nil)
	if err != nil {
		panic(err)
	}
}
