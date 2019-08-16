package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("comapare speed of server", func(t *testing.T) {
		fastServer := makeDelayServer(1 * time.Millisecond)
		slowServer := makeDelayServer(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got, err := Racer(fastURL, slowURL)

		if err != nil {
			t.Error("got an error but not expect one")
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test slow timeout", func(t *testing.T) {
		server := makeDelayServer(30 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("want an error but didn't get one")
		}
	})
}

//makeDelayServer create a new mock http server with alternative delay time
func makeDelayServer(delayTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayTime)
		w.WriteHeader(http.StatusOK)
	}))
}
