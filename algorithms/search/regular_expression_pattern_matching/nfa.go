package regular_expression_pattern_matching

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/depth_first_search/directed_dfs"
	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/directed_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

type NFA struct {
	re      []byte
	digraph *graph.Digraph
}

func NFAConstructor(pattern string) NFA {
	re := []byte(pattern)
	m := len(pattern)
	digraph := graph.NewDigraph(m + 1)
	ops := stack.NewStack()

	for i := 0; i < m; i++ {
		lp := i
		if re[i] == '(' || re[i] == '|' {
			ops.Push(i)
		} else if re[i] == ')' {
			or, _ := ops.Pop()
			if re[or] == '|' {
				lp, _ = ops.Pop()
				digraph.AddEdge(lp, or+1)
				digraph.AddEdge(or, i)
			} else {
				lp = or
			}
		}

		if i < m-1 && re[i+1] == '*' {
			digraph.AddEdge(lp, i+1)
			digraph.AddEdge(i+1, lp)
		}
		if re[i] == '(' || re[i] == ')' || re[i] == '*' {
			digraph.AddEdge(i, i+1)
		}
	}
	return NFA{re, digraph}
}

func (nfa NFA) Recognizes(txt string) bool {
	states := make([]int, 0, len(nfa.re))
	dfs := directed_dfs.NewDirectedDFS(nfa.digraph, 0)
	for i := 0; i < nfa.digraph.Vertices(); i++ {
		if dfs.IsConnected(i) {
			states = append(states, i)
		}
	}
	for i := 0; i < len(txt); i++ {
		matches := make([]int, 0, len(nfa.re))
		for _, s := range states {
			// Make sure we don't use accept state as array index
			// since its value is len(nfa.re) and therefore would be out of bounds
			if s < len(nfa.re) {
				if nfa.re[s] == txt[i] || nfa.re[s] == '.' {
					matches = append(matches, s+1)
				}
			}
		}
		dfs = directed_dfs.NewDirectedDFS(nfa.digraph, matches...)
		states = make([]int, 0, len(nfa.re))
		for i := 0; i < nfa.digraph.Vertices(); i++ {
			if dfs.IsConnected(i) {
				states = append(states, i)
			}
		}
	}
	for _, s := range states {
		if s == len(nfa.re) {
			return true
		}
	}
	return false
}
