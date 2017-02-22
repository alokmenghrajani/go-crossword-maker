package words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	w := Load("../test_wordlist.txt")
	assert.Equal(t, len(w.words), 4)

	ngrams := w.ngrams["wo"]
	assert.Equal(t, 2, len(ngrams))
	assert.Contains(t, ngrams, Ngram{"wood", 0})
	assert.Contains(t, ngrams, Ngram{"world", 0})

	ngrams = w.ngrams["orl"]
	assert.Equal(t, 2, len(ngrams))
	assert.Contains(t, ngrams, Ngram{"neighborly", 6})
	assert.Contains(t, ngrams, Ngram{"world", 1})
}
