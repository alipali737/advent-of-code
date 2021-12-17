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

type Coord struct {
	X, Y int
}

func StringSlice2IntSlice(_strings []string) []int {
	ints := make([]int, len(_strings))
	for i, s := range _strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func DupCount(list []Coord) map[Coord]int {
	duplicate_frequency := make(map[Coord]int)

	for _, item := range list {
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1
		} else {
			duplicate_frequency[item] = 1
		}
	}
	return duplicate_frequency
}

func GetDiagonalPoints(p1 Coord, p2 Coord) (points []Coord) {
	// BB vs SS
	if p1.X > p2.X && p1.Y > p2.Y {
		for i := 0; i <= p1.X-p2.X; i++ {
			points = append(points, Coord{p2.X + i, p2.Y + i})
		}
	}

	// SS vs BB
	if p1.X < p2.X && p1.Y < p2.Y {
		for i := 0; i <= p2.X-p1.X; i++ {
			points = append(points, Coord{p1.X + i, p1.Y + i})
		}
	}

	// BS vs SB
	if p1.X > p2.X && p1.Y < p2.Y {
		for i := 0; i <= p1.X-p2.X; i++ {
			points = append(points, Coord{p1.X - i, p1.Y + i})
		}
	}

	// SB vs BS
	if p1.X < p2.X && p1.Y > p2.Y {
		for i := 0; i <= p2.X-p1.X; i++ {
			points = append(points, Coord{p1.X + i, p1.Y - i})
		}
	}

	return points
}

func GetHoriVertPoints(p1 Coord, p2 Coord) (points []Coord) {

	// X is the same, iterate Y
	if p1.X == p2.X {
		big := Max(p1.Y, p2.Y)
		small := Min(p1.Y, p2.Y)
		for i := 0; i <= big-small; i++ {
			points = append(points, Coord{p1.X, small + i})
		}
	} else if p1.Y == p2.Y {
		// Y is the same, iterate X
		big := Max(p1.X, p2.X)
		small := Min(p1.X, p2.X)
		for i := 0; i <= big-small; i++ {
			points = append(points, Coord{small + i, p1.Y})
		}
	}

	return points
}

func ImportFile(path string) (points []Coord) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c1, c2 Coord
		line := StringSlice2IntSlice(strings.Split(regexp.MustCompile(` -> `).ReplaceAllString(scanner.Text(), ","), ","))
		c1 = Coord{line[0], line[1]}
		c2 = Coord{line[2], line[3]}
		if c1.X == c2.X || c1.Y == c2.Y {
			// fmt.Printf("Considered line: %d, %d\n", c1, c2)
			points = append(points, GetHoriVertPoints(c1, c2)...)
		} else {
			points = append(points, GetDiagonalPoints(c1, c2)...)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return points
}

func PrintBoard(points []Coord, mapX int, mapY int) {
	board := make([][]int, mapY)
	for i := range board {
		board[i] = make([]int, mapX)
	}
	for _, v := range points {
		board[v.Y][v.X]++
	}

	for _, row := range board {
		stringRow := arrayToString(row, ",")
		stringRow = regexp.MustCompile(",").ReplaceAllString(stringRow, " ")
		stringRow = regexp.MustCompile("0").ReplaceAllString(stringRow, ".")
		fmt.Println(stringRow)
	}
}

func main() {
	points := ImportFile("input.txt")
	dup_map := DupCount(points)
	var counter int
	for _, v := range dup_map {
		if v > 1 {
			counter++
		}
	}
	fmt.Printf("Number of overlaps: %d\n", counter)
}
