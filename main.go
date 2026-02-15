package main

import (
	"fmt"
	"piscine/piscine"
)

func main() {
	fmt.Println(piscine.WeAreUnique("", ""))
	fmt.Println(piscine.WeAreUnique("foo", "boo"))
	fmt.Println(piscine.WeAreUnique("abc", "def"))
	fmt.Println(piscine.WeAreUnique("hello", "world"))   // 5
	fmt.Println(piscine.WeAreUnique("aaa", "aaa"))       // 0
	fmt.Println(piscine.WeAreUnique("abc", "abc"))       // 0
	fmt.Println(piscine.WeAreUnique("a", ""))            // 1
	fmt.Println(piscine.WeAreUnique("", "b"))            // 1
	fmt.Println(piscine.WeAreUnique("abcdef", "ghijkl")) // 12
}
