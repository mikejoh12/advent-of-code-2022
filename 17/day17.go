package main

import (
	"fmt"
	"image"
	"os"
)

type chamber struct {
	yStart         int
	highestRock    int
	rockPos        image.Point
	currentRockIdx int
	jetPattern     []byte
	jetPatternIdx  int
	space          [4 * 2022][7]rune
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c *chamber) dropRock() {
	var rocks = [][]image.Point{
		{image.Pt(0, 0), image.Pt(1, 0), image.Pt(2, 0), image.Pt(3, 0)},
		{image.Point{1, 0}, image.Point{0, 1}, image.Pt(1, 1), image.Pt(2, 1), image.Pt(1, 2)},
		{image.Pt(0, 0), image.Pt(1, 0), image.Pt(2, 0), image.Pt(2, 1), image.Pt(2, 2)},
		{image.Pt(0, 0), image.Pt(0, 1), image.Pt(0, 2), image.Pt(0, 3)},
		{image.Pt(0, 0), image.Pt(1, 0), image.Pt(0, 1), image.Pt(1, 1)},
	}

	rock := rocks[c.currentRockIdx%5]
	c.rockPos = image.Pt(2, c.yStart) // Start 2 from left edge
	isFalling, isPushed := true, true

fall_loop:
	for isFalling {
		if isPushed {
			isPushed = false
			dir := c.jetPattern[c.jetPatternIdx % len(c.jetPattern)]
			c.jetPatternIdx++
			var lateral image.Point
			if dir == '<' {
				lateral = image.Pt(-1, 0)
			} else {
				lateral = image.Pt(1, 0)
			}
			newPos := c.rockPos.Add(lateral)
			for _, rockPt := range rock {
				newPt := newPos.Add(rockPt)
				if newPt.Y >= 0 && (newPt.X < 0 || newPt.X > 6 || c.space[newPt.Y][newPt.X] == '#') {
					continue fall_loop
				}
			}
			c.rockPos = newPos
		} else {
			isPushed = true
			newPos := c.rockPos.Add(image.Pt(0, -1))
			for _, rockPt := range rock {
				newPt := newPos.Add(rockPt)
				if newPt.Y < 0 || c.space[newPt.Y][newPt.X] == '#' {
					for _, drawRockPt := range rock {
						drawPt := c.rockPos.Add(drawRockPt)
						c.space[drawPt.Y][drawPt.X] = '#'
						if drawPt.Y > c.highestRock {
							c.highestRock = drawPt.Y
						}
					}
					c.yStart = c.highestRock + 4
					c.currentRockIdx++
					return
				}
			}
			c.rockPos.Y--
		}
	}

}

func newChamber(pattern []byte) *chamber {
	return &chamber{
		yStart: 3, // Start position for rock nr 1. Y increments up.
		jetPattern: pattern,
		highestRock: -1, // Set initial rock floor position
	}
}

func main() {
	chars, _ := os.ReadFile("input1.txt")
	c := newChamber(chars)

	for r := 0; r < 2022; r++ {
		c.dropRock()
	}
	fmt.Println(c.highestRock+1)
}
