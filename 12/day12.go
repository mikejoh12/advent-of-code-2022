package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func getStartPos(grid [][]rune) image.Point {
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				return image.Point{x, y}
			}
		}
	}
	return image.Point{-1, -1}
}

func getAdjList(grid [][]rune) map[image.Point][]image.Point {
	var directions = [4]image.Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	adjList := make(map[image.Point][]image.Point)

	for y, row := range grid {
		for x, r := range row {
			adjList[image.Point{x, y}] = []image.Point{}
			for _, direction := range directions {
				newPos := image.Point{x, y}.Add(direction)
				if newPos.Y >= 0 && newPos.Y < len(grid) && newPos.X >= 0 && newPos.X < len(grid[0]) {

					// Start pos
					if r == 'S' && grid[newPos.Y][newPos.X] == 'a' {
						adjList[image.Point{x, y}] = append(adjList[image.Point{x, y}], newPos)
					}
					// Max one step up
					if grid[newPos.Y][newPos.X] != 'S' && grid[newPos.Y][newPos.X] != 'E' && grid[newPos.Y][newPos.X] <= r+1 {
						adjList[image.Point{x, y}] = append(adjList[image.Point{x, y}], newPos)
					}
					// End pos
					if (r == 'y' || r == 'z') && grid[newPos.Y][newPos.X] == 'E' {
						adjList[image.Point{x, y}] = append(adjList[image.Point{x, y}], newPos)
					}
				}
			}
		}
	}
	return adjList
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var grid [][]rune

	for scanner.Scan() {
		rowStr := scanner.Text()
		rowRunes := make([]rune, 0)
		for _, r := range rowStr {
			rowRunes = append(rowRunes, r)
		}
		grid = append(grid, rowRunes)
	}

	var find func(pos image.Point, steps int, visited map[image.Point]bool, curRune rune)
	minSteps := 99999

	start := getStartPos(grid)
	adjList := getAdjList(grid)
	fmt.Println(adjList)
	fmt.Println("start", start)

	find = func(pos image.Point, steps int, visited map[image.Point]bool, curRune rune) {
		visited[pos] = true

		for _, newPos := range adjList[pos] {
			if _, ok := visited[newPos]; !ok {
				if grid[newPos.Y][newPos.X] == 'E' {
					fmt.Println("Found path - steps", steps+1)
					if steps+1 < minSteps {
						minSteps = steps + 1
					}
					return
				}

				newVisited := make(map[image.Point]bool)
				for c, visited := range visited {
					newVisited[c] = visited
				}
				find(newPos, steps+1, newVisited, grid[newPos.Y][newPos.X])
			}
		}

	}

	find(start, 0, make(map[image.Point]bool), 'S')

	fmt.Println("minSteps", minSteps)
}
