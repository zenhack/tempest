package main

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
	"zenhack.net/go/sandstorm/grain"
	"zenhack.net/go/sandstorm/websession"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	ctx := context.Background()
	_, err := grain.ConnectAPI(ctx, websession.FromHandler(ctx, http.DefaultServeMux))
	if err != nil {
		log.Fatalln(err)
	}

	// We don't need to do anything else in this goroutine; our websession
	// is registered with sandstorm. But, if we fall of the end of main the
	// program will exit, so let's avoid that:
	for {
		time.Sleep(30 * time.Second)
	}
}
