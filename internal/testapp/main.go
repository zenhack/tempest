package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"zenhack.net/go/sandstorm/capnp/grain"
	"zenhack.net/go/sandstorm/exp/websession"

	"zombiezen.com/go/capnproto2/rpc"
	"zombiezen.com/go/capnproto2/server"
)

type UiView struct {
	*websession.HandlerUiView
}

func (*UiView) GetViewInfo(context.Context, grain.UiView_getViewInfo) error { return nil }

func copyStreaming(dest io.Writer, src io.Reader) (int64, error) {
	var (
		buf          [4096]byte
		totalWritten int64
	)
	for {
		countR, errR := src.Read(buf[:])
		countW, errW := dest.Write(buf[:countR])
		totalWritten += int64(countW)
		if flusher, ok := dest.(http.Flusher); ok {
			flusher.Flush()
		}
		switch errR {
		case io.EOF:
			return totalWritten, nil
		case nil:
			if errW != nil {
				return totalWritten, errW
			}
		default:
			return totalWritten, errR
		}
	}
}

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

	http.HandleFunc("/return-headers", func(w http.ResponseWriter, req *http.Request) {
		headers := map[string]string{}
		err := json.NewDecoder(req.Body).Decode(&headers)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error decoding body: " + err.Error()))
			return
		}
		for key, val := range headers {
			w.Header().Set(key, val)
		}
	})

	http.HandleFunc("/etag", func(w http.ResponseWriter, req *http.Request) {
		// NOTE: This does not handle etags in full generality; it's just
		// enough for the one test case in our test suite.
		ifNoneMatch := req.Header.Get("If-None-Match")
		if strings.HasPrefix(ifNoneMatch, "W/") {
			ifNoneMatch = strings.TrimPrefix(ifNoneMatch, "W/")
		}
		value := `"sometag"`

		if req.URL.Query().Get("weak") == "true" {
			value = `W/` + value
		}
		w.Header().Set("ETag", value)

		if ifNoneMatch == value {
			w.WriteHeader(http.StatusNotModified)
		} else {
			w.WriteHeader(http.StatusOK)
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

	http.HandleFunc("/redirect", func(w http.ResponseWriter, req *http.Request) {
		log.Print("Got redirect request....")
		status := req.URL.Query().Get("want-status")
		statusCode, err := strconv.Atoi(status)
		if err != nil {
			log.Print("Error parsing status:", err)
			statusCode = 400
		}
		w.Header().Set("Location", "/redirect-landing")
		w.WriteHeader(statusCode)
	})

	http.HandleFunc("/no-op-handler", func(w http.ResponseWriter, req *http.Request) {
		// Make sure stuff works even if we don't do anything in the handler.
	})

	http.HandleFunc("/echo-body", func(w http.ResponseWriter, req *http.Request) {
		log.Print("Got echo body request")
		w.WriteHeader(200)
		n, err := copyStreaming(w, req.Body)
		if err != nil {
			log.Printf("Error in copyStreaming: %q", err)
		}
		log.Printf("copyStreaming: wrote %d bytes.", n)
	})

	if os.Getenv("SANDSTORM") != "1" {
		log.Print("not running in sandstorm; listening on :8434")
		panic(http.ListenAndServe(":8434", nil))
	}

	file := os.NewFile(3, "<sandstorm rpc socket @ fd #3>")
	conn, err := net.FileConn(file)
	if err != nil {
		panic(err)
	}
	transport := rpc.NewStreamTransport(conn)
	rpc.NewConn(
		transport,
		&rpc.Options{
			BootstrapClient: grain.UiView_ServerToClient(
				&UiView{
					HandlerUiView: &websession.HandlerUiView{
						Handler: http.DefaultServeMux,
					},
				},
				&server.Policy{},
			).Client,
		})
	<-context.Background().Done()
}
