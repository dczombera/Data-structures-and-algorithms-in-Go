package graph

type Bag struct {
	items []int
	Size  int
}

type Graph struct {
	vertices  int
	edges     int
	adjacency []Bag
}

func (b *Bag) add(v int) {
	b.items = append(b.items, v)
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

func (g *Graph) AdjacencyList(vertex int) []int {
	g.validateVertices(vertex)
	return g.adjacency[vertex].items
}

func (g *Graph) Degree(vertex int) int {
	g.validateVertices(vertex)
	return g.adjacency[vertex].Size
}
