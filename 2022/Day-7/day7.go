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

	dirs := map[string][]string{}
	currentPath := ""
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		readText = strings.TrimPrefix(readText, "$ ")
		command := strings.Split(readText, " ")
		switch command[0] {
		case "cd":
			if command[1] == "/" {
				command[1] = "~"
			}
			if command[1] == ".." {
				tmp := strings.Split(currentPath, "/")
				currentPath = strings.Join(tmp[:len(tmp)-1], "/")
				continue
			}
			if len(currentPath) > 0 {
				currentPath += "/" + command[1]
			} else {
				currentPath = command[1]
			}
			if _, ok := dirs[currentPath]; !ok {
				dirs[currentPath] = []string{}
			}

		case "ls":
			continue

		default:
			dirs[currentPath] = append(dirs[currentPath], strings.Join(command, " "))
		}
	}

	paths := []string{}
	for path := range dirs {
		paths = append(paths, path)
	}
	pathLengths := map[string]int{}
	for _, path := range paths {
		pathLengths[path] = strings.Count(path, "/")
	}

	paths = sortMapByValues(pathLengths, true)

	dirSizes := map[string]int{}

	for _, path := range paths {
		size := 0
		for _, item := range dirs[path] {
			split := strings.Split(item, " ")
			switch split[0] {
			case "dir":
				size += dirSizes[path+"/"+split[1]]
			default:
				v, _ := strconv.Atoi(split[0])
				size += v
			}
		}
		dirSizes[path] = size
	}

	p1Answer := 0
	for _, size := range dirSizes {
		if size <= 100000 {
			p1Answer += size
		}
	}

	fmt.Printf("Part 1 - Total directory sizes: %d\n", p1Answer)

	sizeLeft := 70000000 - dirSizes["~"]
	sizeRequired := 30000000 - sizeLeft

	dirsBySize := sortMapByValues(dirSizes, false)
	p2Answer := 0
	for _, dir := range dirsBySize {
		if dirSizes[dir] >= sizeRequired {
			p2Answer = dirSizes[dir]
			break
		}
	}

	fmt.Printf("Part 2 - Directory size to delete: %d\n", p2Answer)

}

func sortMapByValues(m map[string]int, desc bool) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	if desc {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] > m[keys[j]]
		})
	} else {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] < m[keys[j]]
		})
	}

	return keys
}
