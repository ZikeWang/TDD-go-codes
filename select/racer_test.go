package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	want := fastURL
	got := Racer(fastURL, slowURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	fastServer.Close()
	slowServer.Close()
}
