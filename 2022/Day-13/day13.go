package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	leftString := ""
	rightString := ""
	index := 0
	p1 := 0
	lines := []string{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()

		if readText != "" {
			lines = append(lines, readText)
		}

		switch readText {
		case "":

			leftPacket := parse(leftString)
			rightPacket := parse(rightString)
			// fmt.Printf(" Left Packet: %v\n", leftPacket)
			// fmt.Printf("Right Packet: %v\n\n", rightPacket)

			if compare(leftPacket, rightPacket) == -1 {
				p1 += index + 1
			}

			leftString = ""
			rightString = ""
			index++
		default:
			if len(leftString) == 0 {
				leftString = readText
			} else {
				rightString = readText
			}
		}
	}

	leftPacket := parse(leftString)
	rightPacket := parse(rightString)
	// fmt.Printf(" Left Packet: %v\n", leftPacket)
	// fmt.Printf("Right Packet: %v\n\n", rightPacket)

	if compare(leftPacket, rightPacket) == -1 {
		p1 += index + 1
	}

	fmt.Printf("Part 1 - Total: %d\n", p1)

	lines = append(lines, "[[2]]", "[[6]]")

	sort.Slice(lines, func(i, j int) bool {
		return compare(parse(lines[i]), parse(lines[j])) < 0
	})

	var twoDividerIdx int
	var sixDividerIdx int

	for i, curr := range lines {
		if curr == "[[2]]" {
			twoDividerIdx = i + 1
		}
		if curr == "[[6]]" {
			sixDividerIdx = i + 1
		}
	}

	fmt.Printf("%d\n", twoDividerIdx*sixDividerIdx)

}

func parse(s string) []any {
	var data []any
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func isList(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), "[]")
}

func compare(a any, b any) int {
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	aIsList := isList(aType)
	bIsList := isList(bType)

	if !aIsList && !bIsList {
		// Both numbers
		aNum := a.(float64)
		bNum := b.(float64)

		if aNum < bNum {
			return -1
		} else if aNum > bNum {
			return 1
		} else {
			return 0
		}
	}

	if aIsList && bIsList {
		// Both lists
		aList := a.([]any)
		bList := b.([]any)

		compareLen := 0
		if len(aList) < len(bList) {
			compareLen = len(aList)
		} else {
			compareLen = len(bList)
		}

		for i := 0; i < compareLen; i++ {
			c := compare(aList[i], bList[i])
			if c != 0 {
				return c
			}
		}

		if len(aList) < len(bList) {
			return -1
		} else if len(aList) > len(bList) {
			return 1
		} else {
			return 0
		}
	}

	if aIsList && !bIsList {
		bNum := b.(float64)
		return compare(a, []any{bNum})
	}

	if !aIsList && bIsList {
		aNum := a.(float64)
		return compare([]any{aNum}, b)
	}

	return 0
}
