package websession

// Integration tests; these operate against the test app in ../../internal/testapp

import (
	"encoding/json"
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
		t.SkipNow()
	}
	return url
}

type echoBody struct {
	Method  string              `json:"method"`
	Url     *url.URL            `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
	Cookies []*http.Cookie      `json:"cookies"`
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

func TestCorrectInfo(t *testing.T) {
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
