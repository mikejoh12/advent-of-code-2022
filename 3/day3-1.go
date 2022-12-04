package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findCommonR(s1, s2 string) rune {
	for _, r := range s1 {
		if strings.Contains(s2, string(r)) {
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

	for scanner.Scan() {
		t := scanner.Text()
		c1, c2 := t[:len(t)/2], t[len(t)/2:]
		common := findCommonR(c1, c2)

		if common >= 'A' && common <= 'Z' {
			score += int(common) + 27 - 'A'
		} else {
			score += int(common) - 'a' + 1
		}
	}
	fmt.Println(score)
}