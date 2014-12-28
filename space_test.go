package path

import (
	"reflect"
	"testing"
)

func TestSpaceShortestPath(t *testing.T) {

	type Test struct {
		space    *Space
		from     Point
		to       Point
		expected []Point
	}

	tests := []Test{
		{NewSpace([]Wall{{Point{0, 0}, Point{10, 10}}}), Point{2, 5}, Point{5, 2}, []Point{Point{2, 5}, Point{0, 0}, Point{5, 2}}},
	}

	for _, test := range tests {
		result, err := test.space.ShortestPath(test.from, test.to)
		if err != nil {
			t.Errorf(`Expected nil, got "%v".`, err)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf(`Expected %v, got %v.`, test.expected, result)
		}
	}
}
