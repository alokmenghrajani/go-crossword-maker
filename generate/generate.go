package generate

import (
	"../grid"
	"../words"
	"fmt"
)

// TODO: figure out how to process in parallel.

func Generate(words *words.Words, grid *grid.Grid) {
	// grid generation is performed in three steps
	// 1. place a random word in a random position
	// 2. resolve any partial words until there are none left.
	//    print the grid if it's a better grid or backtrack.
  // 3. place words that intersect existing words
  //    repeat step 2 or step 1.
  //
  // Note: step 2 should prune the search space. It does however
  //       trigger the exploration of duplicate states.
	phaseOne(words, grid, 0)
}

func phaseTwo(words *words.Words, g *grid.Grid, score int) bool {
	// Find all the horizontal partial words
	t := g.PartialDown()
	if len(t) > 0 {
		//fmt.Printf("PartialDown: %s\n", t)
		return false
	}

	t = g.PartialRight()
	if len(t) > 0 {
		//fmt.Printf("PartialRight: %s\n", t)
		return false
	}

	// grid is in good state!
	fmt.Printf("score: %d\n", score)
	fmt.Println(g)

	return phaseOne(words, g, score)
}

func phaseOne(words *words.Words, g *grid.Grid, score int) bool {
	for _, word := range words.GetWords() {
		if words.IsUsed(word) {
			continue
		}
		//fmt.Printf("trying to place: %s\n", word)
		for i := 0; i <= g.Size; i++ {
			for j := 0; j <= g.Size-len(word); j++ {
				a, z, ok := g.Place(i, j, grid.DOWN, word)
				if ok {
					//fmt.Printf("placed '%s' at DOWN %d, %d\n", word, i, j)
					words.MarkUsed(word)
					// recurse
					if phaseTwo(words, g, score+1) {
						return true
					}
					g.Unplace(i, j, grid.DOWN, word, a, z)
					//fmt.Printf("unplaced '%s' at DOWN %d, %d\n", word, i, j)
					words.MarkUnused(word)
				}
			}
		}
		for i := 0; i <= g.Size-len(word); i++ {
			for j := 0; j <= g.Size; j++ {
				a, z, ok := g.Place(i, j, grid.RIGHT, word)
				if ok {
					//fmt.Printf("placed '%s' at RIGHT %d, %d\n", word, i, j)
					words.MarkUsed(word)
					// recurse
					if phaseTwo(words, g, score+1) {
						return true
					}
					g.Unplace(i, j, grid.RIGHT, word, a, z)
					//fmt.Printf("unplaced '%s' at RIGHT %d, %d\n", word, i, j)
					words.MarkUnused(word)
				}
			}
		}
	}
	return false
}
