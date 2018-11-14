package edge

import (
	"fmt"
)

type Edge struct {
	eitherVertex int
	otherVertex  int
	weight       float64
}

func NewEdge(v, w int, weight float64) Edge {
	return Edge{v, w, weight}
}

func NewEmptyEdge() Edge {
	return Edge{0, 0, 0.0}
}

func (e Edge) Either() int {
	return e.eitherVertex
}

func (e Edge) Other(vertex int) int {
	if vertex == e.eitherVertex {
		return e.otherVertex
	} else if vertex == e.otherVertex {
		return e.eitherVertex
	}

	panic(fmt.Sprintf("Vertex %v is not a valid vertex for this edge", vertex))
}

func (e Edge) Weight() float64 {
	return e.weight
}

func (this Edge) Compare(other Edge) int {
	if this.weight > other.weight {
		return 1
	} else if this.weight < other.weight {
		return -1
	}
	return 0
}
