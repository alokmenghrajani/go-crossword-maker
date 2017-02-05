package grid

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
	_, _, ok := g.Place(1, 1, DOWN, "foo")
	assert.True(t, ok)
	s := g.String()
	assert.Equal(t, s, " .  #  .  . \n .  f| .  . \n .  o| .  . \n .  o| .  . \n")

	_, _, ok = g.Place(0, 2, RIGHT, "dome")
	assert.True(t, ok)
	s = g.String()
	assert.Equal(t, s, " .  #  .  . \n .  f| .  . \n d- o+ m- e-\n .  o| .  . \n")

	_, _, ok = g.Place(3, 1, DOWN, "be")
	assert.True(t, ok)
	s = g.String()
	assert.Equal(t, s, " .  #  .  # \n .  f| .  b|\n d- o+ m- e+\n .  o| .  # \n")

	_, _, ok = g.Place(2, 0, DOWN, "sam")
	assert.True(t, ok)
	s = g.String()
	assert.Equal(t, s, " .  #  s| # \n .  f| a| b|\n d- o+ m+ e+\n .  o| #  # \n")

	_, _, ok = g.Place(1, 1, RIGHT, "fab")
	assert.True(t, ok)
	s = g.String()
	assert.Equal(t, s, " .  #  s| # \n #  f+ a+ b+\n d- o+ m+ e+\n .  o| #  # \n")

	fmt.Println(g)
}

func TestUnplace(t *testing.T) {
	g := New(4)
	a1, z1, _ := g.Place(1, 1, DOWN, "foo")
	a2, z2, _ := g.Place(2, 0, RIGHT, "go")
	a3, z3, _ := g.Place(0, 2, RIGHT, "dome")
	s := g.String()
	assert.Equal(t, s, " .  #  g- o-\n .  f| .  . \n d- o+ m- e-\n .  o| .  . \n")

	g.Unplace(0, 2, RIGHT, "dome", a3, z3)
	s = g.String()
	assert.Equal(t, s, " .  #  g- o-\n .  f| .  . \n .  o| .  . \n .  o| .  . \n")

	g.Unplace(2, 0, RIGHT, "go", a2, z2)
	s = g.String()
	assert.Equal(t, s, " .  #  .  . \n .  f| .  . \n .  o| .  . \n .  o| .  . \n")

	g.Unplace(1, 1, DOWN, "foo", a1, z1)
	s = g.String()
	assert.Equal(t, s, " .  .  .  . \n .  .  .  . \n .  .  .  . \n .  .  .  . \n")
}

func TestUnplace2(t *testing.T) {
	g := New(4)
	a1, z1, _ := g.Place(1, 0, RIGHT, "foo")
	g.Unplace(1, 0, RIGHT, "foo", a1, z1)
	s := g.String()
	assert.Equal(t, s, " .  .  .  . \n .  .  .  . \n .  .  .  . \n .  .  .  . \n")
}
