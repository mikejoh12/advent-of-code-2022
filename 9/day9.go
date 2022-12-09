package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

func tailWillMove(h, t image.Point) bool {
	return math.Abs(float64(h.X - t.X)) > 1 || math.Abs(float64(h.Y - t.Y)) > 1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	head, tail := image.Point{0, 0}, image.Point{0, 0}
	var positions = map[image.Point]bool{tail: true}

	up := image.Point{0, -1}
	right := image.Point{1, 0}
	down := image.Point{0, 1}
	left := image.Point{-1, 0}

	for scanner.Scan() {
		var dir rune
		var steps int

		t := scanner.Text()
		fmt.Sscanf(t, "%c %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			var moveDir image.Point
			switch dir {
			case 'U':
				moveDir = up
			case 'R':
				moveDir = right
			case 'D':
				moveDir = down
			case 'L':
				moveDir = left
			}

			head = head.Add(moveDir)

			if tailWillMove(head, tail) {
				tail = tail.Add(moveDir)

				switch {
				case (dir == 'U' || dir == 'D') && head.X != tail.X:
					tail.X = head.X
				case (dir == 'L' || dir == 'R') && head.Y != head.X:
					tail.Y = head.Y
				}
			}
			positions[tail] = true
		}
	}
	fmt.Println("Positions:", len(positions))
}