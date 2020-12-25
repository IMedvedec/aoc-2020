package first

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	filename                 = "input/first_input"
	twoNumberSummation int64 = 2020
)

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}
	defer file.Close()

	inputMap := make(map[int64]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("number parse error: %v", err)
		}

		inputMap[number] = struct{}{}

		numberDifference := twoNumberSummation - number
		if _, ok := inputMap[numberDifference]; ok {
			log.Printf("Numbers found: %d, %d\n\t%d + %d = %d\n\t%d * %d = %d\n", number, numberDifference,
				number, numberDifference, number+numberDifference, number, numberDifference, number*numberDifference)
			return
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}
}
