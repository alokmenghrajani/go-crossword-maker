package main

import (
	"fmt"
  "os"
  "bufio"
  "io"
  "strings"
)

type Words struct {
  words map[string]bool
  ngrams map[string][]string
}

func Load(filename string) *Words {
  w := new(Words)
  w.words = make(map[string]bool)
  w.ngrams = make(map[string][]string)

  f, err := os.Open(filename)
	panicIfNotNil(err)
	defer f.Close()
	r := bufio.NewReader(f)
  inHeader := true
  for {
    line, err := r.ReadString('\n')
    if err == io.EOF {
      break;
    }
    panicIfNotNil(err)
    if strings.HasPrefix(line, "--------") {
      inHeader = false;
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
  panicIfFalse(!words.words[word], fmt.Sprintf("expecting %s to be false", word))
  words.words[word] = true
}

func (words *Words) markUnused(word string) {
  panicIfFalse(words.words[word], fmt.Sprintf("expecting %s to be true", word))
  words.words[word] = false
}
