package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Part 1 - Visited points: %d\n", solve(2))
	fmt.Printf("Part 2 - Visited points: %d\n", solve(10))

}

func solve(ropeLength int) int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tailUniquePositions := map[point]bool{point{x: 0, y: 0}: true}
	headPos := point{x: 0, y: 0}
	ropePositions := []point{}
	for knot := 0; knot < ropeLength-1; knot++ {
		ropePositions = append(ropePositions, point{x: 0, y: 0})
	}

	for fileScanner.Scan() {
		readText := fileScanner.Text()
		split := strings.Split(readText, " ")
		steps, _ := strconv.Atoi(split[1])
		tailPosSlice := []point{}
		headPos, ropePositions, tailPosSlice = moveHead(split[0], steps, headPos, ropePositions)
		for _, pos := range tailPosSlice {
			tailUniquePositions[pos] = true
		}
	}

	return len(tailUniquePositions)

}

type point struct {
	x, y int
}

func (p point) isTouchingPoint(other point) bool {
	var diff point
	diff.x = int(math.Abs(float64(p.x - other.x)))
	diff.y = int(math.Abs(float64(p.y - other.y)))
	if diff.x <= 1 && diff.x >= 0 && diff.y <= 1 && diff.y >= 0 {
		return true
	}
	return false
}

func moveHead(direction string, steps int, headPos point, knotPositions []point) (point, []point, []point) {
	var moveOffset point
	tailPositions := []point{}
	switch direction {
	case "R":
		moveOffset = point{x: 1, y: 0}
	case "L":
		moveOffset = point{x: -1, y: 0}
	case "D":
		moveOffset = point{x: 0, y: -1}
	case "U":
		moveOffset = point{x: 0, y: 1}
	}

	fmt.Printf("== %s %d ==\n", direction, steps)
	for i := 0; i < steps; i++ {
		headPos.x += moveOffset.x
		headPos.y += moveOffset.y
		stepKnotPositions := append([]point{headPos}, knotPositions...)
		// fmt.Printf("stepKnotPositions: %v\n", stepKnotPositions)
		// fmt.Printf("    knotPositions: %v\n", knotPositions)
		for j := 1; j < len(stepKnotPositions); j++ {
			res := moveKnot(stepKnotPositions[j-1], stepKnotPositions[j])
			stepKnotPositions[j] = res
			knotPositions[j-1] = res
			tailPositions = append(tailPositions, stepKnotPositions[len(stepKnotPositions)-1])
		}
		fmt.Println("")
	}

	return headPos, knotPositions, tailPositions

}

func moveKnot(headPos point, knotPos point) point {
	if knotPos.isTouchingPoint(headPos) {
		// fmt.Printf("Knot touching! Knot before: [%d,%d], Knot: [%d,%d]\n", headPos.x, headPos.y, knotPos.x, knotPos.y)
		return knotPos
	}

	var diff point
	diff.x = headPos.x - knotPos.x
	if diff.x > 1 {
		diff.x -= 1
	} else if diff.x < -1 {
		diff.x += 1
	}

	diff.y = headPos.y - knotPos.y
	if diff.y > 1 {
		diff.y -= 1
	} else if diff.y < -1 {
		diff.y += 1
	}

	knotPos.x += diff.x
	knotPos.y += diff.y

	// fmt.Printf("Knot moving! Knot before: [%d,%d], Knot: [%d,%d], Diff: [%d, %d]\n", headPos.x, headPos.y, knotPos.x, knotPos.y, diff.x, diff.y)

	return knotPos

}
