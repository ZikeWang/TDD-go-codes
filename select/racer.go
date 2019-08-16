package main

import (
	"fmt"
	"net/http"
	"time"
)

//ConfigurableRacer selects and returns the faster-visiting URL between two candidates within configurable delays
func ConfigurableRacer(a, b string, delay time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(delay):
		return "", fmt.Errorf("time out waiting for '%s' and '%s'", a, b)
	}
}

//Racer has a fixed delay of 10s
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

//ping uses channel to return url response
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
