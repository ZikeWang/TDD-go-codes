package main

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("basic struct's methods running test", func(t *testing.T) {
		iterations := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(iterations)

		for i := 0; i < iterations; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, iterations)
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
