package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var score int

	for scanner.Scan() {
		t := scanner.Text()
		move := strings.Split(t, " ")

		switch {
		case move[0] == "A" && move[1] == "Y":
			score += 1 + 3
		case move[0] == "A" && move[1] == "Z":
			score += 6 + 2
		case move[0] == "B" && move[1] == "X":
			score += 1
		case move[0] == "B" && move[1] == "Z":
			score += 3 + 6
		case move[0] == "C" && move[1] == "X":
			score += 2
		case move[0] == "C" && move[1] == "Y":
			score += 3 + 3
		case move[0] == "A" && move[1] == "X":
			score += 3
		case move[0] == "B" && move[1] == "Y":
			score += 2 + 3
		case move[0] == "C" && move[1] == "Z":
			score += 1 + 6
		}
	}
	
	fmt.Println(score)
}