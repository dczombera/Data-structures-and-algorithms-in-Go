package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/single_source_shortest_and_longest_paths_edge_weighted_directed_graph/critical_path_method"
)

var file string

func init() {
	const (
		defaultFile = "jobs.txt"
		usageFile   = "Filename used to build an edge weighted directed graph used solve the job-scheduling problem"
	)

	flag.StringVar(&file, "file", defaultFile, usageFile)
	flag.StringVar(&file, "f", defaultFile, usageFile+" (shorthand)")
}

func main() {
	flag.Parse()
	cpm := critical_path_method.NewCPM(file)
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		log.Fatalln("EOF")
	}

	jobs, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("jobs\t\tstart\t\t finish")
	fmt.Println("***************************************")
	for i := 0; i < jobs; i++ {
		fmt.Printf("%v\t\t%.1f\t\t%.1f\n", i, cpm.StartTimeOf(i), cpm.FinishTimeOf(i))
	}
	fmt.Println("***************************************")
	fmt.Printf("Overall finish time: %.1f\n", cpm.OverallFinishTime())
}
