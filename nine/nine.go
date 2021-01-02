package nine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	filename      = "input/nine_input"
	preambleCount = 25
)

// Run is the solution starting point.
func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	scanner := bufio.NewScanner(file)

	mostRecentNumbers := make([]int64, 25, 25)
	var numberCounter int

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("number parse error: %v", err)
		}

		if numberCounter >= preambleCount {
			if !checkSummation(number, mostRecentNumbers) {
				log.Printf("Number %d does not follow the XMAS rules.\n", number)
				return
			}
		}

		mostRecentNumbers[numberCounter%preambleCount] = number

		numberCounter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	log.Println("Program is finished.")
}

// checkSummation is a helper function for checking XMAS number encoding rules.
func checkSummation(number int64, mostRecentNumbers []int64) bool {
	length := len(mostRecentNumbers)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j && (mostRecentNumbers[i]+mostRecentNumbers[j] == number) {
				return true
			}
		}
	}

	return false
}
