package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter three times leaves it at 3", func(t *testing.T) {
		counter := newCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := newCounter()

		// sync.WaitGroup allows us to wait for goroutines to finish jobs. 
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			// create 1000 threads and decrement wg by one upon completion of each one
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		// wait until the WaitGroup counter is zero
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d wanted %d", got.Value(), want)
	}
}

func newCounter() *Counter {
	return &Counter{}
}
