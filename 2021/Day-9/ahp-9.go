package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func strSlice2Int(slice []string) []int {
	tmp := []int{}
	for _, v := range slice {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, j)
	}
	return tmp
}

func int2DSlice2Str(slice [][]int) (strSlice [][]string) {
	for _, row := range slice {
		var tmp []string
		for _, v := range row {
			tmp = append(tmp, string(v))
		}
		strSlice = append(strSlice, tmp)
	}
	return strSlice
}

func ImportFile(path string) (heightmap [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "")
		heightmap = append(heightmap, strSlice2Int(str))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return heightmap
}

func getSurrounding(heightmap [][]int, x int, y int) (points []int) {
	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // Up, left, down, right
	mapBoundX := len(heightmap[0])
	mapBoundY := len(heightmap)

skip:
	for _, direction := range directions {
		dx := x + direction[0]
		dy := y + direction[1]
		if dx < 0 || dx >= mapBoundX || dy < 0 || dy >= mapBoundY {
			continue skip
		}

		points = append(points, heightmap[dy][dx])
	}
	return points
}

func main() {
	heightmap := ImportFile("input.txt")
	// fmt.Println(heightmap)
	step1(heightmap)
	step2(heightmap)
}

func calcLowestPoints(heightmap [][]int) []int {
	var lowestPoints []int
	for y, row := range heightmap {
		for x, point := range row {
			var lowest bool = true
			adjacent := getSurrounding(heightmap, x, y)
			for _, v := range adjacent {
				if v <= point {
					lowest = false
				}
			}
			if lowest {
				lowestPoints = append(lowestPoints, point)
				// fmt.Printf("Found lowest point of value %d, at coordinates [%d,%d]\n", point, x, y)
			}
		}
	}
	return lowestPoints
}

func floodFill(heightmapStr [][]string, x int, y int, boundryHeight int) int {
	boundryHeightStr := string(boundryHeight)

	if x < 0 || x >= len(heightmapStr[0]) || y < 0 || y >= len(heightmapStr) { // Index out of range
		return 0
	}

	if heightmapStr[y][x] == boundryHeightStr || heightmapStr[y][x] == "." { // Make sure its not a boundry or we have been there before
		return 0
	}

	heightmapStr[y][x] = "." // Mark the point as we have checked it now

	floodFill(heightmapStr, x+1, y, 9) // Go right
	floodFill(heightmapStr, x-1, y, 9) // Go left
	floodFill(heightmapStr, x, y-1, 9) // Go up
	floodFill(heightmapStr, x, y+1, 9) // Go down

	var counter int
	for _, row := range heightmapStr {
		for _, v := range row {
			if v == "." {
				counter++
			}
		}
	}

	// fmt.Println(heightmapStr)
	return counter

}

func max(slice []int) (int, int) {
	largest := slice[0]
	largestIndex := 0
	for k, v := range slice {
		if v > largest {
			largest = v
			largestIndex = k
		}
	}
	return largest, largestIndex
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func getLargest3Values(slice []int) []int {
	var tmp2 []int
	for i := 0; i < 3; i++ {
		largest, k := max(slice)
		tmp2 = append(tmp2, largest)
		slice = remove(slice, k)
	}
	return tmp2
}

func step1(heightmap [][]int) {

	lowestPoints := calcLowestPoints(heightmap)
	// fmt.Printf("Lowest Point are: %#v\n", lowestPoints)

	var totalRisk int = 0
	for _, v := range lowestPoints {
		totalRisk += v + 1
	}

	fmt.Printf("Total Risk of all low points are is: %d\n", totalRisk)
}

func step2(heightmap [][]int) {
	var lowestPoints [][]int
	for y, row := range heightmap {
		for x, point := range row {
			var lowest bool = true
			adjacent := getSurrounding(heightmap, x, y)
			for _, v := range adjacent {
				if v <= point {
					lowest = false
				}
			}
			if lowest {
				var tmp []int
				tmp = append(tmp, x)
				tmp = append(tmp, y)
				lowestPoints = append(lowestPoints, tmp)
				// fmt.Printf("Found lowest point of value %d, at coordinates [%d,%d]\n", point, x, y)
			}
		}
	}

	var basinSizes []int
	for _, points := range lowestPoints {
		basinSizes = append(basinSizes, floodFill(int2DSlice2Str(heightmap), points[0], points[1], 9))
	}
	fmt.Printf("Basin Sizes = %#v\n", basinSizes)
	l := getLargest3Values(basinSizes)
	total := l[0] * l[1] * l[2]
	fmt.Printf("Step 2 Total : %d\n", total)

}
