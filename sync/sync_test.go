package main

import "testing"

func TestCount(t *testing.T) {
	t.Run("basic method running test", func(t *testing.T) {
		iterations := 3
		counter := Counter{}
		for i := 0; i < iterations; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, iterations)
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
