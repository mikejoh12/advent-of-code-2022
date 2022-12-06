package main

import (
	"fmt"
	"os"
)

func isByteSliceAllDifferent(chars []byte) bool {
	charMap := make(map[byte]bool)
	for _, char := range chars {
		charMap[char] = true
	}
	return len(charMap) == len(chars)
}

func main() {
	chars, _ := os.ReadFile("input.txt")
	var marker1, marker2 int

	for i := 3; i < len(chars)-1; i++ {
		if isByteSliceAllDifferent(chars[i-3:i+1]) {
			marker1 = i+1
			break
		}
	}

	for i := 13; i < len(chars)-1; i++ {
		if isByteSliceAllDifferent(chars[i-13:i+1]) {
			marker2 = i+1
			break
		}
	}

	fmt.Println("part 1:", marker1)
	fmt.Println("part 2:", marker2)
}