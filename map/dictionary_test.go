package main

import "testing"

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func TestMapSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"Test": "It is just a test"}
		got, _ := dictionary.Search("Test")
		want := "It is just a test"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("Test")
		if err == nil {
			t.Fatal("want an error but didn't got one")
		}
		assertError(t, err, ErrNotFound)
	})
}
