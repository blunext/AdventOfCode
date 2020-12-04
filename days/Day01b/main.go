package day01

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
)

func Goooo() {
	fmt.Println("--------- DAY 02 ---------")
	lines := tools.ReadFile(("days/Day01/input.txt"))
	ints := tools.ConvertIntoInts(lines)

	for _, x := range ints {
		for _, y := range ints {
			for _, z := range ints {
				if x+y+z == 2020 {
					fmt.Printf("%d + %d + %d = %d -> %d\n", x, y, z, x+y+z, x*y*z)
				}
			}
		}
	}
}
