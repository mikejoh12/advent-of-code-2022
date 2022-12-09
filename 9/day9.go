package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

func knotWillMove(h, t image.Point) bool {
	return math.Abs(float64(h.X-t.X)) > 1 || math.Abs(float64(h.Y-t.Y)) > 1
}

func moveKnot(head, tail image.Point) image.Point {
	switch {
	case math.Abs(float64(tail.X-head.X)) == 2 && math.Abs(float64(tail.Y-head.Y)) == 2: // 2 steps diagonal
		return image.Point{(head.X+tail.X)/2, (head.Y+tail.Y)/2}
	case tail.X == head.X: // up, down
		return image.Point{tail.X, (tail.Y + head.Y)/2}
	case tail.Y == head.Y: // right, left
		return image.Point{(tail.X+head.X)/2, tail.Y}
	case math.Abs(float64(tail.Y-head.Y)) == 2: // 2 steps up/down and diagonal
		return image.Point{head.X, (head.Y+tail.Y)/2}
	case math.Abs(float64(tail.X-head.X)) == 2: // 2 steps right/left and diagonal
		return image.Point{(head.X+tail.X)/2, head.Y}
	}
	return tail
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	knots := make([]image.Point, 10)
	var positions1 = map[image.Point]bool{knots[1]: true}
	var positions2 = map[image.Point]bool{knots[9]: true}

	for scanner.Scan() {
		var dir rune
		var steps int

		t := scanner.Text()
		fmt.Sscanf(t, "%c %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			switch dir {
			case 'U':
				knots[0] = image.Point{knots[0].X, knots[0].Y - 1}
			case 'R':
				knots[0] = image.Point{knots[0].X + 1, knots[0].Y}
			case 'D':
				knots[0] = image.Point{knots[0].X, knots[0].Y + 1}
			case 'L':
				knots[0] = image.Point{knots[0].X - 1, knots[0].Y}
			}

			for i := 1; i < len(knots); i++ {
				if knotWillMove(knots[i-1], knots[i]) {
					knots[i] = moveKnot(knots[i-1], knots[i])
				}
			}

			positions1[knots[1]] = true
			positions2[knots[9]] = true
		}
	}
	fmt.Println("part 1", len(positions1))
	fmt.Println("part 2:", len(positions2))
}