package grid

import (
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

func TestUnplace3(t *testing.T) {
	g := New(4)
	g.Place(0, 1, DOWN, "foo")
	a2, z2, _ := g.Place(1, 0, RIGHT, "bar")
	g.Unplace(1, 0, RIGHT, "bar", a2, z2)
	s := g.String()
	assert.Equal(t, s, " #  .  .  . \n f| .  .  . \n o| .  .  . \n o| .  .  . \n")
}

func TestPartialDown(t *testing.T) {
	g := New(8)
	g.Place(0, 0, RIGHT, "foo")
	g.Place(1, 1, RIGHT, "bar")
	g.Place(0, 2, RIGHT, "aa")

	g.Place(1, 3, RIGHT, "hello")
	g.Place(2, 4, RIGHT, "dodo")
	g.Place(5, 2, DOWN, "moo")

	g.Place(3, 6, RIGHT, "toto")
	g.Place(3, 7, RIGHT, "world")

	partials := g.PartialDown()
	assert.Contains(t, partials, Partial{"obah", 1, 0})
	assert.Contains(t, partials, Partial{"oa", 2, 0})
	assert.Contains(t, partials, Partial{"ed", 2, 3})
	assert.Contains(t, partials, Partial{"lo", 3, 3})
	assert.Contains(t, partials, Partial{"tw", 3, 6})
	assert.Contains(t, partials, Partial{"ld", 4, 3})
	assert.Contains(t, partials, Partial{"oo", 4, 6})
	assert.Contains(t, partials, Partial{"tr", 5, 6})
	assert.Contains(t, partials, Partial{"ol", 6, 6})
	assert.Equal(t, 9, len(partials))
}

func TestPartialRight(t *testing.T) {
	g := New(8)
	g.Place(0, 0, DOWN, "foo")
	g.Place(1, 1, DOWN, "bar")
	g.Place(2, 0, DOWN, "aa")

	g.Place(3, 1, DOWN, "hello")
	g.Place(4, 2, DOWN, "dodo")
	g.Place(2, 5, RIGHT, "moo")

	g.Place(6, 3, DOWN, "toto")
	g.Place(7, 3, DOWN, "world")

	partials := g.PartialRight()
	assert.Contains(t, partials, Partial{"obah", 0, 1})
	assert.Contains(t, partials, Partial{"oa", 0, 2})
	assert.Contains(t, partials, Partial{"ed", 3, 2})
	assert.Contains(t, partials, Partial{"lo", 3, 3})
	assert.Contains(t, partials, Partial{"tw", 6, 3})
	assert.Contains(t, partials, Partial{"ld", 3, 4})
	assert.Contains(t, partials, Partial{"oo", 6, 4})
	assert.Contains(t, partials, Partial{"tr", 6, 5})
	assert.Contains(t, partials, Partial{"ol", 6, 6})
	assert.Equal(t, 9, len(partials))
}
