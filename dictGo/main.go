package main

import (
	"fmt"

	"github.com/deli-hub/dictGo/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{"first": "First word"}
	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// else {
	// 	fmt.Println(definition)
	// }

	dictionary := mydict.Dictionary{}

	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
