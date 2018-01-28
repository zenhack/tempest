package main

import (
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	rootDomain = defaultTo(os.Getenv("ROOT_DOMAIN"), "local.sandstorm.io")
	listenAddr = defaultTo(os.Getenv("LISTEN_ADDR"), ":8000")
	apps       = map[string]string{
		"app-1": ":8101",
		"app-2": ":8102",
		"app-3": ":8103",
	}

	responseHeaderWhiteList = map[string]headerSanitizer{
		"Content-Type":   compose(onlyOne, mustSatisfy(isMediaType)),
		"Content-Length": compose(onlyOne, mustSatisfy(isNumber)),
	}
)

type headerSanitizer func(hdr []string) []string

func verbatium(want string) func(string) (string, bool) {
	return func(got string) (string, bool) {
		return got, want == got
	}
}

func any(fns ...headerSanitizer) headerSanitizer {
	return func(hdr []string) []string {
		for _, fn := range fns {
			ret := fn(hdr)
			if ret != nil {
				return ret
			}
		}
		return nil
	}
}

func compose(fns ...headerSanitizer) headerSanitizer {
	return func(hdr []string) []string {
		for _, fn := range fns {
			if hdr == nil {
				break
			}
			hdr = fn(hdr)
		}
		return hdr
	}
}

func isMediaType(s string) bool {
	_, _, err := mime.ParseMediaType(s)
	return err == nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func each(fn func(string) (string, bool)) headerSanitizer {
	return func(hdr []string) []string {
		for i := range hdr {
			var ok bool
			hdr[i], ok = fn(hdr[i])
			if !ok {
				return nil
			}
		}
		return hdr
	}
}

func onlyOne(hdr []string) []string {
	if len(hdr) > 1 {
		return nil
	}
	return hdr
}

func mustSatisfy(pred func(string) bool) headerSanitizer {
	return each(func(hdr string) (string, bool) {
		return hdr, pred(hdr)
	})
}

func defaultTo(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func SetAppHeaders(w http.ResponseWriter) {

}

func main() {
	r := mux.NewRouter()
	//rootR := r.Host(rootDomain).Subrouter()

	r.Host("{app}." + rootDomain).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		SetAppHeaders(w)
		appName := mux.Vars(req)["app"]
		appAddr, ok := apps[appName]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch req.Method {
		case "GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS":
		default:
			badRequest(w, "Bad Method")
			return
		}
		var newUrl url.URL
		newUrl.Scheme = "http"
		newUrl.Host = "127.0.0.1" + appAddr
		newUrl.Path = req.URL.Path
		newUrl.RawPath = req.URL.RawPath
		newUrl.RawQuery = req.URL.RawQuery
		newReq := http.Request{
			Method: req.Method,
			URL:    &newUrl,
			Header: http.Header{},
			Body:   req.Body,
		}
		resp, err := http.DefaultTransport.RoundTrip(&newReq)
		if err != nil {
			log.Printf("Error connecting to app %q: %q", appName, err)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error connecting to app."))
			return
		}
		switch resp.StatusCode {
		case 200, 201, 202, 024, 206, 207, 301, 302, 303, 304, 307, 308:
		case 400, 403, 404, 405, 406, 409, 410, 412, 413, 414, 415, 418, 422:
		case 500:
		default:
			switch {
			case resp.StatusCode >= 400 && resp.StatusCode < 500:
				badRequest(w, "Bad Request")
			default:
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}
			return
		}

		for header, sanitizer := range responseHeaderWhiteList {
			val, ok := resp.Header[header]
			if ok {
				w.Header()[header] = sanitizer(val)
			}
		}

		w.WriteHeader(resp.StatusCode)
		// TODO: Limit according to Content-Length & set a timeout.
		io.Copy(w, resp.Body)
	})

	http.Handle("/", r)
	http.ListenAndServe(listenAddr, nil)
}

func badRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`<!doctype html>
		<html>
			<head>
				<meta charset="utf-8" />
				<title>Bad Request</title>
			</head>
			<body>
				<p>Bad Request: ` + msg + `</p>
			</body>
		</html>`))
}
