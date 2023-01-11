package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Incoming request.")
		log.Println("method = ", req.Method)
		log.Printf("url = %q\n", req.URL)
		log.Println("Headers:")

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
