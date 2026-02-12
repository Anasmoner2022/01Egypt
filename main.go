package main

import (
	"fmt"
	"piscine/piscine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
	fmt.Print(piscine.FromTo(0, 5))    // 00, 01, 02, 03, 04, 05
	fmt.Print(piscine.FromTo(5, 0))    // 05, 04, 03, 02, 01, 00
	fmt.Print(piscine.FromTo(-1, 5))   // Invalid
	fmt.Print(piscine.FromTo(95, 100)) // Invalid
}
