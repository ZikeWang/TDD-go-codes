package main

import "sync"

//Counter is a struct which has a int field to record numbers
type Counter struct {
	mu    sync.Mutex
	value int
}

//Inc is a method of Counter to increment the value one by one
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

//Value is a method of Counter to echo value
func (c *Counter) Value() int {
	return c.value
}
