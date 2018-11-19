package edge_weighted_directed_graph

type DirectedEdge struct {
	from   int
	to     int
	weight float64
}

func NewDirectedEdge(from, to int, weight float64) DirectedEdge {
	return DirectedEdge{from, to, weight}
}

func NewEmptyDirectedEdge() DirectedEdge {
	return DirectedEdge{}
}

func (e DirectedEdge) From() int {
	return e.from
}

func (e DirectedEdge) To() int {
	return e.to
}

func (e DirectedEdge) Weight() float64 {
	return e.weight
}
