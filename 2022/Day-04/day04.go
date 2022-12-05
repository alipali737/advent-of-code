package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fullContainments := 0
	partialContainments := 0
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		re := regexp.MustCompile(`\d[\d]*`)
		matches := re.FindAllString(readText, -1)
		pairValues := []int{}
		for _, v := range matches {
			tmp, _ := strconv.Atoi(v)
			pairValues = append(pairValues, tmp)
		}

		if pairValues[0] >= pairValues[2] && pairValues[1] <= pairValues[3] || pairValues[2] >= pairValues[0] && pairValues[3] <= pairValues[1] {
			fullContainments++
		}

		first := makeRange(pairValues[0], pairValues[1])
		second := makeRange(pairValues[2], pairValues[3])

		if sliceIntersectionBetweenTwo(first, second) {
			partialContainments++
		}
	}

	fmt.Printf("Part 1 - Final full containments: %d\n", fullContainments)
	fmt.Printf("Part 2 - Final partial containments: %d\n", partialContainments)

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sliceIntersectionBetweenTwo(s1 []int, s2 []int) bool {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			return true
		}
	}
	return false
}
