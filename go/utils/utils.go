package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(day int) string {
	fileName := fmt.Sprintf("../input/day%02d.txt", day)

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading input file: ", fileName)
	}

	return string(file)
}

func ReadInputLines(day int) []string {
	file := ReadInput(day)
	return strings.Split(file, "\n")
}
