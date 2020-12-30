package seventeen

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
)

const (
	filename    = "input/seventeen_input"
	cycleNumber = 6
	activeCube  = rune('#')
)

// cube is the problem data structure.
type cube struct {
	fields map[string]struct{}
	cache  map[string]string

	// dimension ranges
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

func newCube(cache map[string]string) *cube {
	return &cube{
		fields: make(map[string]struct{}),
		cache:  cache,
	}
}

func (c *cube) runIteration() {
	newCube := newCube(c.cache)

	for x := c.minX - 1; x <= c.maxX+1; x++ {
		for y := c.minY - 1; y <= c.maxY+1; y++ {
			for z := c.minZ - 1; z <= c.maxZ+1; z++ {
				c.newIterationState(newCube, x, y, z, c.isActive(x, y, z))
			}
		}
	}

	c.fields = newCube.fields
	c.minX, c.maxX = newCube.minX, newCube.maxX
	c.minY, c.maxY = newCube.minY, newCube.maxY
	c.minZ, c.maxZ = newCube.minZ, newCube.maxZ
}

func (c *cube) newIterationState(newCube *cube, x, y, z int, state bool) {
	var activeNeighborCount int

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if c.isActive(i, j, k) && !(x == i && y == j && z == k) {
					activeNeighborCount++
				}
			}
		}
	}

	if state {
		if !(activeNeighborCount == 2 || activeNeighborCount == 3) {
			newCube.setInactive(x, y, z)
		} else {
			newCube.setActive(x, y, z)
		}
	} else {
		if activeNeighborCount == 3 {
			newCube.setActive(x, y, z)
		} else {
			newCube.setInactive(x, y, z)
		}
	}
}

func (c *cube) setActive(x, y, z int) {
	key := c.getKey(x, y, z)

	c.fields[key] = struct{}{}

	c.updateDimensions(x, y, z)
}

func (c *cube) setInactive(x, y, z int) {
	key := c.getKey(x, y, z)

	delete(c.fields, key)
}

func (c *cube) isActive(x, y, z int) bool {
	key := c.getKey(x, y, z)

	_, isActive := c.fields[key]
	return isActive
}

func (c *cube) getActiveCount() int {
	return len(c.fields)
}

func (c *cube) getKey(x, y, z int) string {
	hashInput := fmt.Sprintf("%d-%d-%d", x, y, z)

	key, ok := c.cache[hashInput]
	if !ok {
		hasher := sha1.New()
		hasher.Write([]byte(hashInput))
		key = string(hasher.Sum(nil))

		c.cache[hashInput] = key
	}

	return key
}

func (c *cube) updateDimensions(x, y, z int) {
	if x > c.maxX {
		c.maxX = x
	}
	if x < c.minX {
		c.minX = x
	}

	if y > c.maxY {
		c.maxY = y
	}
	if y < c.minY {
		c.minY = y
	}

	if z > c.maxZ {
		c.maxZ = z
	}
	if z < c.minZ {
		c.minZ = z
	}
}

func Run() {
	cube := cubeInitialization()

	for i := 0; i < cycleNumber; i++ {
		cube.runIteration()
	}

	log.Printf("After %d iterations there are %d active cubes.\n", cycleNumber, cube.getActiveCount())
}

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
