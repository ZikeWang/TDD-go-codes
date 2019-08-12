package main

import "testing"

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

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "testAdd"
		value := "this is an Add function test"
		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertValue(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "testAdd"
		value := "this is an Add function test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new value")
		assertError(t, err, ErrKeyExists)
		assertValue(t, dictionary, key, value)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertValue(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err == ErrNotFound {
		t.Errorf("add failed")
	}

	if got != value {
		t.Errorf("got '%s' want '%s'", got, value)
	}
}
