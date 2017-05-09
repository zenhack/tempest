package main

import (
	"net/http"
	"zenhack.net/go/sandstorm/websession"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	websession.ListenAndServe(nil)
}
