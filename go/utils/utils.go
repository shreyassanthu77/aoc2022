package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadInput(day int) string {
	fileName := fmt.Sprintf("day%02d/input.txt", day)

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading input file", fileName)
	}

	return string(file)
}
