package main

import (
	"fmt"

	wcounter "../word-counter"
)

func main() {
	text := "a hello hello there ther a b c d e f g g t r"
	res := wcounter.CalcWordFreq(text)
	fmt.Println(res)
}
