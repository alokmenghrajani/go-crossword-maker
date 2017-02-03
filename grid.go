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
	size int
	grid [][]cell
}

func New(size int) *Grid {
	g := new(Grid)
	g.size = size
	g.grid = make([][]cell, size)
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
		if t, ok := g.grid[x][y+i].(charCell); ok {
			t.isDown = true
			g.grid[x][y+i] = t
		} else {
			// You might think that we can get away without using partial words
			// (and having word insertion ordering solve the partial words issue)
			// but then we'll fail to explore the entire search space because of
			// "circular" partial words. I prefer to go through the words in a fixed
			// order and deal with these partial words.
			g.grid[x][y+i] = charCell{char: word[i], isDown: true, isRight: false}
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
		if t, ok := g.grid[x+i][y].(charCell); ok {
			t.isRight = true
			g.grid[x+i][y] = t
		} else {
			g.grid[x+i][y] = charCell{char: word[i], isDown: false, isRight: true}
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
