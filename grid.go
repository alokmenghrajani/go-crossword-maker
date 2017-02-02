package main

import (
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
	char      byte
	isPartial bool
}

func (c charCell) isCell() {}
func (c charCell) String() string {
	if c.isPartial {
		return fmt.Sprintf("(%c)", c.char)
	} else {
		return fmt.Sprintf(" %c ", c.char)
	}
}

type point struct {
	x, y int
}

type Grid struct {
	size int
	grid [][]cell
	rightPartial map[point]int
	downPartial map[point]int
}

func New(size int) *Grid {
	g := new(Grid)
	g.size = size
	g.grid = make([][]cell, size)
	g.rightPartial = make(map[point]int)
	g.downPartial = make(map[point]int)
	for i := 0; i < size; i++ {
		g.grid[i] = make([]cell, size)
		for j := 0; j < size; j++ {
			g.grid[i][j] = emptyCell{}
		}
	}
	return g
}

func (g *Grid) isValid(x, y int) bool {
	if x < 0 || x >= g.size {
		return false
	}
	if y < 0 || y >= g.size {
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

func (g *Grid) removeRightPartial(x, y int) bool {
	_, ok := g.rightPartial[point{x, y}]
	if ok {
		delete(g.rightPartial, point{x, y})
		return true
	}
	return false
}

func (g *Grid) Place(x, y int, dir dir, word string) bool {
	// we can place a word only if the cells before and after the word are empty or black
	// and if every cell is empty or has the same value as the letter being placed.
	if dir == DOWN {
		return g.placeDown(x, y, word)
	} else {
		return g.placeRight(x, y, word)
	}
}

func (g *Grid) placeDown(x, y int, word string) bool {
	if !g.isEmptyOrBlack(x, y-1) {
		return false
	}
	if !g.isEmptyOrBlack(x, y+len(word)) {
		return false
	}
	for i := 0; i < len(word); i++ {
		if !g.isEmptyOrLetter(x, y+i, word[i]) {
			return false
		}
	}
	if g.isValid(x, y-1) {
		g.grid[x][y-1] = blackCell{}
	}
	if g.isValid(x, y+len(word)) {
		g.grid[x][y+len(word)] = blackCell{}
	}
	for i := 0; i < len(word); i++ {
		if _, ok := g.grid[x][y+i].(charCell); ok {
			g.grid[x][y+i] = charCell{char: word[i], isPartial: false}
		} else if g.isEmptyOrBlack(x-1, y+i) && g.isEmptyOrBlack(x+1, y+i) {
			g.grid[x][y+i] = charCell{char: word[i], isPartial: false}
		} else {
			// Since words are always surrounded by black cells, we only need
			// to look at the cell to the right and left and mark it as partial.
			// The whole partial tracking logic is currently very complicated and
			// hard to correctly test. Perhaps I should look into a different way
			// to track the state of the board. For each cell, I could store
			// if it's part of a down or right word. Then scan the board and find
			// partials that need to be fulfilled.
			//
			// Go all the way left. Look if it matches a partial. If it does, remove it
			// Go one right. Look if it matches a partial. If it does, remove it.
			// Add all the way left to all the way right as a new partial.
			leftOldPartial := 0
			if !g.isEmptyOrBlack(x-1, y+i) {
				t := g.grid[x-1][y+i].(charCell)
				t.isPartial = true
				g.grid[x-1][y+i] = t
				for j:=1; !g.isEmptyOrBlack(x-j, y+i); j++ {
					leftOldPartial = j
				}
			}
			if leftOldPartial > 1 {
				for j:=1; j<leftOldPartial; j++ {
					if g.removeRightPartial(x-j, y+i) {
						panicIfFalse(false, fmt.Sprintf("wasn't expecting a partial at %d-%d,%d+%d", x, j, y, i));
					}
				}
				if !g.removeRightPartial(x-leftOldPartial, y+i) {
					panicIfFalse(false, fmt.Sprintf("was expecting a partial at %d-%d,%d+%d", x, leftOldPartial, y, i));
				}
			}
			rightOldPartial := 0
			if !g.isEmptyOrBlack(x+1, y+i) {
				t := g.grid[x+1][y+i].(charCell)
				t.isPartial = true
				g.grid[x+1][y+i] = t
				for j:=1; !g.isEmptyOrBlack(x+j, y+i); j++ {
					rightOldPartial = j
				}
			}
			if rightOldPartial > 1 {
				for j:=1; j<rightOldPartial; j++ {
					if g.removeRightPartial(x+j, y+i) {
						panicIfFalse(false, fmt.Sprintf("wasn't expecting a partial at %d+%d,%d+%d", x, j, y, i));
					}
				}
				if !g.removeRightPartial(x+rightOldPartial, y+i) {
					panicIfFalse(false, fmt.Sprintf("was expecting a partial at %d+%d,%d+%d", x, rightOldPartial, y, i));
				}
			}
			g.rightPartial[point{x-leftOldPartial, y+i}] = leftOldPartial + rightOldPartial + 1

			// You might think that we can get away without using partial words
			// (and having word insertion ordering solve the partial words issue)
			// but then we'll fail to explore the entire search space because of
			// "circular" partial words. I prefer to go through the words in a fixed
			// order and deal with these partial words.
			g.grid[x][y+i] = charCell{char: word[i], isPartial: true}
		}
	}
	return true
}

func (g *Grid) placeRight(x, y int, word string) bool {
	if !g.isEmptyOrBlack(x-1, y) {
		return false
	}
	if !g.isEmptyOrBlack(x+len(word), y) {
		return false
	}
	for i := 0; i < len(word); i++ {
		if !g.isEmptyOrLetter(x+i, y, word[i]) {
			return false
		}
	}
	if g.isValid(x-1, y) {
		g.grid[x-1][y] = blackCell{}
	}
	if g.isValid(x+len(word), y) {
		g.grid[x+len(word)][y] = blackCell{}
	}
	for i := 0; i < len(word); i++ {
		if _, ok := g.grid[x+i][y].(charCell); ok {
			g.grid[x+i][y] = charCell{char: word[i], isPartial: false}
		} else if g.isEmptyOrBlack(x, y-1) && g.isEmptyOrBlack(x, y+1) {
			g.grid[x+i][y] = charCell{char: word[i], isPartial: false}
		} else {
			g.grid[x+i][y] = charCell{char: word[i], isPartial: true}
		}
	}
	return true
}

func (g *Grid) String() string {
	s := ""
	for j := 0; j < g.size; j++ {
		for i := 0; i < g.size; i++ {
			s += g.grid[i][j].String()
		}
		s += "\n"
	}
	return s
}
