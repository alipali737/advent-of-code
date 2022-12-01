package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	maximum := []int{0, 0, 0}
	currentTotal := 0
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		fmt.Printf("Current text: %s\n", readText)
		if readText == "" || readText == "-" {
			fmt.Printf("Current Total: %d\n", currentTotal)
			for i := 0; i < 3; i++ {
				if currentTotal >= maximum[i] {
					maximum = InsertIntoSlice(maximum, currentTotal, i)
					break
				}
			}
			fmt.Printf("Maximum: %d\n\n", maximum)
			currentTotal = 0
		} else {
			intText, err := strconv.Atoi(readText)
			if err != nil {
				fmt.Printf("failed to convert read text to int: %s\n", err)
			}
			currentTotal += intText
		}
	}

	fmt.Printf("--- Top %d most calories ---\n", len(maximum))
	for i := 0; i < len(maximum); i++ {
		fmt.Printf("%d : %d calories\n", i+1, maximum[i])
	}

	fmt.Printf("\nTotal : %d\n", SumSlice(maximum))

}

func InsertIntoSlice(s []int, v int, i int) []int {
	shifted := s[i : len(s)-1]
	shifted = append([]int{v}, shifted...)
	shifted = append(s[:i], shifted...)
	return shifted
}

func SumSlice(s []int) int {
	total := 0
	for _, value := range s {
		total += value
	}
	return total
}
