package Day03

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
)

var slopes = []struct {
	x, y, testNo int
}{
	{1, 1, 2},
	{3, 1, 7},
	{5, 1, 3},
	{7, 1, 4},
	{1, 2, 2},
}

func check(lines []string, xStep, yStep int) int {
	x, y, t := 0, 0, 0
	for {
		switch field := lines[y][x]; field {
		case '#':
			t++
		}
		x += xStep
		y += yStep

		if y >= len(lines) {
			return t
		}

		if x >= len(lines[0]) {
			x = x - len(lines[0])
		}
	}
}
func Goooo() {
	fmt.Println("--------- DAY 03 ---------")
	lines := tools.ReadFile(("days/Day03/input.txt"))
	trees := check(lines, 3, 1)
	fmt.Printf("trees=%d\n", trees)

	fmt.Println("--------- DAY 03 part 2 ---------")
	for _, s := range slopes {
		no := check(lines, s.x, s.y)
		fmt.Printf("%d, %d = %d\n", s.x, s.y, no)
	}

}
