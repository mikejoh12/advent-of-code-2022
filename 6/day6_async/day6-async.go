package main

import (
	"fmt"
	"os"
)

type markerResult struct {
    value 	int
    length 	int
}

func isByteSliceAllDifferent(chars []byte) bool {
	charMap := make(map[byte]bool)
	for _, char := range chars {
		charMap[char] = true
	}
	return len(charMap) == len(chars)
}

func findMarker(chars *[]byte,  length int, c chan markerResult) {
	for i := length-1; i < len(*chars)-1; i++ {
		if isByteSliceAllDifferent((*chars)[i-length+1:i+1]) {
			c <- markerResult{value: i + 1, length: length}
			break
		}
	}
}

func main() {
	chars, _ := os.ReadFile("input.txt")
    c1 := make(chan markerResult)
	
	go findMarker(&chars, 4, c1)
	go findMarker(&chars, 14, c1)

	res1, res2 := <- c1, <- c1

	fmt.Printf("For length %d result is %d\n", res1.length, res1.value)
	fmt.Printf("For length %d result is %d\n", res2.length, res2.value)
}