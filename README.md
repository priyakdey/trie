# Trie

Trie/Prefix Tree implementation in golang.

## Table of contents:

- [Overview](#overview)
- [Usage](#usage)
  - [Examples](#code-example)
- [Local Development](#local-development)
- [License](#license)

## Overview

Tries are a form of string-indexed look-up data structure, which is used to store
a dictionary list of words that can be searched on in a manner that allows for 
efficient generation of completion lists.  

A prefix trie is an ordered tree data structure used in the representation of a 
set of strings over a finite alphabet set, which allows efficient storage of 
words with common prefixes.


## Usage

To get a reference to the Trie, use the `trie.New()` method.
Operations allowed are:
- Insert(word)
- Contains(word)
- ContainsPrefix(prefix)
- WordsWithPrefix(prefix)
- Delete(word)

_Right now, you can insert English letters. Entering UTF-8 characters, might break_

### Code Example:

```go
package main

import (
  "github.com/priyakdey/trie"
)

func main() {
  trie := trie.New()                              // Create a new tree and returns the ref
  trie.Insert("cat")
  trie.Insert("catwoman")
  
  isPresent := trie.Contains("cat")               // returns true
  isPresent = trie.Contains("catwo")              // returns false
  isPresent = trie.Contains("catwoman")           // returns true

  isPrefixPresent := trie.ContainsPrefix("ca")    //returns true
  isPrefixPresent = trie.ContainsPrefix("caaaa")  //returns false

  recommendations := WordsWithPrefix("ca")        // returns [cat, catwoman]

  trie.Delete("cat")                              // deletes the word from the dictionary

  recommendations = WordsWithPrefix("ca")         // returns [catwoman]
}


```

**More examples can be found under [examples/](examples) folder.**

## Local Development

1. For running all test cases run `make test`
1. For checking coverage run `make cov`.

_`make cov` assumes python3.10 installed and serves the html by running `http.server` module.
If not installed, you can comment this and use some other server like [Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) or `npm serve -s`_



## License

This project(all files) is under MIT License. More details under [LICENSE](LICENSE).
