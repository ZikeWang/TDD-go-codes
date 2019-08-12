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

func TestUpdate(t *testing.T) {
	t.Run("Exist word", func(t *testing.T) {
		key := "testUpdate"
		value := "this is a Update function test"
		want := "new value"

		dictionary := Dictionary{key: value}
		err := dictionary.Update(key, want)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, want)
	})

	t.Run("new word", func(t *testing.T) {
		key := "testUpdate"
		want := "new value"

		dictionary := Dictionary{}
		err := dictionary.Update(key, want)

		assertError(t, err, ErrKeyNotExists)
	})
}

func TestDelete(t *testing.T) {
	key := "testDelete"
	value := "this is a Delete function test"

	dictionary := Dictionary{key: value}
	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", key)
	}
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
