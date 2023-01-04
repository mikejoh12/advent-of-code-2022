package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type instrType int

const (
	turn instrType = iota
	move
)

type instruction struct {
	name          instrType
	steps         int
	turnDirection rune
}

type minMaxPt struct {
	min int
	max int
}

type monkeyMap struct {
	pos          image.Point
	direction    int // 0-right, 1-down, 2-left, 3-up
	grid         [][]rune
	rowMinMax    []minMaxPt
	colMinMax    []minMaxPt
	instructions []instruction
	horPairs     map[image.Point]image.Point
	vertPairs    map[image.Point]image.Point
}

func newMinMaxPts(length int) []minMaxPt {
	m := make([]minMaxPt, length)
	for i := range m {
		m[i].min = math.MaxInt
	}
	return m
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func (m *monkeyMap) print() {
	for _, row := range m.grid {
		fmt.Println(string(row) + "\n")
	}
}

func (m *monkeyMap) move() {
	var offset = map[int]image.Point{
		0: image.Pt(1, 0),
		1: image.Pt(0, 1),
		2: image.Pt(-1, 0),
		3: image.Pt(0, -1),
	}
instrLoop:
	for _, instr := range m.instructions {
		switch instr.name {
		case move:
			for i := 0; i < instr.steps; i++ {
				var newPos image.Point
				switch {
				case m.direction == 0 && m.pos.X == m.rowMinMax[m.pos.Y].max:
					newPos = m.horPairs[m.pos]
				case m.direction == 1 && m.pos.Y == m.colMinMax[m.pos.X].max:
					newPos = m.vertPairs[m.pos]
				case m.direction == 2 && m.pos.X == m.rowMinMax[m.pos.Y].min:
					newPos = m.horPairs[m.pos]
				case m.direction == 3 && m.pos.Y == m.colMinMax[m.pos.X].min:
					newPos = m.vertPairs[m.pos]
				default:
					newPos = m.pos.Add(offset[m.direction])
				}
				if m.grid[newPos.Y][newPos.X] == '.' {
					m.pos = newPos
				} else {
					continue instrLoop
				}
			}
		case turn:
			switch {
			case instr.turnDirection == 'L':
				m.direction = mod(m.direction - 1, 4)
			case instr.turnDirection == 'R':
				m.direction = mod(m.direction + 1, 4)
			}
		}
	}
	fmt.Println("part 1", (m.pos.Y+1)*1000 + (m.pos.X+1)*4 + m.direction)

}

func newMonkeyMap(d string) *monkeyMap {
	chars, _ := os.ReadFile("input.txt")
	data := strings.Split(string(chars), "\n\n")

	m := monkeyMap{
		horPairs: make(map[image.Point]image.Point),
		vertPairs: make(map[image.Point]image.Point),
	}
	var maxX int

	rows := strings.Split(data[0], "\n")
	m.grid = make([][]rune, len(rows))
	for y, row := range rows {
		for x, r := range row {
			m.grid[y] = append(m.grid[y], r)
			if r == '.' {
				maxX = maxInt(maxX, x)
			}
		}
	}

	m.rowMinMax = newMinMaxPts(len(m.grid))
	m.colMinMax = newMinMaxPts(maxX + 1)
	for y, row := range m.grid {
		for x, r := range row {
			if r == '.' || r == '#' {
				m.rowMinMax[y].min = minInt(m.rowMinMax[y].min, x)
				m.rowMinMax[y].max = maxInt(m.rowMinMax[y].max, x)
				m.colMinMax[x].min = minInt(m.colMinMax[x].min, y)
				m.colMinMax[x].max = maxInt(m.colMinMax[x].max, y)
			}
		}
	}
	for y, rowData := range m.rowMinMax {
		m.horPairs[image.Pt(rowData.min, y)] = image.Pt(rowData.max, y)
		m.horPairs[image.Pt(rowData.max, y)] = image.Pt(rowData.min, y)
	}
	for x, colData := range m.colMinMax {
		m.vertPairs[image.Pt(x, colData.min)] = image.Pt(x, colData.max)
		m.vertPairs[image.Pt(x, colData.max)] = image.Pt(x, colData.min)
	}

	m.pos = image.Pt(m.rowMinMax[0].min, 0)
	m.direction = 0

	nrStr := ""
	for i := 0; i < len(data[1]); i++ {
		r := rune(data[1][i])
		if unicode.IsDigit(r) {
			nrStr += string(r)
		}
		if i == len(data[1])-1 {
			nr, _ := strconv.Atoi(nrStr)
			m.instructions = append(m.instructions, instruction{
				name:  move,
				steps: nr,
			})
		}
		if (r == 'R' || r == 'L') && len(nrStr) > 0 {
			nr, _ := strconv.Atoi(nrStr)
			nrStr = ""
			m.instructions = append(m.instructions, instruction{
				name:  move,
				steps: nr,
			})
			m.instructions = append(m.instructions, instruction{
				name:          turn,
				turnDirection: r,
			})
		}
	}
	return &m
}

func main() {
	m := newMonkeyMap("input.txt")
	m.move()
}