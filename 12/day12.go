package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

func getStartPos(grid [][]string) coord {
	for y, row := range grid {
		for x, char := range row {
			if char == "S" {
				return coord{y: y, x: x}
			}
		}
	}
	return coord{-1, -1}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var grid [][]string

	for scanner.Scan() {
		t := scanner.Text()
		row := strings.Split(t, "")
		grid = append(grid, row)
	}

	var find func(pos coord, steps int, visited map[coord]bool, curRune rune)
	var directions = [4]coord{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	minSteps := 99999

	start := getStartPos(grid)

	find = func(pos coord, steps int, visited map[coord]bool, curRune rune) {
		visited[pos] = true
		for _, direction := range directions {
			newPos := coord{y: pos.y + direction.y, x: pos.x + direction.x}
			if _, ok := visited[newPos]; !ok &&
				newPos.y >= 0 && newPos.y < len(grid) && newPos.x >= 0 && newPos.x < len(grid[0]) {
				if (curRune == 'y' || curRune == 'z') && grid[newPos.y][newPos.x] == "E" {
					fmt.Println("Found path", visited)
					if steps+1 < minSteps {
						minSteps = steps + 1
						return
					}
				}

				if rune(grid[newPos.y][newPos.x][0]) <= curRune+1 {
					newVisited := make(map[coord]bool)
					for c, visited := range visited {
						newVisited[c] = visited
					}
					find(newPos, steps+1, newVisited, rune(grid[newPos.y][newPos.x][0]))
				}

			}
		}
	}

	find(start, 0, make(map[coord]bool), 'a'-1)

	fmt.Println(grid, "start", start)
	fmt.Println("minSteps", minSteps)
}
