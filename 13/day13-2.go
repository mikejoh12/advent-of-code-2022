package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

type order int
const (
	inOrder order = iota
	equal
	notInOrder
)	

type packets [][]any

func newPackets(f string) *packets {
	chars, _ := os.ReadFile(f)
	pairStrings := strings.Split(string(chars), "\n\n")
	var output packets
	for _, pairStr := range pairStrings {
		packetStrings := strings.Split(pairStr, "\n")
		for _, packetStr := range packetStrings {
			var packet []any
			json.Unmarshal([]byte(packetStr), &packet)
			output = append(output, packet)
		}
	}
	return &output
}

func (p *packets) add(s []any) {
	(*p) = append((*p), s)
}

func (p *packets) findDecoderKey() int {
	result := 1
	for i, p := range *p {
		if reflect.DeepEqual(p, []any{[]any{float64(2)}}) || reflect.DeepEqual(p, []any{[]any{float64(6)}}) {
			result *= i+1
		}
	}
	return result
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
	p := newPackets("input.txt")
	p.add([]any{[]any{float64(2)}})
	p.add([]any{[]any{float64(6)}})
	sort.Slice(*p, func(i, j int) bool {
		return isInOrder((*p)[i], (*p)[j]) == inOrder
	})
	fmt.Println("part 2:", p.findDecoderKey())
}