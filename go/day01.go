package main

import (
	"fmt"
	"github.com/shreyassanthu77/aoc/go/utils"
	"strings"
)

func main() {
	// 	input := strings.Split(`1abc2
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet`, "\n")
	input := strings.Split(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, "\n")
	input = utils.ReadInputLines(1)

	// part1(input)
	part2(input)
}

func getValue(input string) int {
	var n1 = -1
	var n2 = -1

	for _, char := range input {
		if char >= '0' && char <= '9' {
			if n1 == -1 {
				n1 = int(char - '0')
			} else {
				n2 = int(char - '0')
			}
		}
	}

	if n2 == -1 {
		n2 = n1
	}
	n := n1*10 + n2
	return n
}

func part1(input []string) {
	sum := 0
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		sum += getValue(line)
	}
	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		toReplace := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		replaceWith := []string{"o1e", "t2o", "t3e", "f4r", "f5e", "s6x", "s7n", "e8t", "n9e"}
		for i := 0; i < len(toReplace); i++ {
			line = strings.ReplaceAll(line, toReplace[i], replaceWith[i])
		}
		n := getValue(line)
		sum += n
	}
	fmt.Println(sum)
}
