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
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 8.0}, 80.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{6.0, 6.0}, 18.0},
	}

	for _, testcase := range areaTest {
		got := testcase.shape.Area()
		if got != testcase.want {
			t.Errorf("got %f want %f", got, testcase.want)
		}
	}
}
