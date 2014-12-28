package path

import "testing"

func TestSegmentIntersects(t *testing.T) {
	type Test struct {
		segment1 Segment
		segment2 Segment
		expected *Segment
	}
	tests := []Test{
		// Perpendicular, bisecting
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{0, 10}, Point{10, 0}}, &Segment{Point{5, 5}, Point{5, 5}}},
		// Touching at one end
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{10, 10}, Point{20, 0}}, &Segment{Point{10, 10}, Point{10, 10}}},
		// Collinear, touching
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{10, 10}, Point{20, 20}}, &Segment{Point{10, 10}, Point{10, 10}}},
		// Collinear, not touching
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{11, 11}, Point{20, 20}}, nil},
		// Collinear, touching
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{-10, -10}, Point{0, 0}}, &Segment{Point{0, 0}, Point{0, 0}}},
		// Collinear, not touching
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{-10, -10}, Point{-1, -1}}, nil},
		// Collinear, overlapping
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{0, 0}, Point{5, 5}}, &Segment{Point{0, 0}, Point{5, 5}}},
		// Parallel
		{Segment{Point{0, 0}, Point{10, 10}}, Segment{Point{1, 0}, Point{11, 10}}, nil},
	}
	var result *Segment
	for _, test := range tests {
		result = test.segment1.Intersect(test.segment2)
		if test.expected == nil && result != nil {
			t.Errorf(`Expected nil, got %v.`, result)
		} else if test.expected != nil && result == nil {
			t.Errorf(`Expected %v, got nil.`, *test.expected)
		} else if test.expected != nil && *result != *test.expected {
			t.Errorf(`Expected %v, got %v.`, *test.expected, *result)
		}
	}
}

func TestSegmentContains(t *testing.T) {
	type Test struct {
		point    Point
		expected bool
	}
	tests := []Test{
		{Point{0, 0}, true},
		{Point{5, 5}, true},
		{Point{10, 10}, true},
		{Point{2, 5}, false},
		{Point{5, 2}, false},
		{Point{11, 11}, false},
		{Point{-1, -1}, false},
	}
	s := Segment{Point{0, 0}, Point{10, 10}}
	var result bool
	for _, test := range tests {
		result = s.Contains(test.point)
		if result != test.expected {
			t.Errorf(`Expected %t, got %t.`, test.expected, result)
		}
	}
}
