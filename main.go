package main

import (
	"fmt"
	"piscine/piscine"
)

func main() {
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
	fmt.Print(piscine.FifthAndSkip("e 5£ @ 8* 7 =56 ;"))

}
