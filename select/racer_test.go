package main

import (
	"testing"
)

func TestRacer(t *testing.T) {
	fastURL := "http://www.baidu.com"
	slowURL := "http://www.github.com"

	want := fastURL
	got := Racer(fastURL, slowURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
