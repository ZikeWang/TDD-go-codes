package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 8.0} //focus here, when initialize struct, using {} not ()
	got := Perimeter(rectangle)
	want := 36.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 10.0, height: 8.0}, want: 80.0},
		{name: "Circle", shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 6.0, height: 6.0}, want: 18.0},
	}

	for _, testcase := range areaTest {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.shape.Area()
			if got != testcase.want {
				t.Errorf("%#v, got %f want %f", testcase.shape, got, testcase.want)
			}
		})
	}
}
