package uptime

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	client, _ = NewClient(&Config{BaseURL: server.URL, Token: "test-token", RateMilliseconds: 20})
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v should be %v", got, want)
	}
}
