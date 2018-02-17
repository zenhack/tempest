// +build integration_tests

package websession

// Integration tests; these operate against the test app in ../../internal/testapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
)

// Get the base URL for the test app from the environment. if the requisite environment
// variable is not set, skip the test.
func getAppUrl(t *testing.T) string {
	url := os.Getenv("GO_SANDSTORM_TEST_APP")
	if url == "" {
		t.Fatal("Integration test: GO_SANDSTORM_TEST_APP environment " +
			"variable not defined.")
	}
	return url
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

func TestBasicGetHead(t *testing.T) {
	baseUrl := getAppUrl(t)
	successfulResponse := func(resp *http.Response, err error) {
		if err != nil {
			t.Fatal("Making http request:", err)
		}
		if resp.StatusCode != 200 {
			t.Log("Response:", resp)
			t.Fatal("Got non-200 status code from app")
		}
	}
	successfulResponse(http.Get(baseUrl + "echo-request/hello"))
	successfulResponse(http.Head(baseUrl + "echo-request/hello"))
}

func TestGetCorrectInfo(t *testing.T) {
	baseUrl := getAppUrl(t)

	resp, err := http.Get(baseUrl + "echo-request/hello")
	chkfatal(t, err)
	defer func() {
		if t.Failed() {
			t.Log("Response:", resp)
		}
	}()
	if resp.StatusCode != 200 {
		t.Fatal("Non-zero status code")
	}
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
	baseUrl := getAppUrl(t)

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

	if resp.StatusCode != 200 {
		t.Fatal("Bad status code.")
	}
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

func TestWantStatus(t *testing.T) {
	baseUrl := getAppUrl(t)

	for _, wantStatus := range []int{200, 201, 202, 204, 205} {
		resp, err := http.Get(fmt.Sprintf(
			"%secho-request/status?want-status=%d", baseUrl, wantStatus,
		))
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != wantStatus {
			t.Error("Unexpected status code; wanted", wantStatus, "but got",
				resp.StatusCode, ".")
		}
	}
}

func TestNoOpHandler(t *testing.T) {
	baseUrl := getAppUrl(t)

	resp, err := http.Get(baseUrl + "no-op-handler")
	chkfatal(t, err)
	if resp.StatusCode != 200 {
		t.Fatal("Non-200 status. Response:", resp)
	}
}
