package path

// Space defines a 2D Cartesian plane with origin {0,0}. Walls can be added to
// the space, represented by line segments. The space can then be used to
// determine the shortest path from one point to another around these walls.
type Space struct {
	Walls []Wall
}

// NewSpace returns a new Space.
func NewSpace(walls []Wall) *Space {
	if walls == nil {
		walls = []Wall{}
	}
	return &Space{walls}
}

// AddWall adds a wall to this space.
func (self *Space) AddWall(wall Wall) {
	self.Walls = append(self.Walls, wall)
}

// Path returns a list of control Points which define the shortest path from
// `from` to `to` around this space's walls.
//
// The algorithm makes the handwavy assumption that the path from `from` to
// `to` will either be one line segment (in the case where there is line of
// sight between the two points) or multiple line segments where the control
// points coincide with the edges of one or more walls. There are no curved
// segments to the path, and there is no point in the path which doesn't
// coincide with a wall edge or the destination.
//
// To do this, a weighted digraph is built up whose nodes are `from`, `to` and
// extremeties of all walls.  Nodes in the digraph are connected if there is
// line of sight between them.  The weight of an edge is the 2D distance
// between the points.
//
// The shortest path from `from` to `to` in this graph is taken to be the
// shortest path between the two points in the plane.
func (self Space) ShortestPath(from, to Point) ([]Point, error) {

	// Build a digraph whose nodes are `from`, `to` and the edges of all walls.
	g := NewGraph()
	fromNode := NewNode(from)
	g.AddNode(fromNode)
	toNode := NewNode(to)
	g.AddNode(toNode)
	for _, w := range self.Walls {
		g.AddNode(NewNode(w.Point1))
		g.AddNode(NewNode(w.Point2))
	}

	// Build the connections between nodes. Two nodes are connected if there
	// is line of sight between their points.
	var (
		p1, p2       Point
		s            Segment
		lineOfSight  bool
		intersection *Segment
	)
	for n1, _ := range g.Nodes {
		for n2, _ := range g.Nodes {
			p1 = n1.Payload.(Point)
			p2 = n2.Payload.(Point)
			if p1 == p2 {
				// Don't bother connecting nodes that are the same point.
				continue
			}
			s = p1.Segment(p2)
			lineOfSight = true
			for _, w := range self.Walls {
				intersection = s.Intersect(Segment(w))
				if intersection == nil {
					// There is no intersection, so this wall does not block line of sight.
					continue
				}
				if (*intersection).Point1 != (*intersection).Point2 {
					// This wall is collinear, so does not block line of sight.
					continue
				}
				if (*intersection).Point1 == w.Point1 || (*intersection).Point1 == w.Point2 {
					// The intersection is at a wall extremity, so doesn't block line of sight.
					continue
				}
				// This wall blocks line of sight.
				lineOfSight = false
				break
			}
			if lineOfSight {
				n1.Connect(n2, p1.Distance(p2))
			}
		}
	}

	// Get shorted path and then extract Points
	nodes, err := g.ShortestPath(fromNode, toNode)
	if err != nil {
		return nil, err
	}
	points := make([]Point, len(nodes))
	for i, n := range nodes {
		points[i] = n.Payload.(Point)
	}
	return points, nil
}
