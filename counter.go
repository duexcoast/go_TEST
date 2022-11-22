package main

import "sync"

type Counter struct {
	// sync.Mutex allows us to add locks to our data
	mu sync.Mutex
	value int
}


func(c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func(c *Counter) Value() int {
	return c.value
}