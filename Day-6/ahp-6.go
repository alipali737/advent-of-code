package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ImportFile(path string) (fish []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := "[" + scanner.Text() + "]"
		err := json.Unmarshal([]byte(str), &fish)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning: %v\n", err)
	}
	return fish
}

func main() {
	days := 256
	fish := ImportFile("input.txt")
	// TODO : Divide up the list and use goroutines to iterate and change that part of the list then return it and rebuild list
	// LINK : http://www.golangpatterns.info/concurrency/parallel-for-loop : https://www.golang-book.com/books/intro/10
	for i := 0; i < days; i++ {
		fishToAdd := 0
		for j, v := range fish {
			if v == 0 {
				fishToAdd++
				fish[j] = 6
			} else {
				fish[j]--
			}
		}

		for i := 0; i < fishToAdd; i++ {
			fish = append(fish, 8)
		}
	}
	fmt.Printf("Number of fish on day %d: %d", days, len(fish))
}
