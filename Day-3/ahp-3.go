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

func main() {
	step1()
}
