package path

const EPSILON = 1e-10

type Vector struct {
	X float64
	Y float64
}

// Perp returns the perp product of the two vectors `self` and `vec`.
func (self Vector) Perp(vec Vector) float64 {
	return self.X*vec.Y - self.Y*vec.X
}

// Dot returns the dot product of the two vectors `self` and `vec`.
func (self Vector) Dot(vec Vector) float64 {
	return self.X*vec.X + self.Y*vec.Y
}

// Mult multiplies this vector's components by `m` and returns a new vector.
func (self Vector) Mult(m float64) Vector {
	return Vector{self.X * m, self.Y * m}
}
