package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func newPairs(f string) (output [][]any) {
	chars, _ := os.ReadFile("input.txt")
	pairStrings := strings.Split(string(chars), "\n\n")
	for _, pairStr := range pairStrings {
		packetStrings := strings.Split(pairStr, "\n")
		var newPair []any
		for _, packetStr := range packetStrings {
			new := getSlice(packetStr)
			newPair = append(newPair, new)
		}
		output = append(output, newPair)
	}
	return
}

func getSubList(s string, idx int) (sublist string, endIdx int) {
	stack := make([]rune, 0)
	var startHeight int
	for i, r := range s{
		switch {
		case i == idx && r == '[':
			startHeight = len(stack)
		case r == ']' && len(stack) == startHeight:
			return s[idx:i+1], i
		case r == '[':
			stack = append(stack, '[')
		case r == ']':
			stack = stack[:len(stack)-1]
		}
	}
	return "error in getSubList", -1

}

func getSlice(p1 string) (output []any) {
	for i := 0; i < len(p1); i++ {
		switch {
		case unicode.IsDigit(rune(p1[i])):
			output = append(output, int(p1[i])-'0')
		case i != 0 && p1[i] == '[':
			sub, endIdx := getSubList(p1, i)
			i = endIdx
			newSlice := getSlice(sub)
			output = append(output, newSlice)
		}
	}
	return output
}

func main() {
	pairs := newPairs("input.txt")
	for pairIdx, pair := range pairs {
		fmt.Println(pairIdx, pair)
	}
}