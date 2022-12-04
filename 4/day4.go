package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count1, count2 int

	for scanner.Scan() {
		t := scanner.Text()
		var l1, l2, r1, r2 int
		fmt.Sscanf(t, "%d-%d,%d-%d", &l1, &l2, &r1, &r2)

		switch {
		case (l1 >= r1 && l2 <= r2) || (r1 >= l1 && r2 <= l2):
			count1, count2 = count1+1, count2+1
		case (l1 >= r1 && l1 <= r2) || (l2 >= r1 && l2 <= r2):
			count2++
		}
	}
	fmt.Println("part 1:", count1)
	fmt.Println("part 2:", count2)
}
