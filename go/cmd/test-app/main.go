package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("X-Sandstorm-App-Custom-Header", "some value")
		w.Write([]byte("Test!"))
	})
	panic(http.ListenAndServe(":8000", nil))
}
