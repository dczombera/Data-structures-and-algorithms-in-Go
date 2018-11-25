// Pogram Degrees of Separation finds the degree of separation between two different data points (e.g. persons)
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/breadth_first_search/breadth_first_paths"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/symbol_graph"
)

var delimiter string
var file string
var source string

func init() {
	const (
		defaultDelimiter = "/"
		usageDelimiter   = "Delimiter used in file to separate data"
		defaultFile      = "movies.txt"
		usageFile        = "Filename used to build symbol graph"
		defaultSource    = "Ford, Harrison"
		usageSource      = "Source which is used to find degrees of separation"
	)

	flag.StringVar(&delimiter, "delimiter", defaultDelimiter, usageDelimiter)
	flag.StringVar(&delimiter, "d", defaultDelimiter, usageDelimiter+" (shorthand)")
	flag.StringVar(&file, "file", defaultFile, usageFile)
	flag.StringVar(&file, "f", defaultFile, usageFile+" (shorthand)")
	flag.StringVar(&source, "source", defaultSource, usageSource)
	flag.StringVar(&source, "s", defaultSource, usageSource+" (shorthand)")
}

func main() {
	flag.Parse()
	sg := symbol_graph.NewSymbolGraph(file, delimiter)
	if !sg.Contains(source) {
		log.Fatalf("Could not find source %v in file\n", source)
	}

	i, err := sg.IndexOf(source)
	check(err)

	paths := breadth_first_paths.NewBreadthFirstPaths(sg.Graph(), i)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter individual to find path for (/q to quit): ")
		scanner.Scan()
		sink := scanner.Text()
		if sink == "/q" {
			break
		}

		if !sg.Contains(sink) {
			fmt.Printf("Could not find '%v' in file\n", sink)
			continue
		}

		j, err := sg.IndexOf(sink)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !paths.HasPathTo(j) {
			fmt.Printf("%v not connected to %v\n", sink, source)
			continue
		}

		path, err := paths.PathTo(j)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for curr := path.First; curr != nil; curr = curr.Next {
			fmt.Print(sg.NameOf(curr.Item))
			if curr.Next != nil {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
