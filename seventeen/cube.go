package seventeen

import (
	"crypto/sha1"
	"fmt"
)

// cube is the problem data structure.
type cube struct {
	// fields contains the active cubes by string key
	fields map[string]struct{}
	// cache contains calculated string key hashes for searched cubes
	cache map[string]string

	// dimension ranges
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

// newCube is a cube constructor.
func newCube(cache map[string]string) *cube {
	return &cube{
		fields: make(map[string]struct{}),
		cache:  cache,
	}
}

// runIteration runs one cube iteration.
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

// newIterationState calculates the state of a cube in the new iteration.
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

// setActive sets a cube to active state.
func (c *cube) setActive(x, y, z int) {
	key := c.getKey(x, y, z)

	c.fields[key] = struct{}{}

	c.updateDimensions(x, y, z)
}

// setInactive sets a cube to inactive state.
func (c *cube) setInactive(x, y, z int) {
	key := c.getKey(x, y, z)

	delete(c.fields, key)
}

// isActive is used to check cube state.
func (c *cube) isActive(x, y, z int) bool {
	key := c.getKey(x, y, z)

	_, isActive := c.fields[key]
	return isActive
}

// getActiveCount returns the count of active cubes.
func (c *cube) getActiveCount() int {
	return len(c.fields)
}

// getKey is used to calculate the hash for accessing a cube by coordinates.
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

// updateDimensions is a helper function for cube dimension widening with iterations.
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
