package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
   obj := solutions.Constructor()
   obj.Push(0)
	 obj.Push(1)
	 obj.Push(0)
	 fmt.Println(obj.GetMin())
   obj.Pop()
   fmt.Println(obj.GetMin())
}
