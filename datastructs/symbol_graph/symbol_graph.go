package symbol_graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/graph"
)

type SymbolGraph struct {
	keys        map[string]int
	reverseKeys []string
	graph       *graph.Graph
}

var initCap = 16

func NewSymbolGraph(filename string, delimiter string) *SymbolGraph {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sg := SymbolGraph{make(map[string]int), nil, nil}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), delimiter)
		for _, t := range tokens {
			if !sg.Contains(t) {
				sg.keys[t] = len(sg.keys)
			}
		}
	}

	sg.reverseKeys = make([]string, len(sg.keys))
	for k, v := range sg.keys {
		sg.reverseKeys[v] = k
	}

	g := graph.NewGraph(len(sg.keys))
	// Rewind reader in order to avoid creation of new file
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(f)
	// connect first vertex on each line with all other on same line
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), delimiter)
		v := sg.keys[tokens[0]]
		for _, t := range tokens[1:] {
			g.AddEdge(v, sg.keys[t])
		}
	}
	sg.graph = &g

	return &sg
}

func (sg *SymbolGraph) Contains(s string) bool {
	_, ok := sg.keys[s]
	return ok
}

func (sg *SymbolGraph) IndexOf(s string) (int, error) {
	i, ok := sg.keys[s]
	if !ok {
		return -1, errors.New(fmt.Sprintf("vertex with name %v does not exist", s))
	}
	return i, nil
}

func (sg *SymbolGraph) NameOf(i int) string {
	sg.validateVertex(i)
	return sg.reverseKeys[i]
}

func (sg *SymbolGraph) Graph() *graph.Graph {
	return sg.graph
}

func (sg *SymbolGraph) validateVertex(v int) {
	if v < 0 || v >= sg.graph.Vertices() {
		panic(fmt.Sprintf("Vertex %v out of bounds", v))
	}
}
