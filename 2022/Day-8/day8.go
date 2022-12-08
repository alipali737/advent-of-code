package main

import (
	"bufio"
	"fmt"
	"os"
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

	visible := 0
	trees := [][]int{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		split := strings.Split(readText, "")
		ints := make([]int, len(split))
		for i, tree := range split {
			ints[i], _ = strconv.Atoi(tree)
		}
		trees = append(trees, ints)
	}

	countedTrees := map[string]bool{}

	calcVisibleFromOutsideDirection("left", trees, countedTrees)
	calcVisibleFromOutsideDirection("right", trees, countedTrees)

	calcVisibleFromOutsideDirection("top", trees, countedTrees)
	calcVisibleFromOutsideDirection("bottom", trees, countedTrees)
	visible = len(countedTrees)

	// fmt.Println(countedTrees)

	visible += len(trees)*2 + len(trees[0])*2 - 4
	fmt.Printf("Visible: %d\n", visible)

	highestScenicScore := 0
	for key := range countedTrees {
		// fmt.Printf("\n%s\n", key)
		split := strings.Split(key, ",")
		row_index, _ := strconv.Atoi(split[0])
		column_index, _ := strconv.Atoi(split[1])
		score := calcScenicScore(row_index, column_index, trees)
		if score > highestScenicScore {
			highestScenicScore = score
		}
		// fmt.Printf("Score: %d, Highest: %d\n", score, highestScenicScore)
	}

	fmt.Printf("Highest scenic score: %d\n", highestScenicScore)

}

func calcVisibleFromOutsideDirection(directionSide string, trees [][]int, countedTrees map[string]bool) {

	switch directionSide {
	case "left":
		for row_index := 1; row_index < len(trees)-1; row_index++ {
			max := trees[row_index][0]
			for column_index := 1; column_index < len(trees[row_index])-1; column_index++ {
				tree := trees[row_index][column_index]

				if tree > max {
					max = tree
					countedTrees[fmt.Sprintf("%d,%d", row_index, column_index)] = true
					// fmt.Printf("Left - [%d,%d] : %d : %d\n", row_index, column_index, trees[row_index], tree)
				}
			}
		}
	case "right":
		for row_index := 1; row_index < len(trees)-1; row_index++ {
			max := trees[row_index][len(trees[row_index])-1]
			for column_index := len(trees[row_index]) - 2; column_index >= 1; column_index-- {
				tree := trees[row_index][column_index]

				if tree > max {
					max = tree
					countedTrees[fmt.Sprintf("%d,%d", row_index, column_index)] = true
					// fmt.Printf("Right - [%d,%d] : %d : %d\n", row_index, column_index, trees[row_index], tree)
				}
			}
		}
	case "top":
		for column_index := 1; column_index < len(trees[0])-1; column_index++ {
			column := getColumn(trees, column_index)
			max := column[0]
			for row_index := 1; row_index < len(column)-1; row_index++ {
				tree := column[row_index]
				if tree > max {
					max = tree
					countedTrees[fmt.Sprintf("%d,%d", row_index, column_index)] = true
					// fmt.Printf("Top - [%d,%d] : %d : %d\n", row_index, column_index, column, tree)
				}
			}
		}
	case "bottom":
		for column_index := 1; column_index < len(trees[0])-1; column_index++ {
			column := getColumn(trees, column_index)
			max := column[len(column)-1]
			for row_index := len(column) - 2; row_index > 0; row_index-- {
				tree := column[row_index]
				if tree > max {
					max = tree
					countedTrees[fmt.Sprintf("%d,%d", row_index, column_index)] = true
					// fmt.Printf("Bottom - [%d,%d] : %d : %d\n", row_index, column_index, column, tree)
				}
			}
		}
	}
}

func getColumn(s [][]int, columnIndex int) []int {
	column := make([]int, 0)
	for _, row := range s {
		column = append(column, row[columnIndex])
	}
	return column
}

func calcScenicScore(row_index, column_index int, trees [][]int) int {
	score := 1
	score *= calcVisibleFromInsideDirection("top", trees, row_index, column_index)
	score *= calcVisibleFromInsideDirection("left", trees, row_index, column_index)
	score *= calcVisibleFromInsideDirection("right", trees, row_index, column_index)
	score *= calcVisibleFromInsideDirection("bottom", trees, row_index, column_index)
	return score
}

func calcVisibleFromInsideDirection(directionSide string, trees [][]int, startRowIndex, startColumnIndex int) int {
	visible := 0

	switch directionSide {
	case "right":
		max := trees[startRowIndex][startColumnIndex]
		for column_index := startColumnIndex + 1; column_index < len(trees[startRowIndex]); column_index++ {
			tree := trees[startRowIndex][column_index]

			visible++
			if tree >= max {
				// fmt.Printf("Right - [%d,%d] : %d : %d - visible: %d\n", startRowIndex, column_index, trees[startRowIndex], tree, visible)
				return visible
			}
		}
		// fmt.Printf("Right - reached edge, visible: %d\n", visible)

	case "left":
		max := trees[startRowIndex][startColumnIndex]
		for column_index := startColumnIndex - 1; column_index >= 0; column_index-- {
			tree := trees[startRowIndex][column_index]

			visible++
			if tree >= max {
				// fmt.Printf("Left - [%d,%d] : %d : %d - visible: %d\n", startRowIndex, column_index, trees[startRowIndex], tree, visible)
				return visible
			}
		}
		// fmt.Printf("Left - reached edge, visible: %d\n", visible)

	case "top":
		column := getColumn(trees, startColumnIndex)
		max := column[startRowIndex]
		for row_index := startRowIndex - 1; row_index >= 0; row_index-- {
			tree := column[row_index]

			visible++
			if tree >= max {
				// fmt.Printf("Top - [%d,%d] : %d : %d - visible: %d\n", row_index, startColumnIndex, column, tree, visible)
				return visible
			}
		}
		// fmt.Printf("Top - reached edge, visible: %d\n", visible)

	case "bottom":
		column := getColumn(trees, startColumnIndex)
		max := column[startRowIndex]
		for row_index := startRowIndex + 1; row_index < len(column); row_index++ {
			tree := column[row_index]

			visible++
			if tree >= max {
				// fmt.Printf("Bottom - [%d,%d] : %d : %d - visible: %d\n", row_index, startColumnIndex, column, tree, visible)
				return visible
			}
		}
		// fmt.Printf("Bottom - reached edge, visible: %d\n", visible)
	}
	return visible
}
