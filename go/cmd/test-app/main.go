package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Test!"))
	})
	panic(http.ListenAndServe(":8000", nil))
}
