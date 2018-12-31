package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/regular_expression_pattern_matching"
)

var file string
var pattern string

func init() {
	const (
		defaultFile    = "text.txt"
		usageFile      = "Text file used to search for given regular expression pattern"
		defaultPattern = ""
		usagePattern   = "Regular expression pattern used to find substring in text"
	)

	flag.StringVar(&file, "file", defaultFile, usageFile)
	flag.StringVar(&file, "f", defaultFile, usageFile+" (shorthand)")
	flag.StringVar(&pattern, "regex", defaultPattern, usagePattern)
	flag.StringVar(&pattern, "r", defaultPattern, usagePattern+" (shorthand)")
	flag.StringVar(&pattern, "p", defaultPattern, usagePattern+" (shorthand)")
}

func main() {
	flag.Parse()
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error while trying to open file %v: %v", file, err)
	}

	nfa := regular_expression_pattern_matching.NFAConstructor("(.*" + pattern + ".*)")
	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		line++
		txt := scanner.Text()
		if nfa.Recognizes(txt) {
			fmt.Printf("%v:\t%v\n", line, strings.TrimSpace(txt))
		}
	}
}
