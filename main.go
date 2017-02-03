package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

/**
 * To make a crossword, we start with a list of words.
 * We then search for the best solution; continuously printing better solutions to stdout.
 * TODO: add a flag to create symmetric grids.
 * TODO: add the ability to start from a specific grid.
 */

var (
	words = kingpin.Flag("words", "File to read word list from.").Short('w').Required().String()
	size  = kingpin.Flag("size", "Grid size.").Short('s').Required().Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Loading %s\n", *words)
	words := Load(*words)
	grid := New(*size)
	fmt.Println(grid)
	_ = words
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func panicIfFalse(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}
