package generate

import (
	"../grid"
	. "../utils"
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

// Checks that all the partial words can be solved.
// Returns a map of partial -> how many permutations are valid.
func phaseTwoValidateDown(w *words.Words, g *grid.Grid) (map[grid.Partial][]words.Ngram, bool) {
	r := make(map[grid.Partial][]words.Ngram)

	// For each partial word, count how many words we can still place
	for _, partial := range g.PartialDown() {
		r[partial] = []words.Ngram{}
		ngrams := w.GetNgrams(partial.Partial)
		for _, ngram := range ngrams {
			if w.IsUsed(ngram.Word) {
				continue
			}
			// try to place ngram.word at partial.x, partial.y - ngram.offset
			sb, eb, ok := g.Place(partial.X, partial.Y-ngram.Offset, grid.DOWN, ngram.Word)
			if ok {
				r[partial] = append(r[partial], ngram)
				g.Unplace(partial.X, partial.Y-ngram.Offset, grid.DOWN, ngram.Word, sb, eb)
			}
		}
		if len(r[partial]) == 0 {
			return r, false
		}
	}
	return r, true
}

func phaseTwoValidateRight(w *words.Words, g *grid.Grid) (map[grid.Partial][]words.Ngram, bool) {
	r := make(map[grid.Partial][]words.Ngram)

	// For each partial word, count how many words we can still place
	for _, partial := range g.PartialRight() {
		r[partial] = []words.Ngram{}
		ngrams := w.GetNgrams(partial.Partial)
		for _, ngram := range ngrams {
			if w.IsUsed(ngram.Word) {
				continue
			}
			// try to place ngram.word at partial.x - ngram.offset, partial.y
			sb, eb, ok := g.Place(partial.X-ngram.Offset, partial.Y, grid.RIGHT, ngram.Word)
			if ok {
				r[partial] = append(r[partial], ngram)
				g.Unplace(partial.X-ngram.Offset, partial.Y, grid.RIGHT, ngram.Word, sb, eb)
			}
		}
		if len(r[partial]) == 0 {
			return r, false
		}
	}
	return r, true
}

func phaseTwo(w *words.Words, g *grid.Grid, score int) bool {
	partialDown, valid := phaseTwoValidateDown(w, g)
	if !valid {
		return false
	}
	partialRight, valid := phaseTwoValidateRight(w, g)
	if !valid {
		return false
	}

	// Resolve the first partial with a count of 1.
	for partial, ngrams := range partialDown {
		if len(ngrams) == 1 {
			sb, eb, ok := g.Place(partial.X, partial.Y-ngrams[0].Offset, grid.DOWN, ngrams[0].Word)
			PanicIfFalse(ok, "partialDown contains invalid data")
			w.MarkUsed(ngrams[0].Word)
			valid := phaseTwo(w, g, score)
			g.Unplace(partial.X, partial.Y-ngrams[0].Offset, grid.DOWN, ngrams[0].Word, sb, eb)
			w.MarkUnused(ngrams[0].Word)
			return valid
		}
	}
	for partial, ngrams := range partialRight {
		if len(ngrams) == 1 {
			sb, eb, ok := g.Place(partial.X-ngrams[0].Offset, partial.Y, grid.RIGHT, ngrams[0].Word)
			PanicIfFalse(ok, "partialRight contains invalid data")
			w.MarkUsed(ngrams[0].Word)
			valid := phaseTwo(w, g, score)
			g.Unplace(partial.X-ngrams[0].Offset, partial.Y, grid.RIGHT, ngrams[0].Word, sb, eb)
			w.MarkUnused(ngrams[0].Word)
			return valid
		}
	}

	// Iterate over first partial
	for partial, ngrams := range partialDown {
		valid := false
		for _, ngram := range ngrams {
			sb, eb, ok := g.Place(partial.X, partial.Y-ngram.Offset, grid.DOWN, ngram.Word)
			PanicIfFalse(ok, "partialDown contains invalid data")
			w.MarkUsed(ngram.Word)
			valid = valid || phaseTwo(w, g, score)
			g.Unplace(partial.X, partial.Y-ngram.Offset, grid.DOWN, ngram.Word, sb, eb)
			w.MarkUnused(ngram.Word)
		}
		return valid
	}
	for partial, ngrams := range partialRight {
		valid := false
		for _, ngram := range ngrams {
			sb, eb, ok := g.Place(partial.X-ngram.Offset, partial.Y, grid.RIGHT, ngram.Word)
			PanicIfFalse(ok, "partialRight contains invalid data")
			w.MarkUsed(ngram.Word)
			valid = valid || phaseTwo(w, g, score)
			g.Unplace(partial.X-ngram.Offset, partial.Y, grid.RIGHT, ngram.Word, sb, eb)
			w.MarkUnused(ngram.Word)
		}
		return valid
	}

	// We don't have any partials, grid is in good state!
	fmt.Printf("score: %d\n", score)
	fmt.Println(g)

	return phaseOne(w, g, score)
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
