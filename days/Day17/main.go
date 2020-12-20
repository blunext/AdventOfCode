package Day17

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
)

type vector struct {
	x, y, z int
}

type cube struct {
	state, newState bool
}

type coordinates struct {
	x, y, z int
}

func (c coordinates) getCoordinates() (int, int, int) {
	return c.x, c.y, c.x
}

type grid struct {
	cubes         map[coordinates]*cube
	lookDirection []vector
}

func NewGrid() *grid {
	grid := grid{cubes: make(map[coordinates]*cube)}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				grid.lookDirection = append(grid.lookDirection, vector{x: x, y: y, z: z})
			}
		}
	}
	return &grid
}

func (g *grid) addCube(x, y, z int, state bool) {
	g.cubes[coordinates{x: x, y: y, z: z}] = &cube{state: state}
}

func (g *grid) process() {
	// we work on copy of the cubes map becoue we extend original map when necessary
	for coord, cube := range g.copyCubes() {
		cube.newState = g.checkNeighborsRules(coord, cube.state, true)
	}
	for _, cube := range g.cubes {
		cube.state = cube.newState
	}
}

func (g *grid) copyCubes() map[coordinates]*cube {
	cubesCopy := make(map[coordinates]*cube)
	for k, v := range g.cubes {
		cubesCopy[k] = v
	}
	return cubesCopy
}

func (g *grid) checkNeighborsRules(coord coordinates, state bool, canExtend bool) bool {
	count := g.countNeighbors(coord, canExtend)
	switch state {
	case true:
		if count == 2 || count == 3 {
			return true
		}
	case false:
		if count == 3 {
			return true
		}
	}
	return false
}

func (g *grid) countNeighbors(coord coordinates, canExtend bool) int {
	i := 0
	for _, v := range g.lookDirection {
		looking := coordinates{x: coord.x + v.x, y: coord.y + v.y, z: coord.z + v.z}
		if cube, ok := g.cubes[looking]; ok {
			if cube.state {
				i++
			}
			continue
		}
		// add new cube and check it
		if canExtend {
			g.addCube(looking.x, looking.y, looking.z, false)
			g.cubes[looking].newState = g.checkNeighborsRules(looking, false, false) // no extend
		}

	}
	return i
}

func Goooo() {
	fmt.Println("--------- DAY 17 ---------")

	//lines := tools.ReadFile(("days/Day17/testInput.txt"))
	lines := tools.ReadFile(("days/Day17/input.txt"))

	grid := NewGrid()

	initGrid(lines, grid)

	grid.process()
	grid.process()
	grid.process()
	grid.process()
	grid.process()
	grid.process()

	i := 0
	for _, cube := range grid.copyCubes() {
		if cube.state {
			i++
		}
	}
	fmt.Printf("Part 1: %d\n", i)
}

func initGrid(lines []string, grid *grid) {
	for y, line := range lines {
		for x, state := range line {
			switch state {
			case '.':
				grid.addCube(x, y, 0, false)
			case '#':
				grid.addCube(x, y, 0, true)
			default:
				panic("omg...")
			}
		}
	}
}
