package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func Example1_Add() {
	sum := Add(0, 5)
	fmt.Println(sum)
	//Output: 5
}

func Example2_Add() {
	sum := Add(4, 6)
	fmt.Println(sum)
	//Output: 10
}
