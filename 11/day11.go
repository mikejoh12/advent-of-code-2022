package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type operation int

const (
	add operation = iota
	multiply
	square
)

type monkey struct {
	items         []int
	operation     operation
	operationNr   int
	divisibleTest int
	trueReceiver  int
	falseReceiver int
	inspections   int
}

type monkeyGroup struct {
	monkeys      []monkey
	commonFactor int
}

func newMonkeyGroup(f string) *monkeyGroup {
	monkeys := new(monkeyGroup)
	s, _ := os.ReadFile(f)
	strMonkeys := strings.Split(string(s), "\n\n")
	monkeys.commonFactor = 1

	for _, monkeyStr := range strMonkeys {
		monkeyLines := strings.Split(monkeyStr, "\n")
		var m = monkey{}
		monkeyLines[1] = strings.Replace(monkeyLines[1], "  Starting items: ", "", -1)
		monkeyLines[1] = strings.Replace(monkeyLines[1], ",", "", -1)
		m.items = strToInt(strings.Fields(monkeyLines[1]))

		switch {
		case monkeyLines[2] == "  Operation: new = old * old":
			m.operation = square
		case strings.Contains(monkeyLines[2], "*"):
			m.operation = multiply
			fmt.Sscanf(monkeyLines[2], "  Operation: new = old * %d", &m.operationNr)
		case strings.Contains(monkeyLines[2], "+"):
			m.operation = add
			fmt.Sscanf(monkeyLines[2], "  Operation: new = old + %d", &m.operationNr)
		}

		fmt.Sscanf(monkeyLines[3], "  Test: divisible by %d", &m.divisibleTest)
		fmt.Sscanf(monkeyLines[4], "    If true: throw to monkey %d", &m.trueReceiver)
		fmt.Sscanf(monkeyLines[5], "    If false: throw to monkey %d", &m.falseReceiver)
		monkeys.monkeys, monkeys.commonFactor = append(monkeys.monkeys, m), monkeys.commonFactor*m.divisibleTest
	}

	return monkeys
}

func (mGroup *monkeyGroup) playRounds(n int, divBy3 bool) {
	for i := 0; i < n; i++ {
		for i, m := range mGroup.monkeys {
			for _, item := range m.items {
				worryLevel := item
				switch m.operation {
				case add:
					worryLevel += m.operationNr
				case multiply:
					worryLevel *= m.operationNr
				case square:
					worryLevel *= worryLevel
				}
				if divBy3 {
					worryLevel /= 3
				}
				worryLevel = worryLevel % mGroup.commonFactor
				mGroup.monkeys[i].inspections++
				if worryLevel%m.divisibleTest == 0 {
					mGroup.monkeys[m.trueReceiver].items = append(mGroup.monkeys[m.trueReceiver].items, worryLevel)
				} else {
					mGroup.monkeys[m.falseReceiver].items = append(mGroup.monkeys[m.falseReceiver].items, worryLevel)
				}
			}
			mGroup.monkeys[i].items = []int{}
		}
	}
}

func (mGroup *monkeyGroup) getMonkeyBusiness() int {
	sort.Slice(mGroup.monkeys, func(i, j int) bool {
		return mGroup.monkeys[i].inspections > mGroup.monkeys[j].inspections
	})
	return mGroup.monkeys[0].inspections * mGroup.monkeys[1].inspections
}

func strToInt(s []string) (ints []int) {
	for _, nrStr := range s {
		nr, _ := strconv.Atoi(nrStr)
		ints = append(ints, nr)
	}
	return ints
}

func main() {
	monkeys1, monkeys2 := newMonkeyGroup("input.txt"), newMonkeyGroup("input.txt")
	monkeys1.playRounds(20, true)
	monkeys2.playRounds(10000, false)
	fmt.Println("part 1:", monkeys1.getMonkeyBusiness())
	fmt.Println("part 2:", monkeys2.getMonkeyBusiness())
}
