package generate

import (
	//"fmt"
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

	r := phaseTwoValidateDown(w, g)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, r[grid.Partial{"ar", 1, 0}], 1)
	assert.Equal(t, r[grid.Partial{"ra", 2, 0}], 0)

	w.AddWord("bra")
	r = phaseTwoValidateDown(w, g)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, r[grid.Partial{"ar", 1, 0}], 1)
	assert.Equal(t, r[grid.Partial{"ra", 2, 0}], 0)

	w.AddWord("rag")
	w.AddWord("art")
	r = phaseTwoValidateDown(w, g)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, r[grid.Partial{"ar", 1, 0}], 2)
	assert.Equal(t, r[grid.Partial{"ra", 2, 0}], 1)

	w.MarkUsed("rag")
	r = phaseTwoValidateDown(w, g)
	assert.Equal(t, len(r), 2)
	assert.Equal(t, r[grid.Partial{"ar", 1, 0}], 2)
	assert.Equal(t, r[grid.Partial{"ra", 2, 0}], 0)
}
