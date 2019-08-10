package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	Judgeequal := func(t *testing.T, got, want []int) {
		//slice can only be compared to nil
		//But reflect.DeepEqual is not type safe, so be careful here
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test function SumIndividual", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SumIndividual(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("test function SunAll", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5, 6}

		got := SumAll(numbers1, numbers2)
		want := []int{6, 15}

		Judgeequal(t, got, want)
	})

	t.Run("test function SumAllTails", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 9}

		Judgeequal(t, got, want)
	})

	t.Run("test function SumAllTails with edge examples", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{3, 4, 5}

		got := SumAllTails(numbers1, numbers2)
		want := []int{0, 9}

		Judgeequal(t, got, want)
	})
}
