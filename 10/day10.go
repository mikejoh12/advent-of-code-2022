package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isMeasureCycle(c int) bool {
	return (c - 20) % 40 == 0
}

func drawPixel(d []string, cycle, x int) {
	if cycle%40 >= x - 1 && cycle%40 <= x + 1 {
		d[cycle] = "#"
	} else {
		d[cycle] = "."
	}
}

func getDisplayString(d []string) (output string) {
	for y := 0; y < 6; y++ {
		output += strings.Join(d[y*40:y*40+40], "") + "\n"
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	x, cycle, sigStrSum := 1, 0, 0
	display := make([]string, 240)

	for scanner.Scan() {
		t := scanner.Text()
		var op string
		var k int
		fmt.Sscanf(t, "%s %d", &op, &k)

		drawPixel(display, cycle, x)
		switch op {
		case "noop":
			cycle += 1
			if isMeasureCycle(cycle) {
				sigStrSum += cycle * x
			}
		case "addx":
			cycle++
			if isMeasureCycle(cycle) {
				sigStrSum += cycle * x
			}
			drawPixel(display, cycle, x)
			cycle++
			if isMeasureCycle(cycle) {
				sigStrSum += cycle * x
			}
			x += k
		}
	}

	fmt.Println("part 1:", sigStrSum)
	fmt.Print("part 2:\n", getDisplayString(display))
}