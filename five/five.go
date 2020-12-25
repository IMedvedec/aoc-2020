package five

import (
	"bufio"
	"log"
	"os"
)

const (
	filename          = "input/five_input"
	planeRowNumber    = 128
	planeColumnNumber = 8
)

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var highestSeatID int

	for scanner.Scan() {
		line := scanner.Text()
		seatID := calculateSeatID(line)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	log.Printf("Highest seat id: %d", highestSeatID)
}

func calculateSeatID(line string) int {
	rowPattern, columnPattern := line[:7], line[7:]

	row := calculatePattern(rowPattern, 0, 127)
	column := calculatePattern(columnPattern, 0, 7)

	return row*planeColumnNumber + column
}

func calculatePattern(pattern string, min, max int) int {
	isLower := func(char byte) bool {
		return char == 'F' || char == 'L'
	}

	var result int
	for i := range pattern {
		if isLower(pattern[i]) {
			max = min + ((max - min) / 2)
			result = min
		} else {
			min = max - ((max - min) / 2)
			result = max
		}
	}

	return result
}
