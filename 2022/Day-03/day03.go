package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalP1 := 0
	totalP2 := 0
	groupStrings := []string{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()

		// Part 1
		firstComp := readText[:len(readText)/2]
		secondComp := readText[len(readText)/2:]

		intersection := SliceIntersectionBetweenTwo(strings.Split(firstComp, ""), strings.Split(secondComp, ""))

		totalP1 += CalcItemValue(intersection)

		// Part 2
		groupStrings = append(groupStrings, readText)
		if len(groupStrings) == 3 {
			intersection := SliceIntersectionBetweenThree(strings.Split(groupStrings[0], ""), strings.Split(groupStrings[1], ""), strings.Split(groupStrings[2], ""))
			totalP2 += CalcItemValue(intersection)
			groupStrings = []string{}
		}

	}

	fmt.Printf("Part 1 - Final sum: %d\n", totalP1)
	fmt.Printf("Part 2 - Final sum: %d\n", totalP2)

}

func SliceIntersectionBetweenTwo(s1 []string, s2 []string) rune {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	inter := []string{}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	r := []rune(strings.Join(inter, ""))
	return r[0]
}

func SliceIntersectionBetweenThree(s1 []string, s2 []string, s3 []string) rune {
	hash1 := make(map[string]bool)
	for _, e := range s1 {
		hash1[e] = true
	}
	hash2 := make(map[string]bool)
	for _, e := range s2 {
		hash2[e] = true
	}
	inter := []string{}
	for _, e := range s3 {
		if hash1[e] && hash2[e] {
			inter = append(inter, e)
		}
	}
	r := []rune(strings.Join(inter, ""))
	return r[0]
}

func CalcItemValue(v rune) int {
	ascii := int(v)
	// Uppercase
	if ascii >= 65 && ascii <= 90 {
		return ascii - 38
	}

	// Lowercase
	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	}

	return 0
}
