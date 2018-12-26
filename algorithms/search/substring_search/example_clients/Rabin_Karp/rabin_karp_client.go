package main

import (
	"flag"
	"fmt"

	substring_search "github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/search/substring_search"
)

var text string
var pattern string

func init() {
	const (
		defaultText    = ""
		usageText      = "Text used for to search for given pattern"
		defaultPattern = ""
		usagePattern   = "Pattern used to find substring in text"
	)

	flag.StringVar(&text, "text", defaultText, usageText)
	flag.StringVar(&text, "t", defaultText, usageText+" (shorthand)")
	flag.StringVar(&pattern, "pattern", defaultPattern, usagePattern)
	flag.StringVar(&pattern, "p", defaultPattern, usagePattern+" (shorthand)")
}
func main() {
	flag.Parse()
	rk := substring_search.RabinKarpConstructor(pattern, uint64(256))
	pos := rk.Search(text)
	if pos == len(text) {
		fmt.Printf("No substring for pattern %v found in text %v\n", pattern, text)
	} else {
		fmt.Println("Text: \t", text)
		fmt.Print("Pattern: ")
		for i := 0; i < pos; i++ {
			fmt.Print(" ")
		}
		fmt.Println(pattern)
	}
}
