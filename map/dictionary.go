package main

//Dictionary is a map type
type Dictionary map[string]string

//DictionaryErr is a string type
type DictionaryErr string

/*
var (
	ErrNotFound  = errors.New("couldn't find the word you search")
	ErrKeyExists = errors.New("found existing key you want to add")
)
*/

//ErrNotFound and ErrKeyExists are defined error message
const (
	ErrNotFound     = DictionaryErr("couldn't find the word you search")
	ErrKeyExists    = DictionaryErr("found existing key you want to add")
	ErrKeyNotExists = DictionaryErr("word to update doesn't exist") //in fact, it is similar with ErrNotFound
)

func (d DictionaryErr) Error() string {
	return string(d) //this method's return type is string rather than self-defined type DictionaryErr, so we need an explicit transformation
}

//Search is a function return map search results, including value and error message
func (d Dictionary) Search(key string) (string, error) {
	value, err := d[key]
	if !err {
		return "", ErrNotFound
	}
	return value, nil
}

//Add can add key and value into map
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}
	return nil
}

//Update change the value of appointed key with new value
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

//Delete function deletes appointed key and its value
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
