package grid

import (
	. "../utils"
	"fmt"
)

type dir int

const (
	DOWN dir = iota
	RIGHT
)

type cell interface {
	isCell()
	String() string
}

type blackCell struct{}

func (c blackCell) isCell() {}
func (c blackCell) String() string {
	return " # "
}

type emptyCell struct{}

func (c emptyCell) isCell() {}
func (c emptyCell) String() string {
	return " . "
}

type charCell struct {
	char    byte
	isDown  bool
	isRight bool
}

func (c charCell) isCell() {}
func (c charCell) String() string {
	r := fmt.Sprintf(" %c", c.char)
	if c.isDown && !c.isRight {
		r += "|"
	} else if !c.isDown && c.isRight {
		r += "-"
	} else if c.isDown && c.isRight {
		r += "+"
	} else {
		r = "WTF"
	}
	return r
}

type point struct {
	x, y int
}

type Grid struct {
	Size int
	grid [][]cell
}

func New(size int) *Grid {
	g := &Grid{size, make([][]cell, size)}
	for i := 0; i < size; i++ {
		g.grid[i] = make([]cell, size)
		for j := 0; j < size; j++ {
			g.grid[i][j] = emptyCell{}
		}
	}
	return g
}

func (g *Grid) isValid(x, y int) bool {
	if x < 0 || x >= g.Size {
		return false
	}
	if y < 0 || y >= g.Size {
		return false
	}
	return true
}

func (g *Grid) isEmptyOrBlack(x, y int) bool {
	if !g.isValid(x, y) {
		return true
	}
	if _, ok := g.grid[x][y].(blackCell); ok {
		return true
	}
	if _, ok := g.grid[x][y].(emptyCell); ok {
		return true
	}
	return false
}

func (g *Grid) isEmptyOrLetter(x, y int, char byte) bool {
	if !g.isValid(x, y) {
		return false
	}
	if _, ok := g.grid[x][y].(emptyCell); ok {
		return true
	}
	if c, ok := g.grid[x][y].(charCell); ok {
		if c.char == char {
			return true
		}
	}
	return false
}

/**
 * Attempt to place a word at a given position.
 * Returns information required to reverse the operation. It's cheaper to reverse
 * a Place than it is to clone the grid to allow back-tracking.
 */
func (g *Grid) Place(x, y int, dir dir, word string) (startBlack bool, endBlack bool, ok bool) {
	// we can place a word only if the cells before and after the word are empty or black
	// and if every cell is empty or has the same value as the letter being placed.
	if dir == DOWN {
		return g.placeDown(x, y, word)
	} else {
		return g.placeRight(x, y, word)
	}
}

func (g *Grid) placeDown(x, y int, word string) (startBlack bool, endBlack bool, ok bool) {
	startBlack = false
	endBlack = false
	ok = false
	if !g.isEmptyOrBlack(x, y-1) {
		return
	}
	if !g.isEmptyOrBlack(x, y+len(word)) {
		return
	}
	for i := 0; i < len(word); i++ {
		if !g.isEmptyOrLetter(x, y+i, word[i]) {
			return
		}
	}
	if g.isValid(x, y-1) {
		startBlack = true
		g.grid[x][y-1] = blackCell{}
	}
	if g.isValid(x, y+len(word)) {
		endBlack = true
		g.grid[x][y+len(word)] = blackCell{}
	}
	for i := 0; i < len(word); i++ {
		if t, ok := g.grid[x][y+i].(charCell); ok {
			PanicIfFalse(t.isRight, "expecting isRight to be true")
			t.isDown = true
			g.grid[x][y+i] = t
		} else {
			// You might think that we can get away without using partial words
			// (and having word insertion ordering solve the partial words issue)
			// but then we'll fail to explore the entire search space because of
			// "circular" partial words.
			g.grid[x][y+i] = charCell{char: word[i], isDown: true, isRight: false}
		}
	}
	ok = true
	return
}

func (g *Grid) placeRight(x, y int, word string) (startBlack bool, endBlack bool, ok bool) {
	startBlack = false
	endBlack = false
	ok = false
	if !g.isEmptyOrBlack(x-1, y) {
		return
	}
	if !g.isEmptyOrBlack(x+len(word), y) {
		return
	}
	for i := 0; i < len(word); i++ {
		if !g.isEmptyOrLetter(x+i, y, word[i]) {
			return
		}
	}
	if g.isValid(x-1, y) {
		startBlack = true
		g.grid[x-1][y] = blackCell{}
	}
	if g.isValid(x+len(word), y) {
		endBlack = true
		g.grid[x+len(word)][y] = blackCell{}
	}
	for i := 0; i < len(word); i++ {
		if t, ok := g.grid[x+i][y].(charCell); ok {
			PanicIfFalse(t.isDown, "expecting isDown to be true")
			t.isRight = true
			g.grid[x+i][y] = t
		} else {
			g.grid[x+i][y] = charCell{char: word[i], isDown: false, isRight: true}
		}
	}
	ok = true
	return
}

/**
 * Reverses a word placed by Place.
 * Note: the current implementation only works if the order is preserved. We
 *       could make things more reliable by tracking 4 directions for each
 *       blackCell, but it's probably not going to be useful.
 *       I could make things more fool-proof by pushing the startBlack/endBlack
 *       on a stack and having a parameter-less Undo function.
 */
func (g *Grid) Unplace(x, y int, dir dir, word string, startBlack bool, endBlack bool) {
	if dir == DOWN {
		g.unplaceDown(x, y, word, startBlack, endBlack)
	} else {
		g.unplaceRight(x, y, word, startBlack, endBlack)
	}
}

func (g *Grid) unplaceDown(x, y int, word string, startBlack bool, endBlack bool) {
	// Check that word was placed at x, y
	for i := 0; i < len(word); i++ {
		t, ok := g.grid[x][y+i].(charCell)
		PanicIfFalse(ok, fmt.Sprintf("expecting charCell at x, y+i; got: ", x, y, i, g.grid[x][y+i]))
		PanicIfFalse(t.char == word[i], fmt.Sprintf("expecting specific char: %c != %c", t.char, word[i]))
		PanicIfFalse(t.isDown, "expecting isDown to be set")
	}
	// Revert placeDown
	if startBlack {
		g.grid[x][y-1] = emptyCell{}
	}
	if endBlack {
		g.grid[x][y+len(word)] = emptyCell{}
	}
	for i := 0; i < len(word); i++ {
		t := g.grid[x][y+i].(charCell)
		if t.isRight {
			t.isDown = false
			g.grid[x][y+i] = t
		} else {
			g.grid[x][y+i] = emptyCell{}
		}
	}
}

func (g *Grid) unplaceRight(x, y int, word string, startBlack bool, endBlack bool) {
	// Check that word was placed at x, y
	for i := 0; i < len(word); i++ {
		t, ok := g.grid[x+i][y].(charCell)
		PanicIfFalse(ok, fmt.Sprintf("expecting charCell at %d+%d, %d; got: ", x, i, y, g.grid[x+i][y]))
		PanicIfFalse(t.char == word[i], fmt.Sprintf("expecting specific char: %c != %c", t.char, word[i]))
		PanicIfFalse(t.isRight, "expecting isRight to be set")
	}
	// Revert placeDown
	if startBlack {
		g.grid[x-1][y] = emptyCell{}
	}
	if endBlack {
		g.grid[x+len(word)][y] = emptyCell{}
	}
	for i := 0; i < len(word); i++ {
		t := g.grid[x+i][y].(charCell)
		if t.isDown {
			t.isRight = false
			g.grid[x+i][y] = t
		} else {
			g.grid[x+i][y] = emptyCell{}
		}
	}
}

func (g *Grid) PartialDown() []string {
	r := []string{}
	for i := 0; i < g.Size; i++ {
		partial := ""
		for j := 0; j < g.Size; j++ {
			t, ok := g.grid[i][j].(charCell)
			if ok && !t.isDown {
				PanicIfFalse(t.isRight, "either isDown or isRight should be set")
				partial += fmt.Sprintf("%c", t.char)
			} else {
				if len(partial) > 1 {
					r = append(r, partial)
				}
				partial = ""
			}
		}
		if len(partial) > 1 {
			r = append(r, partial)
		}
	}
	return r
}

func (g *Grid) PartialRight() []string {
	r := []string{}
	for j := 0; j < g.Size; j++ {
		partial := ""
		for i := 0; i < g.Size; i++ {
			t, ok := g.grid[i][j].(charCell)
			if ok && !t.isRight {
				PanicIfFalse(t.isDown, "either isDown or isRight should be set")
				partial += fmt.Sprintf("%c", t.char)
			} else {
				if len(partial) > 1 {
					r = append(r, partial)
				}
				partial = ""
			}
		}
		if len(partial) > 1 {
			r = append(r, partial)
		}
	}
	return r
}

func (g *Grid) String() string {
	s := ""
	for j := 0; j < g.Size; j++ {
		for i := 0; i < g.Size; i++ {
			s += g.grid[i][j].String()
		}
		s += "\n"
	}
	return s
}
