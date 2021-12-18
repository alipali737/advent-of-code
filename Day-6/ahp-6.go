package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
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

func ChunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func Collapse2DArray(slice [][]int) []int {
	var newSlice []int
	for _, v := range slice {
		for _, n := range v {
			newSlice = append(newSlice, n)
		}
	}
	return newSlice
}

func main() {
	days := 256
	divisions := 500000.0
	fish := ImportFile("example.txt")

	// TODO : Divide up the list and use goroutines to iterate and change that part of the list then return it and rebuild list
	// LINK : http://www.golangpatterns.info/concurrency/parallel-for-loop : https://www.golang-book.com/books/intro/10

	/*
		    Start: 3,4,3,1,2
		End Day 0: 2,3,2,0,1
		End Day 1: 1,2,1,6,8,0
		End Day 2: 0,1,0,5,7,6,8
	*/

	for j := 0; j < days; j++ {

		fmt.Printf("-------------- DAY %d --------------\n", j)
		// fmt.Printf("Divisions = %f, chunkSize = %d", divisions, int(math.Ceil(float64(len(fish))/divisions)))
		chunks := ChunkSlice(fish, int(math.Ceil(float64(len(fish))/divisions)))
		// fmt.Printf("Before Iteration: %d\n", chunks)

		c := make(chan string, len(chunks))
		fishToAdd := make([]int, len(chunks))

		for i := range chunks {

			go func(i int, k []int) {
				fishToAdd[i] = 0
				for j, v := range k {
					// fmt.Printf("Evaluating j = %d, v = %d, chunk = %d, chunkIndex = %d\n", j, v, k, i)
					if v == 0 {
						// fmt.Printf("==> Creating new fish on j = %d, v = %d, chunk = %d, chunkIndex = %d\n", j, v, k, i)
						fishToAdd[i]++
						k[j] = 6
					} else {
						k[j]--
					}
					// fmt.Printf("Chunk %d after loop = %d\n", i, k)
				}

				// fmt.Printf("Fish to add: %d, for chunk %d\n", fishToAdd[i], i)
				c <- ""
				chunks[i] = k
			}(i, chunks[i])

		}
		// Wait for all to finish
		for i := 0; i < len(chunks); i++ {
			<-c
		}

		// fmt.Printf("Fish To Add = %d\n", fishToAdd)

		for x := 0; x < len(fishToAdd); x++ {
			for f := 0; f < fishToAdd[x]; f++ {
				var tmp []int
				for _, v := range chunks[x] {
					tmp = append(tmp, v)
				}
				tmp = append(tmp, 8)
				chunks[x] = tmp
			}
		}

		// fmt.Printf("After Iteration: %d\n", chunks)
		fish = Collapse2DArray(chunks)

	}
	fmt.Printf("-------------- END --------------\n")
	fmt.Printf("Number of fish after %d days: %d\n", days, len(fish))
}
