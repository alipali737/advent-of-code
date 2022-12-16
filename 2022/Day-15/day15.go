package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sampleInput string

//go:embed input.txt
var realInput string

func main() {
	input := realInput
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	p1 := 0
	p2 := 0

	/*
		Basic Idea for Part 1: For each sensor,
		calculate the manhattan distance to its beacon.
		Then determine the min and max x coordinates possible for
		line 2000000 which are
		sensor.xCoordinate +- (distance (2000000-sensor.yCoord)) +1
		Put that range into a List,
		sort it by minX,
		then sum the ranges (ignoring overlap).

		Part 2: Just do Part 1 for variable y instead of 2000000
		and stop when the maximum x coordinate already covered
		by a previous range is smaller than the minimum x coordinate
		of the current range (list is sorted by minX
		so there cant be another range covering the gap)
	*/

	sensors := []sensor{}
	lineNo := 2000000
	maxY := 4000000
	// lineNo := 10
	// maxY := 25

	for _, line := range lines {
		sensor := parse(line)
		sensors = append(sensors, sensor)
		// fmt.Printf("Sensor: [x: %d, y: %d], Beacon: [x: %d, y: %d]\n", sensor.position.x, sensor.position.y, sensor.beacon.x, sensor.beacon.y)
	}
	for yIndex := 0; yIndex < maxY; yIndex++ {
		beaconRanges := []point{}
	OUTER:
		for _, sensor := range sensors {
			distance := manhattanDistance(sensor.position, sensor.beacon)
			if inYRange(sensor, yIndex, distance) {
				heightDiffToYLine := int(math.Abs(float64(yIndex) - float64(sensor.position.y)))
				minX := sensor.position.x - (distance - heightDiffToYLine) + 1
				maxX := sensor.position.x + (distance - heightDiffToYLine) + 1
				for _, p := range beaconRanges {
					if p.x <= minX && p.y >= maxX {
						continue OUTER
					}
				}
				beaconRanges = append(beaconRanges, point{x: minX, y: maxX})
			}
		}

		sort.Slice(beaconRanges, func(i, j int) bool {
			return beaconRanges[i].x < beaconRanges[j].x
		})
		currentX := 0
		for _, p := range beaconRanges {
			if currentX < p.x {
				p2 = currentX*maxY + yIndex
				fmt.Printf("Part 2 - %d\n", p2)
				return
			}
			if currentX < p.y {
				currentX = p.y
			}
		}

		// Part 1
		if yIndex == lineNo {
			currentX = beaconRanges[0].x
			for _, p := range beaconRanges {
				if currentX >= p.y {
					continue
				}
				if currentX > p.x {
					p1 += p.y - currentX
				} else {
					p1 += p.y - p.x
				}
				currentX = p.y
			}
			fmt.Printf("Part 1 - %d\n", p1)
		}
	}

}

type sensor struct {
	position point
	beacon   point
}

type point struct {
	x, y int
}

func inYRange(sensor sensor, yIndex int, distance int) bool {
	if sensor.position.y < yIndex {
		return sensor.position.y+distance > yIndex
	} else {
		return sensor.position.y-distance < yIndex
	}
}

func manhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p2.y)-float64(p1.y)) + math.Abs(float64(p2.x)-float64(p1.x)))
}

func parse(s string) sensor {
	re := regexp.MustCompile(".[0-9]+")
	matches := re.FindAllString(s, -1)
	for i, match := range matches {
		matches[i] = strings.TrimPrefix(match, "=")
	}
	ints := stringsToInts(matches)
	newSensor := sensor{
		position: point{
			x: ints[0],
			y: ints[1],
		},
		beacon: point{
			x: ints[2],
			y: ints[3],
		},
	}
	return newSensor
}

func stringsToInts(s []string) []int {
	i := make([]int, 0, len(s))
	for _, str := range s {
		n, _ := strconv.Atoi(str)
		i = append(i, n)
	}
	return i
}
