package second

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	filename = "input/second_input"
)

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}
	defer file.Close()

	var validPasswordCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if analyzeLine(line) {
			validPasswordCount++
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	log.Printf("Valid password count: %d\n", validPasswordCount)
}

func analyzeLine(line string) bool {
	rule := strings.Fields(line)

	occurances := strings.Split(rule[0], "-")
	minOccurances, err := strconv.ParseInt(occurances[0], 10, 64)
	if err != nil {
		log.Fatalf("number parse error: %v", err)
	}
	maxOccurances, err := strconv.ParseInt(occurances[1], 10, 64)
	if err != nil {
		log.Fatalf("number parse error: %v", err)
	}

	// The assignment assumption is that we check the occurance of a single rune
	// in the given password string.
	var countedRune rune
	for _, r := range rule[1][:len(rule[1])-1] {
		countedRune = r
		break
	}

	var foundOccurances int64 = 0
	for _, rune := range rule[2] {
		if rune == countedRune {
			foundOccurances++
		}
	}

	return foundOccurances >= minOccurances && foundOccurances <= maxOccurances
}
