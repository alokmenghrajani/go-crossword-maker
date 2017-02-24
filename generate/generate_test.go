package generate

import (
	"../grid"
	"../words"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhaseTwoValidateDown(t *testing.T) {
	g := grid.New(4)
	g.Place(0, 0, grid.RIGHT, "bar")
	g.Place(1, 1, grid.RIGHT, "rat")

	w := words.New()
	w.AddWord("area")

	r, valid := phaseTwoValidateDown(w, g)
	assert.False(t, valid)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, 1, r[grid.Partial{"ar", 1, 0}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 2, 0}])

	w.AddWord("bra")
	r, valid = phaseTwoValidateDown(w, g)
	assert.False(t, valid)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, 1, r[grid.Partial{"ar", 1, 0}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 2, 0}])

	w.AddWord("rag")
	w.AddWord("art")
	r, valid = phaseTwoValidateDown(w, g)
	assert.True(t, valid)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, 2, r[grid.Partial{"ar", 1, 0}])
	assert.Equal(t, 1, r[grid.Partial{"ra", 2, 0}])

	w.MarkUsed("rag")
	r, valid = phaseTwoValidateDown(w, g)
	assert.False(t, valid)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, 2, r[grid.Partial{"ar", 1, 0}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 2, 0}])
}

func TestPhaseTwoValidateRight(t *testing.T) {
	g := grid.New(4)
	g.Place(0, 0, grid.DOWN, "bar")
	g.Place(1, 1, grid.DOWN, "rat")

	w := words.New()
	w.AddWord("area")

	r, valid := phaseTwoValidateRight(w, g)
	assert.False(t, valid)
	assert.Equal(t, 2, len(r))
	assert.Equal(t, 1, r[grid.Partial{"ar", 0, 1}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 0, 2}])

	w.AddWord("bra")
	r, valid = phaseTwoValidateRight(w, g)
	assert.False(t, valid)
	assert.Equal(t, 2, len(r))
	assert.Equal(t, 1, r[grid.Partial{"ar", 0, 1}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 0, 2}])

	w.AddWord("rag")
	w.AddWord("art")
	r, valid = phaseTwoValidateRight(w, g)
	assert.True(t, valid)
	assert.Equal(t, 2, len(r))
	assert.Equal(t, 2, r[grid.Partial{"ar", 0, 1}])
	assert.Equal(t, 1, r[grid.Partial{"ra", 0, 2}])

	w.MarkUsed("rag")
	r, valid = phaseTwoValidateRight(w, g)
	assert.False(t, valid)
	assert.Equal(t, 2, len(r))
	assert.Equal(t, 2, r[grid.Partial{"ar", 0, 1}])
	assert.Equal(t, 0, r[grid.Partial{"ra", 0, 2}])
}
