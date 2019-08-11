package perimeter

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
	checkshape := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area() //calling method of struct instead of function using interface
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("test rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 8.0}
		want := 80.0
		checkshape(t, rectangle, want)
	})

	t.Run("test circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkshape(t, circle, want)
	})
}
