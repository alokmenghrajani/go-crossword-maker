package words

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
  . "../utils"
)

/**
 * Store all the words in a map and track if they have been picked or not.
 * Also, build a map of every n-gram to word. Loading a large dictionnary
 * can take a few seconds, but we only incur the cost once.
 *
 * TODO: it might be more efficient to use two lists instead?
 * TODO: ability to favor specific words using some kind of score?
 * TODO: ability to indicate that a word should only be picked if another
 *       one was picked. Helps when creating crosswords with specific themes.
 */
type Words struct {
	words  map[string]bool
	ngrams map[string][]string
}

func Load(filename string) *Words {
	w := new(Words)
	w.words = make(map[string]bool)
	w.ngrams = make(map[string][]string)

	f, err := os.Open(filename)
	PanicIfNotNil(err)
	defer f.Close()
	r := bufio.NewReader(f)
	inHeader := true
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		PanicIfNotNil(err)
		if strings.HasPrefix(line, "--------") {
			inHeader = false
		} else if !inHeader {
			w.addWord(line)
		}
	}
	return w
}

func (words *Words) addWord(word string) {
	word = strings.TrimRight(word, "\n")
	words.words[word] = false

	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word); j++ {
			ngram := word[i:j]
			if l, ok := words.ngrams[ngram]; ok {
				words.ngrams[ngram] = append(l, word)
			} else {
				words.ngrams[ngram] = []string{word}
			}
		}
	}
}

func (words *Words) markUsed(word string) {
	PanicIfFalse(!words.words[word], fmt.Sprintf("expecting %s to be false", word))
	words.words[word] = true
}

func (words *Words) markUnused(word string) {
	PanicIfFalse(words.words[word], fmt.Sprintf("expecting %s to be true", word))
	words.words[word] = false
}