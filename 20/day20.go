package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	value    int
	startIdx int
	prev     *node
	next     *node
}

type numberList struct {
	head   *node
	length int
}

func NewNumberList(nrs []int) *numberList {
	n := node{
		value:    nrs[0],
		startIdx: 0,
	}
	nl := numberList{
		head:   &n,
		length: len(nrs),
	}
	current := nl.head
	for i := 1; i < len(nrs); i++ {
		newNode := node{
			value:    nrs[i],
			startIdx: i,
			prev:     current,
		}
		current.next = &newNode
		current = current.next
	}
	current.next = nl.head // make circular
	nl.head.prev = current // make circular
	return &nl
}

func (nl *numberList) mix(times int, nrs []int) {
	for i := 0; i < times; i++ {
		for idx := range nrs {
			nl.move(idx)
		}
	}
}

func (nl *numberList) move(idx int) {
	current := nl.head
	for current.startIdx != idx {
		current = current.next
	}
	if current.value == 0 {
		return
	}
	current.next.prev = current.prev
	current.prev.next = current.next
	if nl.head == current {
		nl.head = current.next
	}
	value := current.value
	lim := value % (nl.length - 1)

	switch {
	case current.value > 0:
		for i := 0; i < lim; i++ {
			current = current.next
		}
	case current.value < 0:
		for i := 0; i >= lim; i-- {
			current = current.prev
		}
	}
	new := &node{
		value:    value,
		startIdx: idx,
		prev:     current,
		next:     current.next,
	}
	current.next = new
	current.next.next.prev = new

}

func (nl *numberList) getGroveCoordSum() int {
	counting := false
	count, sum := 0, 0
	current := nl.head
	for count <= 3000 {
		if count == 1000 || count == 2000 || count == 3000 {
			sum += current.value
		}
		if current.value == 0 {
			counting = true
		}
		if counting {
			count++
		}
		current = current.next
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nrs1, nrs2 []int

	for scanner.Scan() {
		t := scanner.Text()
		var nr int
		fmt.Sscan(t, &nr)
		nrs1 = append(nrs1, nr)
		nrs2 = append(nrs2, nr*811589153)
	}

	l1, l2 := NewNumberList(nrs1), NewNumberList(nrs2)

	l1.mix(1, nrs1)
	l2.mix(10, nrs2)

	fmt.Println(l1.getGroveCoordSum())
	fmt.Println(l2.getGroveCoordSum())
}