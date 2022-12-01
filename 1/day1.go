package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elfCals = []int{0}

	for scanner.Scan() {
		t := scanner.Text()
		cal, _ := strconv.Atoi(t)
		if t != "" {
			elfCals[len(elfCals)-1] += cal
		} else {
			elfCals = append(elfCals, cal)
		}
	}

	sort.Slice(elfCals, func(i, j int) bool {
		return elfCals[i] > elfCals[j]
	})

	fmt.Println("part 1:", elfCals[0])
	fmt.Println("part 2:", elfCals[0]+elfCals[1]+elfCals[2])
}