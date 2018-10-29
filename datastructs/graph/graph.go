package graph

type Node struct {
	val  Value
	next *Node
}

type Bag struct {
	first *Node
	Size  int
}

type Graph struct {
	vertices  int
	edges     int
	adjacency []Bag
}

type Value int

func NewNode(val Value, next *Node) *Node {
	return &Node{val, next}
}

func (b *Bag) add(v int) {
	n := NewNode(Value(v), b.first)
	b.first = n
	b.Size++
}

func NewGraph(sizeV int) Graph {
	return Graph{sizeV, 0, make([]Bag, sizeV)}
}

func (g *Graph) AddEdge(v, w int) {
	g.validateVertices(v, w)
	g.adjacency[v].add(w)
	g.adjacency[w].add(v)
}

func (g *Graph) validateVertices(vertices ...int) {
	for _, v := range vertices {
		if v < 0 || v >= g.vertices {
			panic("vertix out of graph bounds")
		}
	}
}

func (g *Graph) Vertices() int {
	return g.vertices
}

func (g *Graph) Edges() int {
	return g.edges
}

func (g *Graph) AdjacencyList(vertex int) Bag {
	g.validateVertices(vertex)
	return g.adjacency[vertex]
}

func (g *Graph) Degree(vertex int) int {
	g.validateVertices(vertex)
	return g.adjacency[vertex].Size
}
