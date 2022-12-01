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

	var elfCal int
	var elfCals []int

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			cal, _ := strconv.Atoi(t)
			elfCal += cal
		} else {
			elfCals = append(elfCals, elfCal)
			elfCal = 0
		}
	}

	elfCals = append(elfCals, elfCal)
	sort.Slice(elfCals, func(i, j int) bool {
		return elfCals[i] > elfCals[j]
	})
	
	fmt.Println(elfCals[0]+elfCals[1]+elfCals[2])
}
