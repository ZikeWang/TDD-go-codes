package context3

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// T1 means cancel function invoke latency
var T1 = 5 * time.Millisecond

// T2 means go func execution latency
var T2 = 8 * time.Millisecond

// T3 means select default execution latency
var T3 = 10 * time.Millisecond

// Store fetches data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns a handler for calling Store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error whatever you like
		}

		fmt.Fprint(w, data)
	}
}
