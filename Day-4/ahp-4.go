package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SplitLine(line string) []string {
	slice := regexp.MustCompile(`[ ]+`).Split(line, -1)
	if slice[0] == "" {
		slice = slice[1:]
	}
	// fmt.Printf("Slice after regex: %s\n", slice)
	return slice
}

func ImportFile(path string) (calledNums []string, boards [][][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum, currentBoard, currentBoardY := 0, 0, 0
	var board [][]string
	for scanner.Scan() {
		// If first line get the list of called Nums
		if lineNum == 0 {
			calledNums = strings.Split(scanner.Text(), ",")
		}

		// If rest of boards
		if scanner.Text() != "" && lineNum != 0 {
			// TODO : Read boards using 'currentBoard' & 'currentBoardY' & 'board'
			// fmt.Printf("currentBoardY = %d : %s\n", currentBoardY, scanner.Text())
			board = append(board, SplitLine(scanner.Text()))
			currentBoardY++
			if currentBoardY == 5 {
				currentBoardY = 0
				boards = append(boards, board)
				board = make([][]string, 0)
				currentBoard++
			}
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return
}

func CheckWin(board [][]string) (win bool) {
	var marked int
	var axis string = "hori"

	// Horizontal check boundries
	firstBoundry := len(board)
	secondBoundry := len(board[0])

	// Do it twice to cover both axis
	for i := 0; i < 2; i++ {
		for x := 0; x < firstBoundry; x++ {
			marked = 0
			for y := 0; y < secondBoundry; y++ {
				if axis == "hori" {
					if board[x][y] == "X" {
						marked++
					}
				} else {
					if board[y][x] == "X" {
						marked++
					}
				}
			}
			if marked == 5 {
				fmt.Printf("Winner Board: %s\n", board)
				return true
			}
		}
		// Invert vars to check vertically next
		axis = "vert"
	}

	return false
}

func CallNum2(n string, boards [][][]string) (winningBoardIndex []int, newBoards [][][]string) {
	var tmp bool = false
	for i, board := range boards {
		for j, x := range board {
			for y := range x {
				if x[y] == n {
					boards[i][j][y] = "X"
					win := CheckWin(board)
					if win {
						tmp = true
						winningBoardIndex = append(winningBoardIndex, i)
					}
				}
			}
		}
	}
	if tmp {
		newBoards = boards
		return
	}
	return nil, boards
}

func CallNum(n string, boards [][][]string) (win bool, winningBoardIndex int, newBoards [][][]string) {
	for i, board := range boards {
		for j, x := range board {
			for y := range x {
				if x[y] == n {
					boards[i][j][y] = "X"
					win = CheckWin(board)
					if win {
						winningBoardIndex = i
						newBoards = boards
						return
					}
				}
			}
		}
	}
	return false, -1, boards
}

func RemoveBoardArr(boards [][][]string, a []int) [][][]string {
	var byteBoards, byteA []byte
	boardsCopy := boards
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(boards); j++ {
			match := 0
			for rowIndex := range boards[j] {
				byteBoards = []byte(strings.Join(boards[j][rowIndex], ","))
				byteA = []byte(strings.Join(boardsCopy[a[i]][rowIndex], ","))
				if bytes.Equal(byteBoards, byteA) {
					match++
				}
			}
			if match == 5 {
				boards = RemoveBoard(boards, j)
			}
		}
	}
	return boards
}

func RemoveBoard(boards [][][]string, i int) [][][]string {
	boards[i] = boards[len(boards)-1]
	return boards[:len(boards)-1]
}

func Step2() {
	var calledNums []string
	var boards [][][]string
	var lastCalled string
	var counter int
	var winningBoardIndex []int
	var finalBoard [][]string
	calledNums, boards = ImportFile("input.txt")

	for len(boards) > 1 {
		winningBoardIndex, boards = CallNum2(calledNums[counter], boards)
		if len(winningBoardIndex) >= 1 {
			lastCalled = calledNums[counter]
			finalBoard = boards[winningBoardIndex[len(winningBoardIndex)-1]]
			boards = RemoveBoardArr(boards, winningBoardIndex)
		}
		if counter >= len(calledNums)-1 {
			fmt.Printf("Last Number: %s - Board: %s\n", lastCalled, finalBoard)
			break
		} else {
			counter++
		}
	}
	// fmt.Printf("Winning board: %s\n", boards[winningBoardIndex])

	tot := 0
	for x := 0; x < len(boards[0]); x++ {
		for y := 0; y < len(boards[0][0]); y++ {
			if boards[0][x][y] != "X" {
				tmp, _ := strconv.Atoi(boards[0][x][y])
				tot += tmp
			}
		}
	}
	tmp, _ := strconv.Atoi(calledNums[counter-1])
	tot *= tmp

	fmt.Printf("Result: %d, Last Called: %d, Num of Boards left: %d", tot, tmp, len(boards))
}

func Step1() {
	var calledNums []string
	var boards [][][]string
	var win bool
	var counter, winningBoardIndex int
	calledNums, boards = ImportFile("input.txt")

	for !win {
		win, winningBoardIndex, boards = CallNum(calledNums[counter], boards)
		counter++
		if counter >= len(calledNums) {
			log.Panicln("No winner found!")
		}
	}
	// fmt.Printf("Winning board: %s\n", boards[winningBoardIndex])

	tot := 0
	for x := 0; x < len(boards[winningBoardIndex]); x++ {
		for y := 0; y < len(boards[winningBoardIndex][0]); y++ {
			if boards[winningBoardIndex][x][y] != "X" {
				tmp, _ := strconv.Atoi(boards[winningBoardIndex][x][y])
				tot += tmp
			}
		}
	}
	tmp, _ := strconv.Atoi(calledNums[counter-1])
	tot *= tmp

	fmt.Printf("Result: %d\n", tot)
}

func main() {
	Step1()
	Step2()
}
