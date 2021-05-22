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

	/* ADD
	word := "hello"
	definition := "First"
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
	*/

	/* UPDATE
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Print(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	*/

	/* DELETE
	 */
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
