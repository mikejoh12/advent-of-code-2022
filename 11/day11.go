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

func parse(f string) []monkey {
	s, _ := os.ReadFile(f)
	strMonkeys := strings.Split(string(s), "\n\n")

	var monkeys []monkey
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

		monkeys = append(monkeys, m)
	}

	return monkeys
}

func main() {
	monkeys := parse("input.txt")

	for i := 0; i < 20; i++ {
		fmt.Println("round i", i)
		fmt.Println("monkeys", monkeys)
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := item
				switch monkey.operation {
				case "+":
					worryLevel += monkey.operationNr
				case "*":
					worryLevel *= monkey.operationNr
				case "square":
					worryLevel = worryLevel * worryLevel
				}
				worryLevel = worryLevel / 3
				monkeys[i].inspections++
				if worryLevel%monkey.divisibleTest == 0 {
					monkeys[monkey.testTrueReceiver].items = append(monkeys[monkey.testTrueReceiver].items, worryLevel)
				} else {
					monkeys[monkey.testFalseReceiver].items = append(monkeys[monkey.testFalseReceiver].items, worryLevel)
				}
			}
			monkeys[i].items = []int{}
		}
	}

	levels := make([]int, 0)
	for _, monkey := range monkeys {
		levels = append(levels, monkey.inspections)
	}
	sort.Slice(levels, func(i, j int) bool {
		return levels[i] > levels[j]
	})
	fmt.Println(levels[0] * levels[1])
}
