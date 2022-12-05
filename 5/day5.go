package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type move struct {
	qty		int
	from	int
	to		int
}

func rev(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getTopLetters(stacks [][]rune) (result string) {
	for _, stack := range stacks {
		if len(stack) > 0 {
			result += string(stack[len(stack)-1])
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	stacks1 := make([][]rune,9)
	stacks2 := make([][]rune,9)
	moves := make([]move, 0)

	for scanner.Scan() {
		t := scanner.Text()
		switch {
		case strings.Contains(t, "["):
			var stackIdx int
			for i := 1; i < len(t)-1; i += 4 {
				if unicode.IsLetter(rune(t[i])) {
					stacks1[stackIdx] = append([]rune{rune(t[i])}, stacks1[stackIdx]...)
					stacks2[stackIdx] = append([]rune{rune(t[i])}, stacks2[stackIdx]...)
				}
				stackIdx++
			}
		case strings.Contains(t, "move"):
			var qtyFrom, stackFrom, stackTo int
			fmt.Sscanf(t, "move %d from %d to %d", &qtyFrom, &stackFrom, &stackTo)
			moves = append(moves, move{qty: qtyFrom, from: stackFrom-1, to: stackTo-1})
		}
	}

	for _, m := range moves {
		stacks1[m.to] = append(stacks1[m.to], rev(stacks1[m.from][len(stacks1[m.from])-m.qty:])...)
		stacks1[m.from] = stacks1[m.from][:len(stacks1[m.from])-m.qty]

		stacks2[m.to] = append(stacks2[m.to], stacks2[m.from][len(stacks2[m.from])-m.qty:]...)
		stacks2[m.from] = stacks2[m.from][:len(stacks2[m.from])-m.qty]
	}

	fmt.Println("part 1:", getTopLetters(stacks1))
	fmt.Println("part 2:", getTopLetters(stacks2))
}