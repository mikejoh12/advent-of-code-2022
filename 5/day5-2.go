package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3
*/
/*
[N]     [Q]         [N]
[R]     [F] [Q]     [G] [M]
[J]     [Z] [T]     [R] [H] [J]
[T] [H] [G] [R]     [B] [N] [T]
[Z] [J] [J] [G] [F] [Z] [S] [M]
[B] [N] [N] [N] [Q] [W] [L] [Q] [S]
[D] [S] [R] [V] [T] [C] [C] [N] [G]
[F] [R] [C] [F] [L] [Q] [F] [D] [P]
 1   2   3   4   5   6   7   8   9
*/
var stack = [][]rune{
	{'F', 'D', 'B', 'Z', 'T', 'J', 'R', 'N'},
	{'R', 'S', 'N', 'J', 'H'},
	{'C', 'R', 'N', 'J', 'G', 'Z', 'F', 'Q'},
	{'F', 'V', 'N','G','R','T','Q'},
	{'L','T','Q','F'},
	{'Q','C','W','Z','B','R','G','N'},
	{'F', 'C','L','S','N','H','M'},
	{'D','N','Q','M','T','J'},
	{'P','G','S'},
}

func rev(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fmt.Println("start stack", stack)

	for scanner.Scan() {
		t := scanner.Text()
		var qtyFrom, stackFrom, stackTo int
		fmt.Sscanf(t, "move %d from %d to %d", &qtyFrom, &stackFrom, &stackTo)
		fmt.Println(qtyFrom, stackFrom, stackTo)

		stack[stackTo-1] = append(stack[stackTo-1], stack[stackFrom-1][len(stack[stackFrom-1])-qtyFrom:]...)
		stack[stackFrom-1] = stack[stackFrom-1][:len(stack[stackFrom-1])-qtyFrom]
		fmt.Println("updated stack", stack)
	}
	fmt.Println(stack)
	for _, s := range stack {
		fmt.Print(string(s[len(s)-1]))
	}
}