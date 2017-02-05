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

func phaseTwo(words *words.Words, g *grid.Grid) bool {
  // Find all the horizontal partial words
  t := g.PartialDown()
  if len(t) > 0 {
    fmt.Printf("PartialDown returned something\n")
    fmt.Println(t)
    return false
  }

  t = g.PartialRight()
  if len(t) > 0 {
    fmt.Printf("PartialRight returned something\n")
    fmt.Println(t)
    return false
  }

  // grid is in good state!
  fmt.Println(g)

  return phaseOne(words, g)
}

func phaseOne(words *words.Words, g *grid.Grid) bool {
  for _, word := range words.GetWords() {
    if words.IsUsed(word) {
      continue
    }
    fmt.Printf("trying to place: %s\n", word)
    for i:=0; i<=g.Size; i++ {
      for j:=0; j<=g.Size - len(word); j++ {
        a, z, ok := g.Place(i, j, grid.DOWN, word)
        if ok {
          fmt.Printf("placed '%s' at DOWN %d, %d\n", word, i, j)
          words.MarkUsed(word)
          // recurse
          if phaseTwo(words, g) {
            return true
          }
          g.Unplace(i, j, grid.DOWN, word, a, z)
          fmt.Printf("unplaced '%s' at DOWN %d, %d\n", word, i, j)                    
          words.MarkUnused(word)
        }
      }
    }
    for i:=0; i<=g.Size - len(word); i++ {
      for j:=0; j<=g.Size; j++ {
        a, z, ok := g.Place(i, j, grid.RIGHT, word)
        if ok {
          fmt.Printf("placed '%s' at RIGHT %d, %d\n", word, i, j)
          words.MarkUsed(word)
          // recurse
          if phaseTwo(words, g) {
            return true
          }
          g.Unplace(i, j, grid.RIGHT, word, a, z)
          fmt.Printf("unplaced '%s' at RIGHT %d, %d\n", word, i, j)
          words.MarkUnused(word)
        }
      }
    }
  }
  return false
}
