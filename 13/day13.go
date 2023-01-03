package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type order int
const (
	inOrder order = iota
	equal
	notInOrder
)	

func newPairs(f string) (output [][2][]any) {
	chars, _ := os.ReadFile(f)
	pairStrings := strings.Split(string(chars), "\n\n")
	for _, pairStr := range pairStrings {
		packetStrings := strings.Split(pairStr, "\n")
		var newPair [2][]any
		for i, packetStr := range packetStrings {
			var packet []any
			json.Unmarshal([]byte(packetStr), &packet)
			newPair[i] = packet
		}
		output = append(output, newPair)
	}
	return
}

func isInOrder(left, right any) order {
		switch left := left.(type) {
		case float64:
			switch right := right.(type) {
			case float64:
				if left < right {
					return inOrder
				} else if left == right {
					return equal
				}
				return notInOrder
			case []any:
				return isInOrder([]any{left}, right)
			}			

		case []any:
			switch right := right.(type) {
			case float64:
				return isInOrder(left, []any{right})
			case []any:
				if len(left) == 0 && len(right) > 0 {
					return inOrder
				}

				for i, leftEl := range left {
					if i >= len(right) {
						return notInOrder
					}

					orderResult := isInOrder(leftEl, right[i])
					switch orderResult {
					case inOrder, notInOrder:
						return orderResult
					case equal:
						if i == len(left)-1 && len(left) < len(right){
							return inOrder
						}
					}
				}
			}
		}
	return equal
}

func main() {
	p := newPairs("input.txt")
	sum := 0
	for i, pair := range p {
		if isInOrder(pair[0], pair[1]) == inOrder{
			sum += i + 1
		}
	}
	fmt.Println("part 1", sum)
}