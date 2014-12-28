package path

type BoundingBox struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

// NewBoundingBox returns a new BoundingBox representing the square bounding
// box of the given points on a 2D Cartesian plane.
func NewBoundingBox(points []Point) BoundingBox {
	bbox := BoundingBox{}
	i := 0
	for _, v := range points {
		if i == 0 || v.Y > bbox.Top {
			bbox.Top = v.Y
		}
		if i == 0 || v.Y < bbox.Bottom {
			bbox.Bottom = v.Y
		}
		if i == 0 || v.X < bbox.Left {
			bbox.Left = v.X
		}
		if i == 0 || v.X > bbox.Right {
			bbox.Right = v.X
		}
		i++
	}
	return bbox
}

// Contains returns true if the given point is in this bounding box.
func (self BoundingBox) Contains(point Point) bool {
	return point.X >= self.Left && point.X <= self.Right && point.Y >= self.Bottom && point.Y <= self.Top
}
