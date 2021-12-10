package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mostCommonBit(column []string) string {
	counter := make(map[string]int)
	for _, c := range column {
		counter[string(c)]++
	}

	// fmt.Println(counter)

	if counter["0"] > counter["1"] {
		return "0"
	} else {
		return "1"
	}
}

func GetColumn(lines [][]string, columnIndex int) (column []string) {
	column = make([]string, 0)
	for _, row := range lines {
		column = append(column, row[columnIndex])
	}
	return
}

func step1() {
	var gammaRate, epsilonRate, powerConsumption int64
	var lines [][]string

	// GammaRate = most common bit in the corresponding position (column) of each row
	// EpsilonRate = least common bit in the corresponding position (inverted gamma rate)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v", err)
	}

	// fmt.Printf("len=%d cap=%d %v\n", len(lines), cap(lines), lines)
	var binary string
	for i := 0; i <= len(lines[0])-1; i++ {
		binary += mostCommonBit(GetColumn(lines, i))
	}
	gammaRate, _ = strconv.ParseInt(binary, 2, 64)
	fmt.Printf("gammaRate = %v, binary = %s\n", gammaRate, binary)

	temp := ""
	for i := 0; i < len(binary); i++ {
		if string(binary[i]) == "0" {
			temp += "1"
		} else {
			temp += "0"
		}
	}

	epsilonRate, _ = strconv.ParseInt(temp, 2, 64)
	fmt.Printf("epsilon = %v, binary = %s\n", epsilonRate, temp)

	powerConsumption = gammaRate * epsilonRate
	fmt.Printf("Power Consumption = %v\n", powerConsumption)
}

func getOxygenRating(lines [][]string, columnIndex int) []string {
	if len(lines) == 1 {
		return lines[0]
	}

	if columnIndex == len(lines[0]) {
		fmt.Printf("Reached end of cycle without final result: %s", lines)
	}

	var validLines [][]string
	mcb := mostCommonBit(GetColumn(lines, columnIndex))
	for _, line := range lines {
		if line[columnIndex] == mcb {
			validLines = append(validLines, line)
		}
	}
	columnIndex++
	return getOxygenRating(validLines, columnIndex)
}

func getCO2Rating(lines [][]string, columnIndex int) []string {
	if len(lines) == 1 {
		return lines[0]
	}

	if columnIndex == len(lines[0]) {
		fmt.Printf("Reached end of cycle without final result: %s", lines)
	}

	var validLines [][]string
	mcb := mostCommonBit(GetColumn(lines, columnIndex))

	if mcb == "1" {
		mcb = "0"
	} else {
		mcb = "1"
	}

	for _, line := range lines {
		if line[columnIndex] == mcb {
			validLines = append(validLines, line)
		}
	}
	columnIndex++
	return getCO2Rating(validLines, columnIndex)
}

func step2() {
	var lines [][]string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}

	binary := strings.Join(getOxygenRating(lines, 0), "")
	oxygenRating, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Printf("Oxygen Rating: %d, Binary: %s\n", oxygenRating, binary)

	binary = strings.Join(getCO2Rating(lines, 0), "")
	CO2Rating, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Printf("Oxygen Rating: %d, Binary: %s\n", CO2Rating, binary)

	fmt.Printf("Life Support Rating: %d\n", CO2Rating*oxygenRating)

}

func main() {
	step1()
	step2()
}
