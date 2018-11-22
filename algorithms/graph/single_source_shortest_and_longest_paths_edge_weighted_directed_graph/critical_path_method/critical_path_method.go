// Package critical_path_method implements solves the parallel precedence-constrained job scheduling problem
// via the critical path method
// The constructor expects the name of a file which has following input format
// - First line:  the number of jobs
// - Each following line: duration of jobs, number of precedence constrained jobs, each individual precedence constrained job
// Example file content:
// 10
// 41.0 3 1 7 9
// 51.0 1 2
// 50.0 0
// 36.0 0
// 38.0 0
// 45.0 0
// 21.0 2 3 8
// 32.0 2 3 8
// 32.0 1 2
// 29.0 2 4 6

package critical_path_method

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	alp "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/single_source_shortest_and_longest_paths_edge_weighted_directed_graph/acyclic_longest_paths"
	g "github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_directed_graph"
)

type CPM struct {
	source        int
	sink          int
	jobs          int
	criticalPaths alp.AcyclicLP
}

func NewCPM(file string) CPM {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		log.Fatalf("Expected file %v to have jobs number and details for each jobs, got an empty file\n", file)
	}
	jobs, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("Expected number of jobs as first line in file %v, got parsing error %v\n", file, err)
	}

	source := jobs * 2
	sink := jobs*2 + 1
	graph := g.NewEdgeWeightedDigraph(jobs*2 + 2)
	for i := 0; i < jobs; i++ {
		scanLine(scanner)
		tokens := strings.Split(scanner.Text(), " ")
		duration, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			log.Fatalf("Expected duration for job %v, got parsing error %v\n", i, err)
		}

		graph.AddEdge(g.DirectedEdge{source, i, 0.0})
		graph.AddEdge(g.DirectedEdge{i, i + jobs, duration})
		graph.AddEdge(g.DirectedEdge{i + jobs, sink, 0.0})

		precedences, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Expected number of precedences for job %v, got error %v\n", i, err)
		}
		for j := 0; j < precedences; j++ {
			precedence, err := strconv.Atoi(tokens[j+2])
			if err != nil {
				log.Fatalf("Expected precedence job for job %v got error %v\n", i, err)
			}
			graph.AddEdge(g.DirectedEdge{i + jobs, precedence, 0.0})
		}
	}

	criticalPaths, err := alp.NewAcyclicLP(graph, source)
	if err != nil {
		log.Fatalln(err)
	}

	return CPM{source, sink, jobs, criticalPaths}
}

func (cpm CPM) StartTimeOf(v int) float64 {
	return cpm.timeOf(v)
}

func (cpm CPM) FinishTimeOf(v int) float64 {
	return cpm.timeOf(v + cpm.jobs)
}

func (cpm CPM) OverallFinishTime() float64 {
	return cpm.timeOf(cpm.sink)
}

func (cpm CPM) timeOf(v int) float64 {
	time, err := cpm.criticalPaths.DistTo(v)
	if err != nil {
		log.Fatalln(err)
	}
	return time
}

func scanLine(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		log.Fatalf("Expected line with job duration and precedence constraints, got empty line")
	}
}
