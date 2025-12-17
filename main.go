package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	// 1. Initialize the dictionary
	obj := solutions.WordConstructor()

	// 2. Add words
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")

	// 3. Perform searches
	param_1 := obj.Search("a")   // false
	param_2 := obj.Search(".at") // false

	// 4. Add another word and resume search
	obj.AddWord("bat")
	param_3 := obj.Search("b.t")  // true
	param_4 := obj.Search("an.")  // true
	param_5 := obj.Search("a.d.") // false
	param_6 := obj.Search("b.")   // false
	param_7 := obj.Search("a.d")  // true
	param_8 := obj.Search(".")    // false

	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)
}
