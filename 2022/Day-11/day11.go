package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	monkeys := map[int]monkey{}

	currentBlock := []string{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		if readText == "" {
			id, m := parseMonkey(currentBlock)
			monkeys[id] = m
			currentBlock = []string{}
		} else {
			currentBlock = append(currentBlock, readText)
		}
	}

	id, m := parseMonkey(currentBlock)
	monkeys[id] = m

	monkeyLCM := 1
	for _, m := range monkeys {
		monkeyLCM *= m.test.conditionValue
	}
	// for i, m := range monkeys {
	// 	fmt.Printf("%d : %v\n", i, m)
	// }

	for round := 0; round < 10000; round++ {
		for mIndex := 0; mIndex < len(monkeys); mIndex++ {
			monkey := monkeys[mIndex]
			monkeyItems := monkey.items
			for iIndex := 0; iIndex < len(monkeyItems); iIndex++ {
				worryLevel := monkeyItems[iIndex]
				switch monkey.operation[0] {
				case "+":
					n, _ := strconv.Atoi(monkey.operation[1])
					worryLevel += n
				case "*":
					if monkey.operation[1] == "old" {
						worryLevel *= worryLevel
					} else {
						n, _ := strconv.Atoi(monkey.operation[1])
						worryLevel *= n
					}
				}

				if worryLevel > monkeyLCM {
					worryLevel %= monkeyLCM
				}

				if worryLevel%monkey.test.conditionValue == 0 {
					monkeys[mIndex], monkeys[monkey.test.ifTrue] = throwItem(monkeys, worryLevel, mIndex, monkey.test.ifTrue)
				} else {
					monkeys[mIndex], monkeys[monkey.test.ifFalse] = throwItem(monkeys, worryLevel, mIndex, monkey.test.ifFalse)
				}
			}
		}

		if (round+1)%1000 == 0 || round+1 == 1 || round+1 == 20 {
			fmt.Printf("== After round %d ==\n", round+1)
			for mi := 0; mi < len(monkeys); mi++ {
				fmt.Printf("Monkey %d inspected items %d times.\n", mi, monkeys[mi].itemsInspected)
				if len(monkeys[mi].items) > 0 {
					fmt.Printf("-> First Item: %d\n", monkeys[mi].items[0])
				} else {
					fmt.Println("-> No items")
				}
			}
			fmt.Println("")
		}

	}

	// for mi := 0; mi < len(monkeys); mi++ {
	// 	fmt.Printf("Monkey %d inspected items %d times.\n", mi, monkeys[mi].itemsInspected)
	// }

	sorted := byItemsInspected{}
	for _, monkey := range monkeys {
		sorted = append(sorted, monkey)
	}

	sort.Sort(sorted)
	fmt.Printf("Part 1 - Total monkey business: %d\n", sorted[len(sorted)-1].itemsInspected*sorted[len(sorted)-2].itemsInspected)

}

func parseMonkey(lines []string) (int, monkey) {
	m := monkey{
		itemsInspected: 0,
	}

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	tmp := strings.Split(lines[0], "")
	id, _ := strconv.Atoi(tmp[len(tmp)-2])

	tmp = strings.Split(strings.TrimPrefix(lines[1], "Starting items: "), ", ")
	for _, v := range tmp {
		n, _ := strconv.Atoi(v)
		m.items = append(m.items, n)
	}

	m.operation = strings.Split(strings.TrimPrefix(lines[2], "Operation: new = old "), " ")

	t := throwTest{}
	n, _ := strconv.Atoi(strings.TrimPrefix(lines[3], "Test: divisible by "))
	t.conditionValue = n

	n, _ = strconv.Atoi(strings.TrimPrefix(lines[4], "If true: throw to monkey "))
	t.ifTrue = n

	n, _ = strconv.Atoi(strings.TrimPrefix(lines[5], "If false: throw to monkey "))
	t.ifFalse = n

	m.test = t

	return id, m

}

func throwItem(monkeys map[int]monkey, item, fromMonkeyID, toMonkeyID int) (monkey, monkey) {
	fromMonkey := monkeys[fromMonkeyID]
	toMonkey := monkeys[toMonkeyID]
	fromMonkey.items = fromMonkey.items[1:]
	fromMonkey.itemsInspected++
	toMonkey.items = append(toMonkey.items, item)

	return fromMonkey, toMonkey
}

type monkey struct {
	items          []int
	operation      []string
	test           throwTest
	itemsInspected int
}

type throwTest struct {
	conditionValue int
	ifTrue         int
	ifFalse        int
}

type byItemsInspected []monkey

func (b byItemsInspected) Len() int      { return len(b) }
func (b byItemsInspected) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byItemsInspected) Less(i, j int) bool {
	m1value := b[i].itemsInspected
	m2value := b[j].itemsInspected

	return m1value < m2value
}
