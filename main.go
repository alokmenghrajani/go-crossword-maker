package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

/**
 * To make a crossword, we start with a list of words. (TODO: ability to indicate a score
 * for being able to place a given word + dependency to only place a word if another
 * word was placed, can help when creating themes).
 * The user also specifies a desired grid size.
 * We then search for the best solution; continuously printing better solutions to stdout.
 * TODO: add a flag to create symmetric grids.
 * TODO: add the ability to start from a specific grid.
 *
 * Data structures:
 * - Grid: each cell is either empty, black or filled with a letter. If it's
 *         filled with a letter, it may or may not be part of a partial word.
 *         The grid will also track the list of partial words.
 * - Words: list of string. Needs to keep track which words have been placed.
 *          Also, needs some kind of hash table to lookup by n-gram.
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
	fmt.Println("...", words, grid)
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
