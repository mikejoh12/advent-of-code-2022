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

func (m *monkeys) yell(name string) int {
	current := (*m)[name]
	switch current.operation {
	case none:
		return current.value
	case add:
		result := m.yell(current.monkey1) + m.yell(current.monkey2)
		return result
	case subtract:
		result := m.yell(current.monkey1) - m.yell(current.monkey2)
		return result
	case multiply:
		result := m.yell(current.monkey1) * m.yell(current.monkey2)
		return result
	case divide:
		result := m.yell(current.monkey1) / m.yell(current.monkey2)
		return result
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

	fmt.Println(monkeys.yell("root"))
}