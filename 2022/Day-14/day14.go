package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// get outer bounds
	// sort by x & y and take min/max for bounds
	// get map of all points to true
	// simulate sand, if it go outside of bounds it is done
	// increment count for each sand rested
	startPoint := point{x: 500, y: 0}
	filledSpots := map[point]bool{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		splitLine := strings.Split(readText, " -> ")
		splitPoints := []point{}
		for _, p := range splitLine {
			point := parsePoint(p)
			splitPoints = append(splitPoints, point)
			filledSpots[point] = true
		}
		for i := 1; i < len(splitPoints); i++ {
			generatedPoints := getLineCoords(splitPoints[i-1], splitPoints[i])
			for _, p := range generatedPoints {
				filledSpots[p] = true
			}
		}
	}
	// fmt.Println(filledSpots)

	top_left, top_right, bottom_left, bottom_right := getBounds(filledSpots)
	// printBoard(filledSpots, top_left, bottom_right)

	sandRested := 0
	for {
		var hasRested bool
		filledSpots, hasRested = simulateSand(filledSpots, startPoint, top_left, top_right, bottom_left, bottom_right, true)
		if hasRested {
			sandRested++
		} else {
			break
		}
	}

	fmt.Printf("Part 1 - Units of sand rested: %d\n", sandRested)

	for {
		var hasRested bool
		filledSpots, hasRested = simulateSand(filledSpots, startPoint, top_left, top_right, bottom_left, bottom_right, false)
		if hasRested {
			sandRested++
		} else {
			break
		}
	}

	fmt.Printf("Part 2 - Units of sand rested: %d\n", sandRested)

}

type point struct {
	x, y int
}

func parsePoint(s string) point {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return point{x, y}
}

func getLineCoords(p1, p2 point) []point {
	points := []point{}
	if p1.x < p2.x {
		// Line going right
		// fmt.Printf("Creating line right! p1: %v, p2: %v\n", p1, p2)
		for i := p1.x; i < p2.x+1; i++ {
			points = append(points, point{i, p1.y})
		}
		return points
	}
	if p1.x > p2.x {
		// Line going left
		// fmt.Printf("Creating line left! p1: %v, p2: %v\n", p1, p2)
		for i := p2.x; i < p1.x+1; i++ {
			points = append(points, point{i, p1.y})
		}
		return points
	}

	if p1.y < p2.y {
		// Line going down
		// fmt.Printf("Creating line down! p1: %v, p2: %v\n", p1, p2)
		for i := p1.y; i < p2.y+1; i++ {
			points = append(points, point{p1.x, i})
		}
		return points
	}
	if p1.y > p2.y {
		//Line going up
		// fmt.Printf("Creating line up! p1: %v, p2: %v\n", p1, p2)
		for i := p2.y; i < p1.y+1; i++ {
			points = append(points, point{p1.x, i})
		}
		return points
	}
	return points
}

func getBounds(m map[point]bool) (point, point, point, point) {
	keys := []point{}
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].x < keys[j].x
	})
	leftX := keys[0].x
	rightX := keys[len(keys)-1].x

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].y < keys[j].y
	})
	topY := keys[0].y
	bottomY := keys[len(keys)-1].y

	top_left := point{x: leftX, y: topY}
	top_right := point{x: rightX, y: topY}
	bottom_left := point{x: leftX, y: bottomY}
	bottom_right := point{x: rightX, y: bottomY}

	return top_left, top_right, bottom_left, bottom_right
}

func simulateSand(filledSpots map[point]bool, startPoint, top_left, top_right, bottom_left, bottom_right point, endless bool) (map[point]bool, bool) {
	notRested := true
	sand := point{startPoint.x, startPoint.y}
	if _, ok := filledSpots[startPoint]; ok {
		return filledSpots, false
	}
	for notRested {
		// fmt.Printf("Sand: %v\n", sand)
		// If can go down
		if _, ok := filledSpots[point{sand.x, sand.y + 1}]; !ok {
			sand.y++
			// fmt.Printf("Sand can go down! Sand: %v\n", sand)
		} else if _, ok := filledSpots[point{sand.x - 1, sand.y + 1}]; !ok {
			// If can go left diag
			sand.x--
			sand.y++
			// fmt.Printf("Sand can go left diag! Sand: %v\n", sand)
		} else if _, ok := filledSpots[point{sand.x + 1, sand.y + 1}]; !ok {
			// If can go right diag
			sand.x++
			sand.y++
			// fmt.Printf("Sand can go right diag! Sand: %v\n", sand)
		} else {
			notRested = false
			// fmt.Printf("Sand rested! Sand: %v\n", sand)
		}

		if endless {
			out := outOfBounds(sand, top_left, top_right, bottom_left, bottom_right)
			if out {
				// fmt.Printf("Sand out of bounds! Sand: %v, Out: %t\n", sand, out)
				return filledSpots, false
			}
		} else {
			if sand.y == bottom_left.y+1 {
				notRested = false
			}
		}
	}

	filledSpots[sand] = true
	return filledSpots, true
}

func outOfBounds(sand, top_left, top_right, bottom_left, bottom_right point) bool {
	return sand.x < top_left.x || sand.x > top_right.x || sand.y > bottom_left.y
}

func printBoard(filledSpots map[point]bool, top_left, bottom_right point) {
	for row := top_left.y; row < bottom_right.y+1; row++ {
		line := ""
		for column := top_left.x; column < bottom_right.x+1; column++ {
			if column == 500 && row == 0 {
				line += "+"
			} else if filledSpots[point{column, row}] {
				line += "#"
			} else {
				line += "."
			}

		}
		fmt.Println(line)
	}
}
