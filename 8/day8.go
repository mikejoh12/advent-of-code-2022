package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(s []string) (result []int) {
	for _, nr := range s {
		intNr, _ := strconv.Atoi(nr)
		result = append(result, intNr)
	}
	return
}

func isDirVis(nrSlice []int, height int) bool {
	for _, nr := range nrSlice {
		if nr >= height {
			return false
		}
	}
	return true
}

func getDirections(yPos, xPos int, grid [][]int) [4][]int {
	var upSlice, downSlice, leftSlice []int
	for y := yPos - 1; y >= 0; y-- {
		upSlice = append(upSlice, grid[y][xPos])
	}
	for y := yPos + 1; y < len(grid); y++ {
		downSlice = append(downSlice, grid[y][xPos])
	}
	for x := xPos - 1; x >= 0; x-- {
		leftSlice = append(leftSlice, grid[yPos][x])
	}
	return [4][]int{upSlice, downSlice, leftSlice, grid[yPos][xPos+1:]}
}

func isVis(yPos, xPos, height int, directions [4][]int) bool {
	for _, direction := range directions {
		if isDirVis(direction, height) {
			return true
		}
	}
	return false
}

func scorePos(directions [4][]int, height int) int {
	score := 1
	for _, direction := range directions {
		viewCount := 0
		for _, nr := range direction {
			if nr >= height {
				viewCount++
				break
			}
			viewCount++
		}
		score *= viewCount
	}
	return score
}

func ReadFile(f string) [][]int {
	file, _ := os.Open(f)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	grid := make([][]int, 0)

	for scanner.Scan() {
		t := scanner.Text()
		heights := strToInt(strings.Split(t, ""))
		grid = append(grid, heights)
	}
	return grid
}

func Part1and2(grid [][]int) (int, int) {
	visible, highestScenic := 0, 0
	for y, row := range grid {
		for x := range row {
			directions := getDirections(y, x, grid)
			if isVis(y, x, grid[y][x], directions) {
				visible++
			}
			score := scorePos(directions, grid[y][x])
			if score > highestScenic {
				highestScenic = score
			}

		}
	}
	return visible, highestScenic
}

func main() {
	grid := ReadFile("input.txt")
	visible, highestScenic := Part1and2(grid)
	fmt.Println("part 1:", visible)
	fmt.Println("part 2", highestScenic)
}