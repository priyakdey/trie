package main

import (
	"fmt"

	"github.com/priyakdey/trie"
)

func main() {
	t := trie.New()
	t.Insert("cap")
	t.Insert("captain")

	fmt.Println("Before deleting the word `cap`: ", t.WordsWithPrefix("cap"))

	t.Delete("cap")

	fmt.Println("After deleting the word `cap`: ", t.WordsWithPrefix("cap"))
	fmt.Println(t.WordsWithPrefix("cap"))

}
