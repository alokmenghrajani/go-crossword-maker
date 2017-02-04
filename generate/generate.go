package generate

import (
  "fmt"
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
  phaseOne(words, grid)
}

func phaseOne(words *words.Words, g *grid.Grid) {
  for _, word := range words.GetWords() {
    fmt.Printf("trying to place: %s\n", word)
    for i:=0; i<=g.Size - len(word); i++ {
      for j:=0; j<=g.Size - len(word); j++ {
        if n, ok := g.Place(i, j, grid.DOWN, word); ok {
          fmt.Println("success")
          fmt.Println(n)
        } else {
          fmt.Println("fail!")
        }
      }
    }
  }
}
