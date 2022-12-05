package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	stackCharacters := [][]string{}
	stacksSectionComplete := false
	stackStrings := []string{}
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		// fmt.Printf("Current Line: \"%s\", stacksSectionComplete = %t\n", readText, stacksSectionComplete)
		if readText == "" {
			stacksSectionComplete = true
			stackCharacters = stripStacks(stackCharacters)
			// print2DSlice(stackCharacters)
			// fmt.Println("---------")
			stackCharacters = rotateMatrix(stackCharacters)
			// print2DSlice(stackCharacters)
			for _, row := range stackCharacters {
				stack := ""
				stack += strings.Join(row, "")
				stackStrings = append(stackStrings, strings.ReplaceAll(stack, "-", ""))
			}
			fmt.Println(stackStrings)
			continue
		}

		if !stacksSectionComplete {
			stackCharacters = append(stackCharacters, strings.Split(readText, ""))
			continue
		}

		re := regexp.MustCompile(`\d[\d,]*`)
		strcmd := re.FindAllString(readText, -1)
		cmd := []int{}
		for _, v := range strcmd {
			tmp, _ := strconv.Atoi(v)
			cmd = append(cmd, tmp)
		}

		stackStrings = moveCrates(stackStrings, cmd[0], cmd[1], cmd[2])
		// fmt.Println("----")

	}

	finalTextP1 := ""
	for _, row := range stackStrings {
		split := strings.Split(row, "")
		finalTextP1 += split[len(split)-1]
	}
	fmt.Printf("Part 1 - Final text: %s\n", finalTextP1)

}

func stripStacks(s [][]string) [][]string {
	newSlice := [][]string{}

	for _, line := range s {
		lineText := strings.Join(line, "")
		lineText = strings.ReplaceAll(lineText, "[", "-")
		lineText = strings.ReplaceAll(lineText, "]", "-")
		lineText = strings.ReplaceAll(lineText, " ", "-")
		lineText = lineText[1 : len(lineText)-1]
		lineText = strings.ReplaceAll(lineText, "---", "")

		// fmt.Println(lineText)

		newSlice = append(newSlice, strings.Split(lineText, ""))
	}

	newSlice[0] = strings.Split("---"+strings.Join(newSlice[0], ""), "")
	// newSlice = newSlice[:len(newSlice)-1]

	return newSlice
}

func rotateMatrix(s [][]string) [][]string {
	size := len(s)
	layer_count := size / 2
	for layer := 0; layer < layer_count; layer++ {
		first := layer
		last := size - first - 1
		// fmt.Printf("Layer %d: first: %d, last: %d\n", layer, first, last)

		for element := first; element < last; element++ {

			offset := element - first

			top := s[first][element]
			right_side := s[element][last]
			bottom := s[last][last-offset]
			left_side := s[last-offset][first]

			s[first][element] = left_side
			s[element][last] = top
			s[last][last-offset] = right_side
			s[last-offset][first] = bottom

		}
	}

	return s
}

func print2DSlice(s [][]string) {
	for _, row := range s {
		for _, column := range row {
			fmt.Printf("%s", column)
		}
		fmt.Printf("\n")
	}
}

func moveCrates(s []string, amount int, src_index int, dst_index int) []string {
	src_index--
	dst_index--

	src := strings.Split(s[src_index], "")
	dst := strings.Split(s[dst_index], "")

	dst = append(dst, src[len(src)-amount:]...)
	src = src[:len(src)-amount]

	// fmt.Printf("%d : %s -[%s]-> %s\n", i+1, src, dst[len(dst)-1], dst)

	s[src_index] = strings.Join(src, "")
	s[dst_index] = strings.Join(dst, "")
	return s
}
