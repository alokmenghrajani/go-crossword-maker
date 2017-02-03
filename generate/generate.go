package generate

import (
  "../grid"
  "../words"
)

// TODO: figure out how to process in parallel.

func Generate(words *words.Words, grid *grid.Grid) {
	// grid generation is performed in two steps
	// 1. place a random word in a random position
	// 2. resolve any partial words until there are none left.
	//    print the grid if it's a better grid or backtrack.
	// repeat step 1.
  
}
