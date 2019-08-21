package main

//Counter is a struct which has a int field to record numbers
type Counter struct {
	value int
}

//Inc is a method of Counter to increment the value one by one
func (c *Counter) Inc() {
	c.value++
}

//Value is a method of Counter to echo value
func (c *Counter) Value() int {
	return c.value
}
