package path

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

// Vector returns the vector `point` - `self`.
func (self Point) Vector(point Point) Vector {
	return Vector{point.X - self.X, point.Y - self.Y}
}

// Segment returns a segment from `self` to `point`.
func (self Point) Segment(point Point) Segment {
	return Segment{self, point}
}

// Distance returns the distance between `self` and `point`.
func (self Point) Distance(point Point) float64 {
	return math.Sqrt(math.Pow(point.X-self.X, 2) + math.Pow(point.Y-self.Y, 2))
}

// Add adds the given vector to this point.
func (self Point) Add(vec Vector) Point {
	return Point{self.X + vec.X, self.Y + vec.Y}
}

func (self Point) String() string {
	return fmt.Sprintf("{%f,%f}", self.X, self.Y)
}
