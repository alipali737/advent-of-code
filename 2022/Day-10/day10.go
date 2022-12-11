package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	instructions := []string{}
	for fileScanner.Scan() {
		instructions = append(instructions, fileScanner.Text())
	}

	registry := 1
	cycle := 0
	index := 0
	signalStrength := 0
	processing := 0
	msg := ""
	for {
		cycle++

		pixel := ""
		relativeCycle := cycle - (40 * (cycle / 40)) - 1

		if relativeCycle >= registry-1 && relativeCycle <= registry+1 {
			pixel += "█"
		} else {
			pixel += " "
		}

		if cycle%40 == 0 {
			pixel += "\n"
		}

		if cycle == 20 || (cycle-20)%40 == 0 {
			// fmt.Printf("Calculating signal strenght! Registry: %d, Cycle: %d, Strenght: %d, Processing: %d\n", registry, cycle, registry*cycle, processing)
			signalStrength += registry * cycle
		}

		inst := instructions[index]
		fmt.Printf("New instruction: %s, Cycle: %d, Processing: %d, Registry: %d, Sprite: %d-%d, Pixel: %s\n", inst, cycle, processing, registry, registry-1, registry+1, pixel)
		if inst == "noop" {
			index++
		} else {
			processing++
			if processing == 2 {
				strip := strings.Split(inst, " ")
				n, _ := strconv.Atoi(strip[1])
				registry += n
				processing = 0
				index++
			}
		}

		msg += pixel

		if index >= len(instructions) {
			break
		}
	}
	fmt.Print(msg)

	fmt.Printf("\nPart 1 - Sum of Signal Strengths: %d\n", signalStrength)

}
