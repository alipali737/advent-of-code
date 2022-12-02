package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		Opponent:
		- A : Rock
		- B : Paper
		- C : Scissors

		You:
		- X : Rock
		- Y : Paper
		- Z : Scissors

		Points:
		Single round score = shape_you_pick + round_outcome
		Shape you pick:
		- 1 : Rock
		- 2 : Paper
		- 3 : Scissors
		Round outcome:
		- 0 : Loss
		- 3 : Draw
		- 6 : Win
	*/
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	outcomes := map[string]int{
		"AX": 4, "AY": 8, "AZ": 3,
		"BX": 1, "BY": 5, "BZ": 9,
		"CX": 7, "CY": 2, "CZ": 6,
	}

	// Op Letter -> value of your play
	lose := map[string]int{
		"A": 3,
		"B": 1,
		"C": 2,
	}
	draw := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	win := map[string]int{
		"A": 2,
		"B": 3,
		"C": 1,
	}

	totalScoreP1 := 0
	totalScoreP2 := 0
	for fileScanner.Scan() {
		readText := fileScanner.Text()

		// Part 1
		splitText := strings.Split(readText, " ")
		combined := splitText[0] + splitText[1]
		totalScoreP1 += outcomes[combined]

		// Part 2
		switch splitText[1] {
		case "X":
			totalScoreP2 += 0 + lose[splitText[0]]
		case "Y":
			totalScoreP2 += 3 + draw[splitText[0]]
		case "Z":
			totalScoreP2 += 6 + win[splitText[0]]
		}
	}

	fmt.Printf("Part 1 - Final outcome score: %d\n", totalScoreP1)
	fmt.Printf("Part 2 - Final outcome score: %d\n", totalScoreP2)
}
