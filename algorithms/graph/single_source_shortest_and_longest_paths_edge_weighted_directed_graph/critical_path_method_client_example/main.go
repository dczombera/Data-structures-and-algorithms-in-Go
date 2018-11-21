package main

import "flag"

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
}
