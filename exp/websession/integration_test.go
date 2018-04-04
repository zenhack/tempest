// +build integration_tests

package websession

// Integration tests; these operate against the test app in ../../internal/testapp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var baseUrl string

// http.Client for use in the tests; we set a timeout because hanging is a
// common failure mode.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

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
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func expectStatus(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Fatal("Unexpected status code; expected", want, "but got", got)
	}
}

func TestBasicGetHead(t *testing.T) {
	successfulResponse := func(resp *http.Response, err error) {
		chkfatal(t, err)
		expectStatus(t, 200, resp.StatusCode)
	}
	successfulResponse(httpClient.Get(baseUrl + "echo-request/hello"))
	successfulResponse(httpClient.Head(baseUrl + "echo-request/hello"))
}

func TestGetCorrectInfo(t *testing.T) {

	resp, err := httpClient.Get(baseUrl + "echo-request/hello")
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

	resp, err := httpClient.Do(req)
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

func TestAdditionalRequestHeaders(t *testing.T) {
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

func TestAdditionalResponseHeaders(t *testing.T) {
	resp, err := httpClient.Post(baseUrl+"return-headers", "application/json", strings.NewReader(`
		{
			"X-Sandstorm-App-Foo": "bar",
			"X-Sandstorm-App-Baz": "quux",
			"X-OC-MTime": "2341"
		}
	`))
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)

	chkHeader := func(name, want string) {
		got := resp.Header.Get(name)
		if got != want {
			t.Errorf("Incorrect value for header %q; wanted %q but got %q.",
				name, want, got)
		}
	}

	chkHeader("X-Sandstorm-App-Foo", "bar")
	chkHeader("X-Sandstorm-App-Baz", "quux")
	chkHeader("X-OC-MTime", "2341")
}

func TestResponseContentType(t *testing.T) {
	resp, err := httpClient.Get(baseUrl + "echo-request/content-type")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Expected content type application/json but got %q", contentType)
	}
}

func TestResponseContentEncoding(t *testing.T) {
	resp, err := httpClient.Get(baseUrl + "echo-request/content-encoding")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)

	encoding := resp.Header.Get("Content-Encoding")
	if encoding != "identity" {
		t.Fatalf("Expected content encoding identity, but got %q.", encoding)
	}
}

func TestResponseContentLength(t *testing.T) {
	resp, err := httpClient.Get(baseUrl + "content-length")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)

	contentLength := resp.Header.Get("Content-Length")
	reportLength, err := strconv.ParseUint(contentLength, 10, 64)
	chkfatal(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	chkfatal(t, err)

	realLength := uint64(len(body))
	if realLength != reportLength {
		t.Fatalf("Content-Header indicated a length of %v, but body was "+
			"actually %v bytes long.", reportLength, realLength)
	}
}

func TestWantStatus(t *testing.T) {
	statusCodes := []int{
		// content:
		200, 201, 202,
		// noContent:
		204, 205,
		// preconditionFailed:
		412,
		// clientError
		400, 403, 404, 405, 406, 409, 410, 413, 414, 415, 418, 422,
		// serverError
		500,
	}
	for _, wantStatus := range statusCodes {
		url := fmt.Sprintf("%secho-request/status?want-status=%d", baseUrl, wantStatus)

		for _, method := range []string{"GET", "HEAD", "DELETE"} {
			req, err := http.NewRequest(method, url, nil)
			chkfatal(t, err)
			resp, err := httpClient.Do(req)
			chkfatal(t, err)
			expectStatus(t, wantStatus, resp.StatusCode)
		}
	}
}

// The server should translate any unrecognized 4xx statuses to 400.
func TestBadClientStatus(t *testing.T) {
	for _, wantStatus := range []int{401, 433, 407, 402} {
		resp, err := httpClient.Get(fmt.Sprintf(
			"%secho-request/status?want-status=%d", baseUrl, wantStatus,
		))
		chkfatal(t, err)
		expectStatus(t, 400, resp.StatusCode)
	}
}

// The server should translate any unrecognized non 4xx statuses to 500.
func TestBadOtherStatus(t *testing.T) {
	statusCodes := []int{0, 42, 124, 243, 339, 343, 501, 503, 555, 623}
	for _, wantStatus := range statusCodes {
		resp, err := httpClient.Get(fmt.Sprintf(
			"%secho-request/status?want-status=%d", baseUrl, wantStatus,
		))
		chkfatal(t, err)
		expectStatus(t, 500, resp.StatusCode)
	}
}

func TestWantLanguage(t *testing.T) {
	for _, wantLang := range []string{"en-US", "de-DE"} {
		resp, err := httpClient.Get(
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
	resp, err := httpClient.Get(baseUrl + "no-op-handler")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
}

func TestETag(t *testing.T) {
	resp, err := httpClient.Get(baseUrl + "etag")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	etag := resp.Header.Get("ETag")

	if etag != `"sometag"` {
		t.Errorf(`Expected ETag value "sometag", but got %q`, etag)
	}

	resp, err = httpClient.Get(baseUrl + "etag?weak=true")
	chkfatal(t, err)
	expectStatus(t, 200, resp.StatusCode)
	etag = resp.Header.Get("ETag")

	if etag != `W/"sometag"` {
		t.Errorf(`Expected ETag value W/"sometag", but got %q`, etag)
	}
}

func TestNoContentBodies(t *testing.T) {
	resp, err := httpClient.Get(baseUrl + "echo-request/status?want-status=205")
	chkfatal(t, err)
	expectStatus(t, 205, resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	chkfatal(t, err)
	if len(data) != 0 {
		t.Fatalf("No-content response returned a body: %q", data)
	}
}

func TestNotModified(t *testing.T) {
	sentETag := `"sometag"`
	req, err := http.NewRequest("GET", baseUrl+"etag", nil)
	chkfatal(t, err)
	req.Header.Set("If-None-Match", sentETag)
	resp, err := httpClient.Do(req)
	chkfatal(t, err)
	expectStatus(t, 304, resp.StatusCode)
	receivedETag := resp.Header.Get("ETag")
	if receivedETag != sentETag {
		t.Fatalf("Got 304 status as expected, but etags didn't match: "+
			"sent %q but received %q.", sentETag, receivedETag)
	}
}

func TestRedirects(t *testing.T) {
	// make sure we don't follow redirects:
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	cases := []struct{ send, recv int }{
		{301, 301},
		{302, 303},
		{303, 303},
		{307, 307},
		{308, 308},
	}
	for _, v := range cases {
		url := fmt.Sprintf("%sredirect?want-status=%d", baseUrl, v.send)
		resp, err := client.Get(url)
		chkfatal(t, err)
		expectStatus(t, v.recv, resp.StatusCode)
		wantLoc := "/redirect-landing"
		gotLoc := resp.Header.Get("Location")
		if gotLoc != wantLoc {
			t.Fatalf("Unexpected Location header %q; expected %q.", gotLoc, wantLoc)
		}
	}
}

// Basic test for PUT/POST/PATCH
func TestPRequest(t *testing.T) {
	methods := []string{"PUT", "POST", "PATCH"}

	sendBody := "Hello, World!\n"

	for _, method := range methods {
		req, err := http.NewRequest(
			method,
			baseUrl+"echo-request/p-request",
			strings.NewReader(sendBody),
		)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "text/plain")
		req.Header.Set("Content-Encoding", "identity")
		resp, err := httpClient.Do(req)
		chkfatal(t, err)

		body := &echoBody{}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		chkfatal(t, err)
		defer func() {
			if t.Failed() {
				t.Logf("response body: %q", bodyBytes)
			}
		}()
		err = json.Unmarshal(bodyBytes, body)
		chkfatal(t, err)

		if body.Method != method {
			t.Errorf("Incorrect method; expected %q but got %q.",
				method, body.Method)
		}

		headers := []struct{ key, val string }{
			{"Content-Type", "text/plain"},
			{"Content-Encoding", "identity"},
		}

		for _, header := range headers {
			actual := body.Headers.Get(header.key)
			if actual != header.val {
				t.Errorf("Incorrect %q header: wanted %q but got %q.",
					header.key, header.val, actual)
			}
		}

		if string(body.Body) != sendBody {
			t.Errorf("Incorrect body; expected %q but got %q.",
				sendBody, body.Body)
		}
	}
}

func TestStreaming(t *testing.T) {
	r, w := io.Pipe()

	// If we just pass in the pipe, the client will block until there's data,
	// so we give it something else to read from before we start writing data,
	// via a MultiReader.
	req, err := http.NewRequest("POST", baseUrl+"echo-body",
		io.MultiReader(strings.NewReader("FirstLine\n"), r))

	chkfatal(t, err)
	resp, err := httpClient.Do(req)
	chkfatal(t, err)
	body := bufio.NewReader(resp.Body)

	line, err := body.ReadString('\n')
	chkfatal(t, err)
	t.Logf("line = %q", line)
	if line != "FirstLine\n" {
		t.Fatalf("Wrote %q but read %q", "FirstLine\n", line)
	}

	msg := "Hello, World\n"
	for i := 0; i < 10; i++ {
		_, err = w.Write([]byte(msg))
		chkfatal(t, err)
		line, err := body.ReadString('\n')
		chkfatal(t, err)
		if line != msg {
			t.Fatalf("Wrote %q but read %q", msg, line)
		}
	}
	w.Close()
	b, err := body.ReadByte()
	switch err {
	case io.EOF:
	case nil:
		t.Fatalf("Expected end of file, but read byte %c", b)
	default:
		t.Fatalf("Expected end of file, but got error %q", err)
	}
}
