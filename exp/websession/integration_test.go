// +build integration_tests

package websession

// Integration tests; these operate against the test app in ../../internal/testapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
)

var baseUrl string

func init() {
	// Get the base URL for the test app from the environment. if
	// the requisite environment variable is not set, panic.
	baseUrl = os.Getenv("GO_SANDSTORM_TEST_APP")
	if baseUrl == "" {
		panic("Integration test: GO_SANDSTORM_TEST_APP environment " +
			"variable not defined.")
	}
}

type echoBody struct {
	Method  string         `json:"method"`
	Url     *url.URL       `json:"url"`
	Headers http.Header    `json:"headers"`
	Body    []byte         `json:"body"`
	Cookies []*http.Cookie `json:"cookies"`
}

func chkfatal(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func expectStatus(t *testing.T, want, got int) {
	if want != got {
		t.Fatal("Unexpected status code; expected", want, "but got", got)
	}
}

func TestBasicGetHead(t *testing.T) {
	successfulResponse := func(resp *http.Response, err error) {
		chkfatal(t, err)
		expectStatus(t, 200, resp.StatusCode)
	}
	successfulResponse(http.Get(baseUrl + "echo-request/hello"))
	successfulResponse(http.Head(baseUrl + "echo-request/hello"))
}

func TestGetCorrectInfo(t *testing.T) {

	resp, err := http.Get(baseUrl + "echo-request/hello")
	chkfatal(t, err)
	defer func() {
		if t.Failed() {
			t.Log("Response:", resp)
		}
	}()
	expectStatus(t, 200, resp.StatusCode)
	body := &echoBody{}
	err = json.NewDecoder(resp.Body).Decode(body)
	chkfatal(t, err)

	if body.Method != "GET" {
		t.Error("Wrong method:", body.Method)
	}
	if body.Url.Path != "/echo-request/hello" {
		t.Error("Wrong path:", body.Url.Path)
	}
}

func testHeader(t *testing.T, name, sendVal, wantRecvVal string) {
	req, err := http.NewRequest("GET", baseUrl+"echo-request/"+name+"-header", nil)
	chkfatal(t, err)
	req.Header.Set(name, sendVal)

	resp, err := http.DefaultClient.Do(req)
	chkfatal(t, err)
	defer func() {
		if t.Failed() {
			t.Log("Response:", resp)
		}
	}()

	expectStatus(t, 200, resp.StatusCode)
	body := &echoBody{}
	err = json.NewDecoder(resp.Body).Decode(body)
	chkfatal(t, err)
	gotRecvVal := body.Headers.Get(name)
	if gotRecvVal != wantRecvVal {
		t.Fatalf("Server did not see correct value for %s header; "+
			"wanted %q but got %q", name, wantRecvVal, gotRecvVal)
	}
}

func TestAcceptHeader(t *testing.T) {
	testHeader(t, "Accept", "application/json", "application/json; q=1")
}

func TestAcceptEncodingHeader(t *testing.T) {
	testHeader(t, "Accept-Encoding", "something-non-standard", "something-non-standard;q=1")
	testHeader(t, "Accept-Encoding", "gzip", "gzip;q=1")
	testHeader(t, "Accept-Encoding", "*;q=1", "*;q=1")
}

func TestAdditionalHeaders(t *testing.T) {
	testHeader(t, "X-Foo-Disallowed-Header", "Bar", "")
	testHeader(t, "X-Sandstorm-App-Baz", "quux", "quux")
	testHeader(t, "OC-Total-Length", "234", "234")
}

func TestETagPrecondition(t *testing.T) {
	testHeader(t, "If-Match", "*", "*")
	testHeader(t, "If-None-Match", `"foobarbaz"`, `"foobarbaz"`)
	testHeader(t, "If-None-Match", `W/"foobarbaz"`, `W/"foobarbaz"`)
	testHeader(t, "If-None-Match", `W/"foobarbaz", "quux"`, `W/"foobarbaz", "quux"`)

	// "net/http" won't magically deal with the etag header unless we use
	// http.FileSystem, so even though the following should arguably
	// return "Precondition Failed," we can ignore it to check whether the
	// header is getting through:
	testHeader(t, "If-None-Match", "*", "*")
	testHeader(t, "If-Match", `"foobarbaz"`, `"foobarbaz"`)
	testHeader(t, "If-Match", `W/"foobarbaz"`, `W/"foobarbaz"`)
	testHeader(t, "If-Match", `W/"foobarbaz", "quux"`, `W/"foobarbaz", "quux"`)
}

func TestResponseContentType(t *testing.T) {
	resp, err := http.Get(baseUrl + "echo-request/content-type")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Expected content type application/json but got %q", contentType)
	}
}

func TestResponseContentEncoding(t *testing.T) {
	resp, err := http.Get(baseUrl + "echo-request/content-encoding")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)

	encoding := resp.Header.Get("Content-Encoding")
	if encoding != "identity" {
		t.Fatalf("Expected content encoding identity, but got %q.", encoding)
	}
}

func TestResponseContentLength(t *testing.T) {
	resp, err := http.Get(baseUrl + "content-length")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)

	contentLength := resp.Header.Get("Content-Length")
	reportLength, err := strconv.ParseUint(contentLength, 10, 64)
	chkfatal(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	chkfatal(t, err)

	realLength := uint64(len(body))
	if realLength != reportLength {
		t.Fatal("Content-Header indicated a length of %d, but body was ",
			"actually %d bytes long.", reportLength, realLength)
	}
}

func TestWantStatus(t *testing.T) {
	for _, wantStatus := range []int{200, 201, 202, 204, 205} {
		resp, err := http.Get(fmt.Sprintf(
			"%secho-request/status?want-status=%d", baseUrl, wantStatus,
		))
		chkfatal(t, err)
		expectStatus(t, wantStatus, resp.StatusCode)
	}
}

func TestWantLanguage(t *testing.T) {
	for _, wantLang := range []string{"en-US", "de-DE"} {
		resp, err := http.Get(
			baseUrl + "echo-request/lang?want-lang=" + wantLang,
		)
		chkfatal(t, err)
		expectStatus(t, 200, resp.StatusCode)
		gotLang := resp.Header.Get("Content-Language")
		if wantLang != gotLang {
			t.Fatalf("Unexpected Content-Language; wanted %q but got %q.",
				wantLang, gotLang)
		}
	}
}

func TestNoOpHandler(t *testing.T) {
	resp, err := http.Get(baseUrl + "no-op-handler")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
}

func TestETag(t *testing.T) {
	resp, err := http.Get(baseUrl + "etag")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	etag := resp.Header.Get("ETag")

	if etag != `"sometag"` {
		t.Errorf(`Expected ETag value "sometag", but got %q`, etag)
	}

	resp, err = http.Get(baseUrl + "etag?weak=true")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	etag = resp.Header.Get("ETag")

	if etag != `W/"sometag"` {
		t.Errorf(`Expected ETag value W/"sometag", but got %q`, etag)
	}
}

func TestNoContentBodies(t *testing.T) {
	resp, err := http.Get(baseUrl + "echo-request/status?want-status=205")
	chkfatal(t, err)
	expectStatus(t, 205, resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	chkfatal(t, err)
	if len(data) != 0 {
		t.Fatalf("No-content response returned a body: %q", data)
	}
}
