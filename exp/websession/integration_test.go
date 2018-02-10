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

func TestAcceptHeader(t *testing.T) {
	baseUrl := getAppUrl(t)

	req, err := http.NewRequest("GET", baseUrl+"echo-request/acceptheader", nil)
	chkfatal(t, err)
	req.Header.Set("Accept", "application/json")

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
	sawAccept := body.Headers.Get("Accept")
	if sawAccept != "application/json; q=1" {
		t.Fatalf("Server did not see correct value for Accept header; "+
			"wanted \"application/json; q=1\" but got %q", sawAccept)
	}
}

func TestSetCookie(t *testing.T) {
	baseUrl := getAppUrl(t)

	resp, err := http.Get(baseUrl + "set-cookie")
	chkfatal(t, err)
	cookies := resp.Cookies()
	defer func() {
		if t.Failed() {
			t.Log("Response:", resp)
			t.Log("Cookies:", cookies)
		}
	}()
	if len(cookies) != 1 {
		t.Fatal("Wrong number of cookies; expected 1 but got", len(cookies))
	}
	if cookies[0].Name != "bob" {
		t.Error("Wrong cookie name:", cookies[0].Name)
	}
	if cookies[0].Value != "chocolate chip" {
		t.Error("Wrong cookie name:", cookies[0].Value)
	}
}

func TestSendCookie(t *testing.T) {
	baseUrl := getAppUrl(t)

	req, err := http.NewRequest("GET", baseUrl+"echorequest/sendcookie", nil)
	chkfatal(t, err)

	req.AddCookie(&http.Cookie{
		Name:   "testcookie",
		Value:  "milk_and",
		Secure: true,
	})

	t.Log(req)

	resp, err := http.DefaultClient.Do(req)
	chkfatal(t, err)
	body := &echoBody{}
	err = json.NewDecoder(resp.Body).Decode(body)
	found := false
	for _, cookie := range body.Cookies {
		if cookie.Name == "test-cookie" && cookie.Value == "milk and" {
			found = true
			break
		}
	}
	if !found {
		t.Log("cookie header:", body.Headers.Get("Cookie"))
		t.Log(body)
		t.Fatal("The cookie we sent was not found!")
	}
}
