package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type operation int
const (
	add		operation = iota
	subtract
	multiply
	divide
	none
)

type monkey struct {
	operation 	operation
	monkey1		string
	monkey2		string
	value		int
}

type monkeys map[string]monkey

func newMonkeys() *monkeys {
	m := make(monkeys)
	return &m
}

func (m *monkeys) yell(name string, invalidOnHumn bool) (int, bool) {
	current := (*m)[name]
	if invalidOnHumn && name == "humn" {
		return -1, false
	} else if current.operation == none {
		return current.value, true
	}

	res1, ok1 := m.yell(current.monkey1, invalidOnHumn)
	res2, ok2 := m.yell(current.monkey2, invalidOnHumn)
	if !ok1 || !ok2 {
		return -1, false
	}

	switch current.operation {
	case add:
		return res1+res2, true
	case subtract:
		return res1-res2, true
	case multiply:
		return res1*res2, true
	case divide:
		return res1/res2, true
	}
	return -1, false
}

func (m *monkeys) findHuman(subtree string, curNr int) int{
	current := (*m)[subtree]
	if subtree == "humn" {
		return curNr
	}
	res1, ok1 := m.yell(current.monkey1, true)
	res2, ok2 := m.yell(current.monkey2, true)
	var unknown string
	var nr int
	if ok1 {
		unknown = current.monkey2
		nr = res1
	} else if ok2 {
		unknown = current.monkey1
		nr = res2
	}
	switch {
	case subtree == "root":
		return m.findHuman(unknown, nr)
	case current.operation == add:
		return m.findHuman(unknown, curNr - nr)
	case current.operation == subtract:
		if ok1 {
			return m.findHuman(unknown, nr-curNr)
		} else {
			return m.findHuman(unknown, nr+curNr)
		}
	case current.operation == multiply:
		return m.findHuman(unknown, curNr/nr)
	case current.operation == divide:
		if ok1 {
			return m.findHuman(unknown, nr/curNr)
		} else {
			return m.findHuman(unknown, nr*curNr)
		}
	}
	return -1
}

func (m *monkeys) add(name string, monkey monkey) {
	(*m)[name] = monkey
} 

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	monkeys := newMonkeys()

	for scanner.Scan() {
		t := scanner.Text()
		data := strings.Split(t, ": ")
		name := data[0]
		newMonkey := monkey{}
		switch {
		case strings.Contains(data[1], "+"):
			newMonkey.operation = add
			fmt.Sscanf(data[1], "%s + %s", &newMonkey.monkey1, &newMonkey.monkey2)
		case strings.Contains(data[1], "-"):
			newMonkey.operation = subtract
			fmt.Sscanf(data[1], "%s - %s", &newMonkey.monkey1, &newMonkey.monkey2)
		case strings.Contains(data[1], "*"):
			newMonkey.operation = multiply
			fmt.Sscanf(data[1], "%s * %s", &newMonkey.monkey1, &newMonkey.monkey2)
		case strings.Contains(data[1], "/"):
			newMonkey.operation = divide
			fmt.Sscanf(data[1], "%s / %s", &newMonkey.monkey1, &newMonkey.monkey2)
		default:
			fmt.Sscanf(data[1], "%d", &newMonkey.value)
			newMonkey.operation = none
		}
		monkeys.add(name, newMonkey)
	}

	part1, _ := monkeys.yell("root", false)
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", monkeys.findHuman("root", -1))
}