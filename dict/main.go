package main

import (
	"fmt"

	"github.com/deli-hub/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}
