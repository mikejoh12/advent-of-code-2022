package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id                int
	items             []int
	operation         string
	operationNr       int
	divisibleTest     int
	testTrueReceiver  int
	testFalseReceiver int
	inspections       int
}

func strToInt(s []string) (ints []int) {
	for _, nrStr := range s {
		nr, _ := strconv.Atoi(nrStr)
		ints = append(ints, nr)
	}
	return ints
}

func parse(f string) (monkeys []monkey, cFactor int) {
	s, _ := os.ReadFile(f)
	strMonkeys := strings.Split(string(s), "\n\n")
	cFactor = 1

	for _, monkeyStr := range strMonkeys {
		monkeyLines := strings.Split(monkeyStr, "\n")
		var m = monkey{}
		fmt.Sscanf(monkeyLines[0], "Monkey %d:", &m.id)

		monkeyLines[1] = strings.Replace(monkeyLines[1], "  Starting items: ", "", -1)
		monkeyLines[1] = strings.Replace(monkeyLines[1], ",", "", -1)
		m.items = strToInt(strings.Fields(monkeyLines[1]))

		if monkeyLines[2] == "  Operation: new = old * old" {
			m.operation = "square"
		} else {
			fmt.Sscanf(monkeyLines[2], "  Operation: new = old %s %d", &m.operation, &m.operationNr)
		}

		fmt.Sscanf(monkeyLines[3], "  Test: divisible by %d", &m.divisibleTest)
		fmt.Sscanf(monkeyLines[4], "    If true: throw to monkey %d", &m.testTrueReceiver)
		fmt.Sscanf(monkeyLines[5], "    If false: throw to monkey %d", &m.testFalseReceiver)

		monkeys, cFactor = append(monkeys, m), cFactor*m.divisibleTest
	}

	return monkeys, cFactor
}

func main() {
	monkeys, cFactor := parse("input.txt")

	for i := 0; i < 10000; i++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				worryLevel := item
				switch m.operation {
				case "+":
					worryLevel += m.operationNr
				case "*":
					worryLevel *= m.operationNr
				case "square":
					worryLevel = worryLevel * worryLevel
				}
				worryLevel = worryLevel % cFactor
				monkeys[i].inspections++
				if worryLevel%m.divisibleTest == 0 {
					monkeys[m.testTrueReceiver].items = append(monkeys[m.testTrueReceiver].items, worryLevel)
				} else {
					monkeys[m.testFalseReceiver].items = append(monkeys[m.testFalseReceiver].items, worryLevel)
				}
			}
			monkeys[i].items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}
