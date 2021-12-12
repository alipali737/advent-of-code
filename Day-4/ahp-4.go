package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringSlice2IntSlice(x []string) (y []int) {
	for _, i := range x {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		y = append(y, j)
	}
	return
}

func ImportFile(path string) (calledNums []int, boards [][][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum, currentBoard, currentBoardY := 0, 0, 0
	board := make([][]int, 0)
	for scanner.Scan() {
		// If first line get the list of called Nums
		if lineNum == 0 {
			calledNums = StringSlice2IntSlice(strings.Split(scanner.Text(), ","))
		}

		// If rest of boards
		if scanner.Text() != "\n" {
			// TODO : Read boards using 'currentBoard' & 'currentBoardY' & 'board'
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return
}

func Step1() {
	var calledNums []int
	var boards [][][]int
	calledNums, boards = ImportFile("input.txt")
	fmt.Println(calledNums)
	fmt.Println(boards)
}

func main() {
	Step1()
}
