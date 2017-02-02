package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	g := New(4)
	s := g.String()
	assert.Equal(t, s, " .  .  .  . \n .  .  .  . \n .  .  .  . \n .  .  .  . \n")
}

func TestPlace(t *testing.T) {
	g := New(4)
	assert.True(t, g.Place(1, 1, DOWN, "foo"))
	s := g.String()
	assert.Equal(t, s, " .  #  .  . \n .  f  .  . \n .  o  .  . \n .  o  .  . \n")

	assert.True(t, g.Place(0, 2, RIGHT, "dome"))
	s = g.String()
	assert.Equal(t, s, " .  #  .  . \n .  f  .  . \n d  o  m  e \n .  o  .  . \n")

	assert.True(t, g.Place(3, 1, DOWN, "be"))
	s = g.String()
	assert.Equal(t, s, " .  #  .  # \n .  f  .  b \n d  o  m  e \n .  o  .  # \n")

	assert.True(t, g.Place(2, 0, DOWN, "sam"))
	s = g.String()
	assert.Equal(t, s, " .  #  s  # \n . (f)(a)(b)\n d  o  m  e \n .  o  #  # \n")

	fmt.Println(g)
	fmt.Printf("here: %s", g.rightPartial)	

	assert.True(t, g.Place(1, 1, RIGHT, "fab"))
	s = g.String()
	assert.Equal(t, s, " .  #  s  # \n #  f  a  b \n d  o  m  e \n .  o  #  # \n")
}
