package Day17

import (
	"Go-AdventOfCode2020/days/Day17/fourD"
	"Go-AdventOfCode2020/days/Day17/threeD"
	"Go-AdventOfCode2020/tools"
	"fmt"
)

func Goooo() {
	fmt.Println("--------- DAY 17 ---------")

	//lines := tools.ReadFile(("days/Day17/testInput.txt"))
	lines := tools.ReadFile(("days/Day17/input.txt"))

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
