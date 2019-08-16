package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := makeDelayServer(0 * time.Millisecond)
	slowServer := makeDelayServer(20 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	want := fastURL
	got := Racer(fastURL, slowURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

//makeDelayServer create a new mock http server with alternative delay time
func makeDelayServer(delayTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayTime)
		w.WriteHeader(http.StatusOK)
	}))
}
