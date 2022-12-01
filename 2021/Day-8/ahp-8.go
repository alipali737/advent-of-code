package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ImportFile(path string) (signals [][]string, outputs [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), "|")
		signals = append(signals, strings.Split(tmp[0], " ")[:10])
		outputs = append(outputs, strings.Split(tmp[1], " ")[1:])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return signals, outputs
}

func main() {
	signals, outputs := ImportFile("input.txt")
	// fmt.Printf("Signals[0] = %#v\n", signals[0])
	// fmt.Printf("Outputs[0] = %#v\n", outputs[0])
	step1(outputs)
	step2(signals, outputs)

}

func step1(outputs [][]string) {
	counter := 0
	for _, set := range outputs {
		for _, v := range set {
			length := len(v)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				counter++
			}
		}
	}
	fmt.Printf("Step 1 counter = %d\n", counter)
}

func matchNumberOfSegments(s string, pattern string, n int) bool {
	splitPattern := strings.Split(pattern, "")
	counter := 0
	for _, v := range splitPattern {
		if regexp.MustCompile(v).MatchString(s) {
			counter++
		}
	}
	return counter >= n
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func step2(signals [][]string, outputs [][]string) {

	total := 0
	for index := 0; index < len(signals); index++ {
		fmt.Printf("----- New Segment -----\n")
		matched := make([]string, 10)
		currentSignals := signals[index]
		currentOutputs := outputs[index]

		// Calc 1, 4, 7, 8
	lengthCalc:
		for _, v := range currentSignals {
			// Number 1
			if len(v) == 2 {
				matched[1] = v
				continue lengthCalc
			}

			// Number 7
			if len(v) == 3 {
				matched[7] = v
				continue lengthCalc
			}

			// Number 4
			if len(v) == 4 {
				matched[4] = v
				continue lengthCalc
			}

			// Number 8
			if len(v) == 7 {
				matched[8] = v
				continue lengthCalc
			}
		}

	otherCalc:
		for _, v := range currentSignals {
			// Number 3, 5, 2
			if len(v) == 5 {
				if matchNumberOfSegments(v, matched[1], 2) {
					matched[3] = v
					continue otherCalc
				}
				if matchNumberOfSegments(v, matched[4], 3) {
					matched[5] = v
					continue otherCalc
				}
				if matchNumberOfSegments(v, matched[4], 2) {
					matched[2] = v
					continue otherCalc
				}
			}

			// Number 0, 9, 6
			if len(v) == 6 {
				if matchNumberOfSegments(v, matched[1], 2) && !matchNumberOfSegments(v, matched[4], 4) {
					matched[0] = v
					continue otherCalc
				}

				if matchNumberOfSegments(v, matched[4], 4) {
					matched[9] = v
					continue otherCalc
				}

				matched[6] = v
			}
			if !contains(matched, v) {
				fmt.Printf("Couldn't find a assignment for \"%s\"\n", v)
			}
		}

		// fmt.Printf("Signals: %#v\n", matched)

		tmp := ""
		for _, v := range currentOutputs {
			for index, signal := range matched {
				if len(v) == len(signal) && matchNumberOfSegments(v, signal, len(signal)) {
					tmp += fmt.Sprint(index)
				}
			}
		}

		// fmt.Printf("Decoded : %#v : %s\n", currentOutputs, tmp)
		tmp2, _ := strconv.Atoi(tmp)
		total += tmp2
	}

	fmt.Printf("Step 2 total = %d\n", total)

	/*

		Get these using their unique lengths: 1, 7, 4, 8
		Determine 3 using 1.. shares all segments
		Determine 5 using 4.. shares all but 1 segment
		Determine 2 using 4.. shares all but 2 segments
		Determine 0 using 1.. shares all segments
		Determine 9 using 4.. shares all segments
		Determine 6 using None.. final value

		Lengths:
		0 : 6
		1 : 2
		2 : 5
		3 : 5
		4 : 4
		5 : 5
		6 : 6
		7 : 3
		8 : 7
		9 : 6
	*/
}
