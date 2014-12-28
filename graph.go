package path

import (
	"fmt"
	"math"
)

type Node struct {
	Payload interface{}
	Edges   map[*Node]float64
}

// NewNode returns a new Node.
func NewNode(payload interface{}) *Node {
	return &Node{payload, make(map[*Node]float64, 0)}
}

// Connect adds an edge to the given node with the given cost.
func (self *Node) Connect(node *Node, cost float64) {
	self.Edges[node] = cost
}

func (self *Node) String() string {
	return fmt.Sprintf("%v", self.Payload)
}

type Graph struct {
	Nodes map[*Node]bool
}

// NewGraph returns a new Graph
func NewGraph() Graph {
	return Graph{make(map[*Node]bool, 0)}
}

// AddNode adds a node to the graph. The same node cannot be added twice.
func (self Graph) AddNode(node *Node) {
	self.Nodes[node] = true
}

// Use Dijkstra's algorithm to find the shortest path from `from` to `to`.
//
// Adapted from http://www.cse.ust.hk/~dekai/271/notes/L10/L10.pdf.
func (self Graph) ShortestPath(from *Node, to *Node) ([]*Node, error) {
	costs := map[*Node]float64{}
	pred := map[*Node]*Node{from: nil}
	pqueue := map[*Node]float64{}
	for n, _ := range self.Nodes {
		costs[n] = math.Inf(1)
		pqueue[n] = math.Inf(1)
	}
	costs[from] = 0.0
	pqueue[from] = 0.0
	var (
		min     float64
		minNode *Node
	)
	for len(pqueue) > 0 {
		// Pick min from queue.
		min = math.Inf(1)
		minNode = nil
		for n, c := range pqueue {
			if c < min {
				min = c
				minNode = n
			}
		}

		// Iterate over minNode's edges
		for n, c := range minNode.Edges {
			if costs[minNode]+c < costs[n] {
				costs[n] = costs[minNode] + c
				pqueue[n] = costs[n]
				pred[n] = minNode
			}
		}

		// Remove this node from the queue - it's visited.
		delete(pqueue, minNode)
	}

	nodes := []*Node{}
	current := to
	for current != nil {
		nodes = append([]*Node{current}, nodes...)
		current = pred[current]
	}

	return nodes, nil
}
