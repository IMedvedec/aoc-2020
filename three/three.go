package three

import (
	"bufio"
	"log"
	"os"
)

const (
	filename    = "input/three_input"
	coordsRight = 3
	coordsDown  = 1
	empty       = '.'
	tree        = '#'
)

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	x, y := 0, 0
	var treeHitCount int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if isTree(line, x) {
			treeHitCount++
		}
		x += coordsRight
		y += coordsDown
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	log.Printf("Tree hit count: %d\n", treeHitCount)
}

func isTree(line string, x int) bool {
	patternLength := len(line)
	if patternLength == 0 {
		return false
	}

	index := func() int {
		if x < patternLength {
			return x
		}

		multiplicator := x / patternLength
		return x - (multiplicator * patternLength)
	}()

	return line[index] == tree
}
