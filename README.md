# go-crossword-maker
A crossword maker written in Go (also known as a crossword grid compiler, setter, creator,
generator or composer). Solves the constraint satisfaction problem related to building crosswords
from a list of words.

This code is currently work in progress. You can run `go test -v .` and that's about it.

# building and running tests
To run the tests
```
go test -v ./grid ./words ./generate .
```

To run the code
(TODO: godep piece)
```
go build
./go-crossword-maker --wordlist ukacd.txt --size 4
```

# motivation for writing this in Go
I need a crossword maker which will perform reasonably well with a small wordlist. I'm building
a crossword with a very specific theme. It's a hard problem to solve manually and I found a
bunch of academic papers on this topic but not much usable code.

So I figured I'll implement my own piece of code and see how far I can get.

I picked Go because I want the end-result to be a command line tool that can run across multiple
different platforms. Go makes cross-compiling easy.

I also figured that Go enables writing efficient parallel processing code and that it will be
easier to keep all the cores busy. The box I'm planning to run this has multiple CPUs with 10 cores each.

# other ideas
- how hard would it be to make this run on a GPU?
