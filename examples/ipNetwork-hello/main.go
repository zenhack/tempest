package main

import (
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	grain_capnp "zenhack.net/go/sandstorm/capnp/grain"
	ip_capnp "zenhack.net/go/sandstorm/capnp/ip"
	ws_capnp "zenhack.net/go/sandstorm/capnp/websession"
	"zenhack.net/go/sandstorm/grain"
	"zenhack.net/go/sandstorm/ip"
	"zenhack.net/go/sandstorm/websession"
	"zombiezen.com/go/capnproto2"
)

var thePage = `<!doctype html>
<html>
    <head><title>ipNetwork example app</title>
        <script>
            window.addEventListener("load", function(_event) {
                document.getElementById("request_cap").addEventListener("click", function() {
                    var rpcId = Math.random();
                    window.parent.postMessage({powerboxRequest: {
                        rpcId: rpcId,
                        // ID for ipNetwork:
                        query: ["EAZQAQEAABEBF1EEAQH_QCAqemtXgqkAAAA"],
                    }}, "*");
                    window.addEventListener("message", function(event) {
                        if (event.data.rpcId !== rpcId) {
                            return;
                        }

                        if (event.data.error) {
                            console.error("rpc errored:", event.data.error);
                            return;
                        }

                        var xhr = new XMLHttpRequest();
                        xhr.open("POST", "/ip-network-cap", true);

                        // Hack to pass bytes through unprocessed.
                        xhr.overrideMimeType("text/plain; charset=x-user-defined");

                        xhr.send(event.data.token);
                    }, false);
                });
            });
        </script>
    </head>
    <body>
        <button id="request_cap">Request ipNetwork Capability</button>
        <form action="/sayhello" method="post">
            <div>
                <label for="message">Message: </label>
                <input type="text" name="message" />
            </div>
            <div>
                <label for="address">Address (e.g. example.com:4444):</label>
                <input type="text" name="address" />
            </div>
            <button type="submit">Submit</button>
        </form>
    </body>
</html>
`

type myUiView struct {
	Ctx context.Context
}

func (v myUiView) NewSession(p grain_capnp.UiView_newSession) error {
	mux := http.NewServeMux()
	sessionCtx := p.Params.Context()
	dialer := &ip.IpNetworkDialer{}
	dialer.Ctx = v.Ctx

	checkErr := func(err error) bool {
		if err != nil {
			println(err.Error())
		}
		return err != nil
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(thePage))
	})

	mux.HandleFunc("/ip-network-cap", func(w http.ResponseWriter, req *http.Request) {
		buf, err := ioutil.ReadAll(req.Body)
		if checkErr(err) {
			return
		}
		results, err := sessionCtx.ClaimRequest(
			v.Ctx,
			func(p grain_capnp.SessionContext_claimRequest_Params) error {
				p.SetRequestToken(string(buf))
				return nil
			}).Struct()
		if checkErr(err) {
			return
		}
		capability, err := results.Cap()
		if checkErr(err) {
			return
		}
		dialer.IpNetwork = ip_capnp.IpNetwork{Client: capnp.ToInterface(capability).Client()}
	})

	mux.HandleFunc("/sayhello", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if checkErr(err) {
			return
		}
		address := req.PostFormValue("address")
		message := req.PostFormValue("message")

		conn, err := dialer.Dial("tcp", address)
		if checkErr(err) {
			return
		}
		defer conn.Close()
		conn.Write([]byte(message))
	})
	client := ws_capnp.WebSession_ServerToClient(websession.FromHandler(v.Ctx, mux)).Client
	p.Results.SetSession(grain_capnp.UiSession{Client: client})
	return nil
}

func (v myUiView) GetViewInfo(p grain_capnp.UiView_getViewInfo) error {
	return nil
}

func (v myUiView) NewRequestSession(p grain_capnp.UiView_newRequestSession) error {
	return nil
}

func (v myUiView) NewOfferSession(p grain_capnp.UiView_newOfferSession) error {
	return nil
}

func main() {
	ctx := context.Background()
	_, err := grain.ConnectAPI(ctx, myUiView{ctx})
	if err != nil {
		log.Fatalln(err)
	}
	<-ctx.Done()
}
