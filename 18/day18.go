package main

import (
	"bufio"
	"fmt"
	"os"
)
type coord struct {
	x	int
	y	int
	z	int
}
type space [25][25][25]bool

func (s *space) getPtExposed(c coord) (count int) {
	var adjacent = [6]coord{
		{0, 0, 1}, // above
		{0, 0, -1}, // below
		{1, 0, 0}, // right
		{-1, 0, 0}, // left
		{0, 1, 0}, // behind
		{0, -1, 0}, // before
	}
	for _, ptOffset := range adjacent {
		pt := coord{c.x+ptOffset.x, c.y+ptOffset.y, c.z+ptOffset.z}
		if pt.x < 0 || pt.x > 24 || pt.y < 0 || pt.y > 24 || pt.z < 0 || pt.z > 24 {
			count++
			continue
		}
		if !s[pt.x][pt.y][pt.z] {
			count++
		}
	}
	return
}

func (s *space) getSfcArea() (total int) {
	for x := range s {
		for y := range s[0] {
			for z := range s[0][0] {
				if s[x][y][z] {
					ptArea := s.getPtExposed(coord{x,y,z})
					total += ptArea
				}
			}
		}
	}
	return
}

func NewSpace(coords []coord) *space {
	var s space
	for _, c := range coords {
		s[c.x][c.y][c.z] = true
	}
	return &s
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var coords []coord
	for scanner.Scan() {
		var x, y, z int
		t := scanner.Text()
		fmt.Sscanf(t, "%d,%d,%d", &x, &y, &z)
		coords = append(coords, coord{x,y,z})
	}

	s := NewSpace(coords)
	fmt.Println("part 1", s.getSfcArea())
}