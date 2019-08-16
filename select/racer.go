package main

import (
	"net/http"
	"time"
)

//Racer selects and returns the faster-visiting URL between two candidates
func Racer(a, b string) (winner string) {
	aDuration := measureDurationTime(a)
	bDuration := measureDurationTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

//measureDurationTime calculates response time to an url
func measureDurationTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
