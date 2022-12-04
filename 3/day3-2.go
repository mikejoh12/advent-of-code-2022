package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findCommonR(s1, s2, s3 string) rune {
	for _, r := range s1 {
		if strings.Contains(s2, string(r)) && strings.Contains(s3, string(r)) {
			return r
		}
	}
	return '?'
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var score int
	var rucksacks []string

	for scanner.Scan() {
		t := scanner.Text()
		rucksacks = append(rucksacks, t)
	}

	for i := 0; i < len(rucksacks) - 2; i += 3 {
		common := findCommonR(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		if common >= 'A' && common <= 'Z' {
			score += int(common) + 27 - 'A'
		} else {
			score += int(common) - 'a' + 1
		}
	}

	fmt.Println(score)
}