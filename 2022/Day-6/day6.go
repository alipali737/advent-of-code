package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to parse input file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	packetQueue := make([]string, 0)
	foundPacket := false
	messageQueue := make([]string, 0)
	foundMessage := false
	counter := 0
	for fileScanner.Scan() {
		readText := fileScanner.Text()
		for _, letter := range strings.Split(readText, "") {
			counter++
			if !foundPacket {
				packetQueue = append(packetQueue, letter)
				if len(packetQueue) == 4 {
					if compareQueue(packetQueue) {
						fmt.Printf("Part 1 - Found 4 unqiue letters as start-of-packet: %s, counter: %d\n", packetQueue, counter)
						foundPacket = true
					}
					packetQueue = packetQueue[1:]
				}
			}
			if !foundMessage {
				messageQueue = append(messageQueue, letter)
				if len(messageQueue) == 14 {
					if compareQueue(messageQueue) {
						fmt.Printf("Part 2 - Found 14 unqiue letters as start-of-message: %s, counter: %d\n", messageQueue, counter)
						foundMessage = true
					}
					messageQueue = messageQueue[1:]
				}
			}

			if foundMessage && foundPacket {
				break
			}
			// fmt.Printf("%d : [%s] -> %s\n", counter, letter, queue)
		}
	}
}

func compareQueue(q []string) bool {
	unique := make(map[string]bool, len(q))
	uq := make([]string, len(unique))
	for _, elem := range q {
		if !unique[elem] {
			uq = append(uq, elem)
			unique[elem] = true
		}
	}

	return len(uq) == len(q)
}
