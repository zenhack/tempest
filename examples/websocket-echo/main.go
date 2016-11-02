package main

import (
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"time"
	"zenhack.net/go/sandstorm/grain"
	"zenhack.net/go/sandstorm/websession"
)

var mainPage = []byte(`<!doctype html>
<html>
	<head>
		<title>Websocket echo server</title>
		<script>
		document.addEventListener("DOMContentLoaded", function(event) {

			window.addEventListener("message", function(event) {
				if (event.data.rpcId !== "0") {
					return;
				}
				if (event.data.error) {
					console.log("ERROR: " + event.data.error);
					return;
				}
				var elt = document.getElementById("offer-iframe");
				elt.setAttribute("src", event.data.uri);
			});

			window.parent.postMessage({
				renderTemplate: {
					rpcId: "0",
					template: window.location.protocol.replace("http", "ws") +
							"//$API_HOST/.sandstorm-token/$API_TOKEN/connect"
				}
			}, "*");

		});
		</script>
	</head>
	<body>
		<p>Connect to the echo server at:</p>

		<iframe id="offer-iframe"></iframe>
	</body>
</html>
`)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write(mainPage)
	})

	http.Handle("/connect", websocket.Handler(func(conn *websocket.Conn) {
		io.Copy(conn, conn)
	}))

	ctx := context.Background()
	_, err := grain.ConnectAPI(ctx, websession.FromHandler(ctx, http.DefaultServeMux))
	if err != nil {
		log.Fatalln(err)
	}
	<-ctx.Done()
}
