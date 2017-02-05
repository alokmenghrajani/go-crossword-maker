package words

import (
	. "../utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
 * Store all the words in a list and track if they have been picked or not.
 * Also, build a map of every n-gram to word. Loading a large dictionnary
 * can take a few seconds, but we only incur the cost once.
 *
 * TODO: ability to favor specific words using some kind of score?
 * TODO: ability to indicate that a word should only be picked if another
 *       one was picked. Helps when creating crosswords with specific themes.
 */
type Words struct {
	words  []string
	used   map[string]bool
	ngrams map[string][]string
}

func Load(filename string) *Words {
	w := &Words{[]string{}, make(map[string]bool), make(map[string][]string)}

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
	words.words = append(words.words, word)
	words.used[word] = false

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

func (words *Words) GetWords() []string {
	return words.words
}

func (words *Words) IsUsed(word string) bool {
	return words.used[word]
}

func (words *Words) MarkUsed(word string) {
	PanicIfFalse(!words.used[word], fmt.Sprintf("expecting %s to be false", word))
	words.used[word] = true
}

func (words *Words) MarkUnused(word string) {
	PanicIfFalse(words.used[word], fmt.Sprintf("expecting %s to be true", word))
	words.used[word] = false
}
