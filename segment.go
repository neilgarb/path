package path

import "math"

type Segment struct {
	Point1 Point
	Point2 Point
}

// Intersects returns the intersection segment of the two segments `self` and
// `segment`.
//
// If the segments intersect at a point, the segment will be a single point.
// If the segments overlap, the segment will be a line.
// If the two segments don't intersect, nil is returned.
//
// Adapted from http://geomalgorithms.com/a05-_intersect-1.html.
func (self Segment) Intersect(segment Segment) *Segment {
	u := self.Point1.Vector(self.Point2)
	v := segment.Point1.Vector(segment.Point2)
	w := segment.Point1.Vector(self.Point1)
	d := u.Perp(v)
	if math.Abs(d) < EPSILON {
		if u.Perp(w) != 0.0 || v.Perp(w) != 0.0 {
			return nil
		}
		du := u.Dot(u)
		dv := v.Dot(v)
		if du == 0.0 && dv == 0.0 {
			if self.Point1 != segment.Point1 {
				return nil
			}
			return &Segment{self.Point1, self.Point1}
		}
		if du == 0.0 {
			if !self.Contains(self.Point1) {
				return nil
			}
			return &Segment{self.Point1, self.Point1}
		}
		if dv == 0.0 {
			if !self.Contains(segment.Point1) {
				return nil
			}
			return &Segment{segment.Point1, segment.Point1}
		}
		w2 := segment.Point1.Vector(self.Point2)
		var t0, t1 float64
		if v.X != 0.0 {
			t0 = w.X / v.X
			t1 = w2.X / v.X
		} else {
			t0 = w.Y / v.Y
			t1 = w2.Y / v.Y
		}
		if t0 > t1 {
			t0, t1 = t1, t0
		}
		if t0 > 1.0 || t1 < 0.0 {
			return nil
		}
		t0 = math.Max(0.0, t0)
		t1 = math.Min(t1, 1.0)
		if t0 == t1 {
			p := segment.Point1.Add(v.Mult(t0))
			return &Segment{p, p}
		}
		p1 := segment.Point1.Add(v.Mult(t0))
		p2 := segment.Point1.Add(v.Mult(t1))
		return &Segment{p1, p2}
	}
	sI := v.Perp(w) / d
	if sI < 0 || sI > 1 {
		return nil
	}
	tI := u.Perp(w) / d
	if tI < 0 || tI > 1 {
		return nil
	}
	p := self.Point1.Add(u.Mult(sI))
	return &Segment{p, p}
}

// Contains returns true if the given point lies on this line segment.
//
// Adapted from http://stackoverflow.com/a/328122.
func (self Segment) Contains(point Point) bool {
	prod := (point.Y-self.Point1.Y)*(self.Point2.X-self.Point1.X) - (point.X-self.Point1.X)*(self.Point2.Y-self.Point1.Y)
	if math.Abs(prod) > EPSILON {
		return false
	}
	dot := (point.X-self.Point1.X)*(self.Point2.X-self.Point1.X) + (point.Y-self.Point1.Y)*(self.Point2.Y-self.Point1.Y)
	if dot < 0 {
		return false
	}
	length := (self.Point2.X-self.Point1.X)*(self.Point2.X-self.Point1.X) + (self.Point2.Y-self.Point1.Y)*(self.Point2.Y-self.Point1.Y)
	if dot > length {
		return false
	}
	return true
}
