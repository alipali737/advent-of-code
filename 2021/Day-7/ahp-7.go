package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func ImportFile(path string) (crabs []int) {
	bytes, _ := ioutil.ReadFile(path)
	puzzleInput := strings.Split(string(bytes), ",")
	for _, v := range puzzleInput {
		tmp, _ := strconv.Atoi(v)
		crabs = append(crabs, tmp)
	}

	return crabs
}

func main() {
	crabs := ImportFile("input.txt")
	fmt.Println("-------- Step 1 --------")
	step1(crabs)
	fmt.Println("-------- Step 2 --------")
	step2(crabs)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func step1(crabs []int) {
	var lowest, aligned int = 99999999, 0
iter:
	for _, target := range crabs {
		fuel := 0
		for _, current := range crabs {
			fuel += int(math.Abs(float64(current - target)))
			if fuel > lowest {
				continue iter
			}
		}
		if fuel < lowest {
			lowest = fuel
			aligned = target
		}
	}

	fmt.Printf("Lowest Fuel Count = %d when aligned to %d", lowest, aligned)
}

func step2(crabs []int) {
	var lowest, aligned int = 99999999, 0
	_, max := MinMax(crabs)
iter:
	for target := 0; target < max; target++ {
		fuel := 0
		for _, current := range crabs {
			n := math.Abs(float64(current - target))
			fuel += int(((1.0 / 2.0) * n) * (n + 1))
			if fuel > lowest {
				continue iter
			}
		}
		if fuel < lowest {
			lowest = fuel
			aligned = target
		}
	}

	fmt.Printf("Lowest Fuel Count = %d when aligned to %d", lowest, aligned)
}
