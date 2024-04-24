package main

import (
	"fmt"

	. "github.com/shreyassanthu77/aoc/go/utils/parsers"
)

func main() {
	parser := Many(
		Char('a'),
	)

	res := parser.Run("aabc")
	fmt.Printf("Result: %+v\n", res)
}
