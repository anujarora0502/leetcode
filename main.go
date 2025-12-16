package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {

	// 1. ["Trie", []]
	obj := solutions.TrieConstructor()
	fmt.Println("null // Constructor")

	// 2. ["insert", ["ab"]]
	obj.Insert("ab")
	fmt.Println("null // insert(\"ab\")")

	// 3. ["search", ["abc"]]
	fmt.Printf("%v // search(\"abc\")\n", obj.Search("abc"))

	// 4. ["search", ["ab"]]
	fmt.Printf("%v // search(\"ab\")\n", obj.Search("ab"))

	// 5. ["startsWith", ["abc"]]
	fmt.Printf("%v // startsWith(\"abc\")\n", obj.StartsWith("abc"))

	// 6. ["startsWith", ["ab"]]
	fmt.Printf("%v // startsWith(\"ab\")\n", obj.StartsWith("ab"))

	// 7. ["insert", ["ab"]] (Duplicate insert)
	obj.Insert("ab")
	fmt.Println("null // insert(\"ab\")")

	// 8. ["search", ["abc"]]
	fmt.Printf("%v // search(\"abc\")\n", obj.Search("abc"))

	// 9. ["startsWith", ["abc"]]
	fmt.Printf("%v // startsWith(\"abc\")\n", obj.StartsWith("abc"))

	// 10. ["insert", ["abc"]]
	obj.Insert("abc")
	fmt.Println("null // insert(\"abc\")")

	// 11. ["search", ["abc"]]
	fmt.Printf("%v // search(\"abc\")\n", obj.Search("abc"))

	// 12. ["startsWith", ["abc"]]
	fmt.Printf("%v // startsWith(\"abc\")\n", obj.StartsWith("abc"))
}
