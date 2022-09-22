package fourD

type vector struct {
	x, y, z, w int
}

type cube struct {
	State, newState bool
}

type coordinates struct {
	x, y, z, w int
}

func (c coordinates) getCoordinates() (int, int, int, int) {
	return c.x, c.y, c.x, c.w
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
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					grid.lookDirection = append(grid.lookDirection, vector{x: x, y: y, z: z, w: w})
				}
			}
		}
	}
	return &grid
}

func (g *grid) addCube(x, y, z, w int, state bool) {
	g.cubes[coordinates{x: x, y: y, z: z, w: w}] = &cube{State: state}
}

func (g *grid) Process() {
	// we work on copy of the cubes map becoue we extend original map when necessary
	for coord, cube := range g.CopyCubes() {
		cube.newState = g.checkNeighborsRules(coord, cube.State, true)
	}
	for _, cube := range g.cubes {
		cube.State = cube.newState
	}
}

func (g *grid) CopyCubes() map[coordinates]*cube {
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
		looking := coordinates{x: coord.x + v.x, y: coord.y + v.y, z: coord.z + v.z, w: coord.w + v.w}
		if cube, ok := g.cubes[looking]; ok {
			if cube.State {
				i++
			}
			continue
		}
		// add new cube and check it
		if canExtend {
			g.addCube(looking.x, looking.y, looking.z, looking.w, false)
			g.cubes[looking].newState = g.checkNeighborsRules(looking, false, false) // no extend
		}

	}
	return i
}

func InitGrid(lines []string, grid *grid) {
	for y, line := range lines {
		for x, state := range line {
			switch state {
			case '.':
				grid.addCube(x, y, 0, 0, false)
			case '#':
				grid.addCube(x, y, 0, 0, true)
			default:
				panic("omg...")
			}
		}
	}
}
