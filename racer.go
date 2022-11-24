package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {

	// select lets you wait on multiple channels. The first one to send a value "wins"
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// chan struct{} is the smallest data type available from a memory perspective
	// we get no allocation versus a bool. we're not sending anything, so why allocate?
	// always make channels versus assigning to a var. var ch chan struct{} initializes with a zero value,
	// which would be nil in this case.
	// If we try to send to a nil channel with <-, it will block forever bc you can't send to nil channels.
	ch := make(chan struct{})
	go func() {
		// sends signal into channel once we have compelted http.Get
		http.Get(url)
		close(ch)
	}()
	return ch
}
