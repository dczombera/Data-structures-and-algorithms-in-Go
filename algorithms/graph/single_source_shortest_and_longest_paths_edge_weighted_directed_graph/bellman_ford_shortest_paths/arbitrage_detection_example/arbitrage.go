package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	bfsp "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/single_source_shortest_and_longest_paths_edge_weighted_directed_graph/bellman_ford_shortest_paths"
	graph "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

var file string

func init() {
	const (
		defaultFile = "rates.txt"
		usageFile   = "File containing exchange table"
	)
	flag.StringVar(&file, "file", defaultFile, usageFile)
	flag.StringVar(&file, "f", defaultFile, usageFile+" (shorthand)")
}

func main() {
	flag.Parse()
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error while trying to open file %v: %v", file, err)
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		log.Fatalf("Expected to get number of elements, found empty line")
	}
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("Error while parsing number of entries: %v", err)
	}

	names := make([]string, size)
	g := graph.NewEdgeWeightedDigraph(size)
	for v := 0; v < size; v++ {
		scanLine(scanner)
		tokens := strings.Split(scanner.Text(), " ")
		names[v] = tokens[0]
		for w := 0; w < size; w++ {
			weight, err := strconv.ParseFloat(tokens[w+1], 64)
			if err != nil {
				log.Fatalf("Error while parsing exchange rate: %v", err)
			}
			g.AddEdge(graph.DirectedEdge{v, w, -math.Log(weight)})
		}
	}

	for _, e := range g.Edges() {
		fmt.Println(e)
	}
	cycle := bfsp.NewBellmanFordSP(g, 0)
	if cycle.HasNegativeCycle() {
		amount := 1000.0
		for curr := cycle.NegativeCycle().First; curr != nil; curr = curr.Next {
			e := curr.Item
			fmt.Printf("%.5f %v ", amount, names[e.From])
			amount *= math.Exp(-e.Weight)
			fmt.Printf("= %.5f %v\n", amount, names[e.To])
		}
	} else {
		fmt.Println("No arbitrage opportunity found")
	}
}

func scanLine(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		log.Fatalf("Error while parsing file: Expect line with currency exchange information")
	}
}
