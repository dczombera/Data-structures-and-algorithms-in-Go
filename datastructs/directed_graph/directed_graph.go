package directed_graph

type Digraph struct {
	vertices  int
	edges     int
	adjacency [][]int
	indegree  []int
}

func NewDigraph(sizeV int) *Digraph {
	dg := Digraph{sizeV, 0, make([][]int, sizeV), make([]int, sizeV)}
	return &dg
}

func (dg *Digraph) AddEdge(v, w int) {
	dg.validateVertices(v, w)
	dg.adjacency[v] = append(dg.adjacency[v], w)
	dg.indegree[w]++
	dg.edges++
}

func (dg *Digraph) AdjacencyList(v int) []int {
	dg.validateVertex(v)
	return dg.adjacency[v]
}

func (dg *Digraph) Vertices() int {
	return dg.vertices
}

func (dg *Digraph) Edges() int {
	return dg.edges
}

func (dg *Digraph) Indegree(v int) int {
	dg.validateVertex(v)
	return dg.indegree[v]
}

func (dg *Digraph) Outdegree(v int) int {
	dg.validateVertex(v)
	return len(dg.adjacency[v])
}

func (dg *Digraph) Reverse() *Digraph {
	size := dg.Vertices()
	reverseDg := Digraph{size, dg.Edges(), make([][]int, size), make([]int, size)}
	for v := 0; v < dg.Vertices(); v++ {
		for _, w := range dg.adjacency[v] {
			reverseDg.AddEdge(w, v)
		}
	}

	return &reverseDg
}

func (dg *Digraph) validateVertices(vertices ...int) {
	for _, v := range vertices {
		dg.validateVertex(v)
	}
}

func (dg *Digraph) validateVertex(v int) {
	if v < 0 || v >= dg.Vertices() {
		panic("Vertex out of bounds")
	}
}
