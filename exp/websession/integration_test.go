package websession

// Integration tests; these operate against the test app in ../../internal/testapp

import (
	"net/http"
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

func TestBasicGet(t *testing.T) {
	baseUrl := getAppUrl(t)
	resp, err := http.Get(baseUrl + "echo-request/hello")
	if err != nil {
		t.Fatal("Making http request:", err)
	}
	t.Log("Response:", resp)
	if resp.StatusCode != 200 {
		t.Fatal("Got non-200 status code from app")
	}
}
