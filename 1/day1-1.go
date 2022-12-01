package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elfCal, maxCal int

	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			cal, _ := strconv.Atoi(t)
			elfCal += cal
		} else {
			if elfCal > maxCal {
				maxCal = elfCal
			}
			elfCal = 0
		}
	}
	if elfCal > maxCal {
		elfCal = maxCal
	}
	fmt.Println(maxCal)
}
