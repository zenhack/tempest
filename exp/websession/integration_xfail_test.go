// +build integration_tests,xfail

package websession

import (
	"encoding/json"
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
