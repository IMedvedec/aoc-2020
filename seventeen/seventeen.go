package seventeen

import (
	"bufio"
	"log"
	"os"
)

const (
	filename    = "input/seventeen_input"
	cycleNumber = 6
	activeCube  = rune('#')
)

// Run is the solution starting point.
func Run() {
	cube := cubeInitialization()

	for i := 0; i < cycleNumber; i++ {
		cube.runIteration()
	}

	log.Printf("After %d iterations there are %d active cubes.\n", cycleNumber, cube.getActiveCount())
}

// cubeInitialization is the input parser. It prepares the initial cube states.
func cubeInitialization() *cube {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	scanner := bufio.NewScanner(file)

	cache := make(map[string]string)
	cube := newCube(cache)

	y, z := 0, 0
	for scanner.Scan() {
		line := scanner.Text()

		for x, state := range line {
			if state == activeCube {
				cube.setActive(x, y, z)
			}
		}

		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return cube
}
