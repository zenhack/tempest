package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func ServeApp(appName string, w http.ResponseWriter, req *http.Request) {
	SetAppHeaders(w)
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
}
