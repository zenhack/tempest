// +build integration_tests,xfail

package websession

// Tests which are currently known to fail.

import (
	"encoding/json"
	"mime"
	"net/http"
	"testing"
)

func TestSetCookie(t *testing.T) {
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
	chkfatal(t, err)
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

func TestDisposition(t *testing.T) {
	mkRequest := func(query string) (mediatype string, params map[string]string) {
		resp, err := http.Get(baseUrl + "echo-request/disposition" + query)
		chkfatal(t, err)
		expectStatus(t, 200, resp.StatusCode)
		disposition := resp.Header.Get("Content-Disposition")
		typ, params, err := mime.ParseMediaType(disposition)
		if err != nil {
			t.Logf("Content-Disposition: %q", disposition)
		}
		chkfatal(t, err)
		return typ, params
	}

	typ, _ := mkRequest("")
	if typ != "inline" {
		t.Fatalf("Unexpected Content-Disposition; wanted inline but got %q", typ)
	}

	typ, params := mkRequest("?filename=hello.txt")
	if typ != "attachment" {
		t.Fatalf("Unexpected Content-Disposition; wanted attachment "+
			"but got %q", typ)
	}
	filename := params["filename"]
	if filename != "hello.txt" {
		t.Fatalf("Unexpected file name in Content-Disposition; wanted hello.txt"+
			" but got %q", filename)
	}
}
