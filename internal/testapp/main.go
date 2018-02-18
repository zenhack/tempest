package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"strconv"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/exp/websession"

	"zombiezen.com/go/capnproto2/rpc"
)

type UiView struct {
	*websession.HandlerUiView
}

func (*UiView) GetViewInfo(grain.UiView_getViewInfo) error { return nil }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`<a href="/static/">main</a>`))
	})

	http.Handle("/static/", http.FileServer(http.Dir("")))

	http.HandleFunc("/echo-request/", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		wantStatus := req.URL.Query().Get("want-status")
		if wantStatus == "" {
			wantStatus = "200"
		}
		var status int
		n, err := fmt.Sscanf(wantStatus, "%d", &status)
		if err != nil {
			log.Printf("Error parsing want-status %q: %v", wantStatus, err)
			status = 400
		} else if n != 1 {
			// I(zenhack) think this is impossible, but the docs for fmt.Sscanf
			// don't actually say; let's be careful.
			log.Printf("Error parsing want-status %q: fmt.Sscanf reported "+
				"no succesful parses.", wantStatus)
			status = 400
		}
		wantLang := req.URL.Query().Get("want-lang")
		if wantLang != "" {
			w.Header().Set("Content-Language", wantLang)
		}
		wantFilename := req.URL.Query().Get("filename")
		if wantFilename == "" {
			w.Header().Set("Content-Disposition", "inline")
		} else {
			w.Header().Set("Content-Disposition", mime.FormatMediaType(
				"attachment", map[string]string{"filename": wantFilename},
			))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "identity")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"url":     req.URL,
			"method":  req.Method,
			"headers": req.Header,
			"body":    body,
			"cookies": req.Cookies(),
		})
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/content-length", func(w http.ResponseWriter, req *http.Request) {
		payload := "Hello, Sandstorm!\n"
		w.Header().Set("Content-Length", strconv.Itoa(len(payload)))
		w.Write([]byte(payload))
	})

	http.HandleFunc("/set-cookie", func(w http.ResponseWriter, req *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "bob",
			Value:  "chocolate chip",
			Secure: true,
		})
	})

	http.HandleFunc("/no-op-handler", func(w http.ResponseWriter, req *http.Request) {
		// Make sure stuff works even if we don't do anything in the handler.
	})

	file := os.NewFile(3, "<sandstorm rpc socket @ fd #3>")
	conn, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	transport := rpc.StreamTransport(conn)
	rpc.NewConn(transport, rpc.MainInterface(grain.UiView_ServerToClient(&UiView{
		HandlerUiView: &websession.HandlerUiView{
			Handler: http.DefaultServeMux,
		},
	}).Client))
	<-context.Background().Done()
}
