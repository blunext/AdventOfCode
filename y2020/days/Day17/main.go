package Day17

import (
	"fmt"

	"AdventOfCode/tools"
	"AdventOfCode/y2020/days/Day17/fourD"
	"AdventOfCode/y2020/days/Day17/threeD"
)

func Goooo() {
	fmt.Println("--------- DAY 17 ---------")

	// lines := tools.ReadLines(("y2020/days/Day17/testInput.txt"))
	lines := tools.ReadLines(("y2020/days/Day17/input.txt"))

	grid := threeD.NewGrid()

	threeD.InitGrid(lines, grid)

	grid.Process()
	grid.Process()
	grid.Process()
	grid.Process()
	grid.Process()
	grid.Process()

	i := 0
	for _, cube := range grid.CopyCubes() {
		if cube.State {
			i++
		}
	}
	fmt.Printf("Part 1: %d\n", i)

	// -------------------

	grid4 := fourD.NewGrid()

	fourD.InitGrid(lines, grid4)

	grid4.Process()
	grid4.Process()
	grid4.Process()
	grid4.Process()
	grid4.Process()
	grid4.Process()

	i = 0
	for _, cube := range grid4.CopyCubes() {
		if cube.State {
			i++
		}
	}
	fmt.Printf("Part 2: %d\n", i)

}
