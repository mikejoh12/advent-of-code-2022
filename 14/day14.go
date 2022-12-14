package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)

type path []image.Point
type grid [][]rune

func newGrid(dx, dy int) *grid {
	g := make(grid, dy)
	for i := range g {
		g[i] = make([]rune, dx)
	}
	return &g
}

func (g *grid) drawFloor(maxY int) {
	floorY := maxY + 2
	for x := 0; x < len((*g)[0]); x++ {
		(*g)[floorY][x] = '#'
	}
}

func (g *grid) dropSand(maxY int) bool {
	down := image.Point{0, 1}
	downLeft := image.Point{-1, 1}
	downRight := image.Point{1, 1}
	sandCoord := image.Point{500, 0}

	for {
		if sandCoord.Y == maxY || (*g)[sandCoord.Y][sandCoord.X] == 'o' {
			return false
		}

		downCoord := sandCoord.Add(down)
		if (*g)[downCoord.Y][downCoord.X] == 0 {
			sandCoord = sandCoord.Add(down)
			continue
		}
		downLeftCoord := sandCoord.Add(downLeft)
		if (*g)[downLeftCoord.Y][downLeftCoord.X] == 0 {
			sandCoord = sandCoord.Add(downLeft)
			continue
		}
		downRightCoord := sandCoord.Add(downRight)
		if (*g)[downRightCoord.Y][downRightCoord.X] == 0 {
			sandCoord = sandCoord.Add(downRight)
			continue
		}

		(*g)[sandCoord.Y][sandCoord.X] = 'o'
		return true
	}
}

func (g *grid) drawPaths(paths []path) {
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			switch {
			case path[i].X == path[i+1].X: // vertical line
				minY := min(path[i].Y, path[i+1].Y)
				maxY := max(path[i].Y, path[i+1].Y)
				for y := minY; y <= maxY; y++ {
					(*g)[y][path[i].X] = '#'
				}
			case path[i].Y == path[i+1].Y: // horizontal line
				minX := min(path[i].X, path[i+1].X)
				maxX := max(path[i].X, path[i+1].X)
				for x := minX; x <= maxX; x++ {
					(*g)[path[i].Y][x] = '#'
				}
			}
		}
	}
}

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

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var paths []path
	var maxY int

	for scanner.Scan() {
		t := scanner.Text()
		points := strings.Split(t, " -> ")
		var onePath path
		for _, p := range points {
			var x, y int
			fmt.Sscanf(p, "%d,%d", &x, &y)
			onePath = append(onePath, image.Point{x, y})
			maxY = max(maxY, y)
		}
		paths = append(paths, onePath)
	}

	grid1, grid2 := newGrid(1000, 200), newGrid(1000, 200)
	grid1.drawPaths(paths)
	grid2.drawPaths(paths)
	grid2.drawFloor(maxY)

	sand1 := 0
	for grid1.dropSand(199) {
		sand1++
	}

	sand2 := 0
	for grid2.dropSand(199) {
		sand2++
	}
	fmt.Println("part 1", sand1)
	fmt.Println("part 2", sand2)
}