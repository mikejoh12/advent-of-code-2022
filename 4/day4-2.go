package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func main() {

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println(t)
		sides := strings.Split(t, ",")
		left := strings.Split(sides[0], "-")
		right := strings.Split(sides[1], "-")
		l1 := toInt(left[0])
		l2 := toInt(left[1])
		r1 := toInt(right[0])
		r2 := toInt(right[1])

		if (l1 >= r1 && l2 <= r2) || (r1 >= l1 && r2 <= l2) {
			count++
		} else if (l1 >= r1 && l1 <= r2) || (l2 >= r1 && l2 <= r2) {
			count++
		}

	}
	fmt.Println(count)
}