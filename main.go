package main

import (
	"./generate"
	"./grid"
	"./words"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

/**
 Tasks:
 - grid.go should track which words have been placed instead of generate.go
 - is it worth improving the way words get removed? I.e. undo-stack?
 - better scoring system. Right now, it's simply number of words placed in grid. should
   favor intersections.
 */

/**
 * To generate a crossword, we start with a list of words.
 * We then search for the best placement of words; continuously printing better solutions to stdout.
 * TODO: add a flag to create symmetric grids.
 * TODO: add the ability to start from a specific grid.
 * TODO: setup godep + update README
 * TODO: setup travis
 * TODO: support non-square grids?
 */

var (
	wordlist = kingpin.Flag("wordlist", "File to read word list from.").Short('w').Required().String()
	size     = kingpin.Flag("size", "Grid size.").Short('s').Required().Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Loading %s\n", *wordlist)
	words := words.Load(*wordlist)
	grid := grid.New(*size)
	fmt.Println("Starting search")
	generate.Generate(words, grid)
}
