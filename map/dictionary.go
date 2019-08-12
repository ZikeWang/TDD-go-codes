package main

import "errors"

//Dictionary is a map type I defined
type Dictionary map[string]string

//ErrNotFound is an error message I defined as a variant
var ErrNotFound = errors.New("couldn't find the word you search")

//Search is a function return map search results, including value and error message
func (d Dictionary) Search(key string) (string, error) {
	value, err := d[key]
	if !err {
		return "", ErrNotFound
	}
	return value, nil
}
