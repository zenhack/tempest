package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Incoming request. method = %v, headers:", req.Method)
		for k, vs := range req.Header {
			for _, v := range vs {
				log.Printf("%v: %v", k, v)
			}
		}

		w.Header().Set("X-Sandstorm-App-Custom-Header", "some value")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Test!"))
	})
	panic(http.ListenAndServe(":8000", nil))
}
