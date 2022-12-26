package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type elf struct {
	position image.Point
	proposed image.Point
}

type elves struct {
	locations  []elf
	occupied   map[image.Point]bool
	proposed   map[image.Point]int
	uLeft      image.Point
	lRight     image.Point
	directions [][]image.Point
}

func (e *elves) firstHalfElves() {
elfLoop:
	for elfIdx, elf := range e.locations {
		// Scan all around - don't move if all empty
		isClearAround := true
		for _, dir := range e.directions {
			for _, minorDir := range dir {
				pos := elf.position.Add(minorDir)
				if _, ok := e.occupied[pos]; ok {
					isClearAround = false
				}
			}
		}
		if isClearAround {
			continue elfLoop
		}

	moveLoop:
		for dirIdx, dir := range e.directions {
			for _, minorDir := range dir {
				pos := elf.position.Add(minorDir)
				if _, ok := e.occupied[pos]; ok {
					continue moveLoop
				}
			}
			switch dirIdx {
			case 0:
				e.locations[elfIdx].proposed = elf.position.Add(dir[0])
				e.proposed[e.locations[elfIdx].proposed]++
			case 1:
				e.locations[elfIdx].proposed = elf.position.Add(dir[0])
				e.proposed[e.locations[elfIdx].proposed]++
			case 2:
				e.locations[elfIdx].proposed = elf.position.Add(dir[0])
				e.proposed[e.locations[elfIdx].proposed]++
			case 3:
				e.locations[elfIdx].proposed = elf.position.Add(dir[0])
				e.proposed[e.locations[elfIdx].proposed]++
			}
			continue elfLoop
		}
	}
}

func (e *elves) secondHalfElves() (moves bool) {
	for elfIdx, cElf := range e.locations {
		count, ok := e.proposed[cElf.proposed]
		if ok && count == 1 {
			delete(e.occupied, cElf.position)
			e.locations[elfIdx].position = cElf.proposed
			e.occupied[cElf.proposed] = true
			e.uLeft.Y = min(e.uLeft.Y, cElf.proposed.Y)
			e.uLeft.X = min(e.uLeft.X, cElf.proposed.X)
			e.lRight.Y = max(e.lRight.Y, cElf.proposed.Y)
			e.lRight.X = max(e.lRight.X, cElf.proposed.X)
			moves = true
		}
	}

	// clear all proposals
	for elfIdx, cElf := range e.locations {
		delete(e.proposed, cElf.proposed)
		e.locations[elfIdx].proposed = e.locations[elfIdx].position // set to current pos
	}
	e.directions = append(e.directions[1:], e.directions[0])
	return
}

func NewElves(cd []byte) *elves {
	rows := strings.Split(string(cd), "\n")
	newElves := elves{
		occupied: make(map[image.Point]bool),
		proposed: make(map[image.Point]int),
		directions: [][]image.Point{
			{{0, -1}, {1, -1}, {-1, -1}},
			{{0, 1}, {1, 1}, {-1, 1}},
			{{-1, 0}, {-1, -1}, {-1, 1}},
			{{1, 0}, {1, -1}, {1, 1}},
		},
	}
	for y, row := range rows {
		for x, r := range row {
			if r == '#' {
				newElf := elf{
					position: image.Pt(x, y),
				}
				newElves.locations = append(newElves.locations, newElf)
				newElves.occupied[image.Pt(x, y)] = true
				newElves.uLeft.Y = min(newElves.uLeft.Y, y)
				newElves.uLeft.X = min(newElves.uLeft.X, x)
				newElves.lRight.Y = max(newElves.lRight.Y, y)
				newElves.lRight.X = max(newElves.lRight.X, x)
			}
		}
	}
	return &newElves
}

// Print prints out the elf positions
func (e *elves) print() {
	for y := e.uLeft.Y; y <= e.lRight.Y; y++ {
		var row string
		for x := e.uLeft.X; x <= e.lRight.X; x++ {
			if _, ok := e.occupied[image.Pt(x, y)]; ok {
				row += "#"
			} else {
				row += " "
			}
		}
		fmt.Println(row)
	}
}

func (e *elves) gndTiles() int {
	return (e.lRight.X-e.uLeft.X+1)*(e.lRight.Y-e.uLeft.Y+1) - len(e.locations)
}

func main() {
	craterData, _ := os.ReadFile("input.txt")
	e := NewElves(craterData)

	moving := true
	round := 1
	for moving {
		e.firstHalfElves()
		moving = e.secondHalfElves()
		if !moving {
			fmt.Println("part 2:", round)
			break
		}
		if round == 10 {
			fmt.Println("part 1", e.gndTiles())
		}
		round++
	}
}
