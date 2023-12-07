package d09

import (
	"fmt"
	"log"
	"sort"
	"time"

	"AdventOfCode/tools"
)

type point struct {
	val     byte
	lowest  bool
	flooded bool
	basin   int
}

func prepareData(data []string) [][]point {
	var p9 = point{val: 9}
	var rows [][]point
	var line []point

	x := len(data[0])
	for i := 0; i < x+2; i++ {
		line = append(line, p9)
	}
	rows = append(rows, line)
	for _, s := range data {
		line = []point{p9}
		ints := tools.ConvertStrOfDigitsToInts(s)
		for _, v := range ints {
			line = append(line, point{val: byte(v)})
		}
		line = append(line, p9)
		rows = append(rows, line)
	}
	line = []point{}
	for i := 0; i < x+2; i++ {
		line = append(line, p9)
	}
	rows = append(rows, line)
	return rows
}

func process(matrix [][]point) {
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x <= len(matrix[y])-1; x++ {
			val := matrix[y][x].val
			if val < matrix[y-1][x].val && val < matrix[y+1][x].val && val < matrix[y][x-1].val && val < matrix[y][x+1].val {
				matrix[y][x].lowest = true
			}
		}
	}
}

func riskPoints(matrix [][]point) int {
	sum := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x <= len(matrix[y])-1; x++ {
			if matrix[y][x].lowest {
				sum += int(matrix[y][x].val) + 1
			}
		}
	}
	return sum
}

func makeRain(matrix [][]point) {
	basin := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x <= len(matrix[y])-1; x++ {
			if matrix[y][x].lowest {
				basin++
				flood(matrix, y, x, basin)
			}
		}
	}
}

func flood(matrix [][]point, y, x, basin int) {
	matrix[y][x].flooded = true
	matrix[y][x].basin = basin
	val := matrix[y][x].val
	floodNeighbour(matrix, y-1, x, val, basin)
	floodNeighbour(matrix, y+1, x, val, basin)
	floodNeighbour(matrix, y, x-1, val, basin)
	floodNeighbour(matrix, y, x+1, val, basin)
}

func floodNeighbour(matrix [][]point, y, x int, val byte, basin int) {
	p := matrix[y][x]
	if p.val != 9 && !p.flooded && val < p.val {
		flood(matrix, y, x, basin)
	}
}

func countBasinSizes(matrix [][]point) []int {
	basins := make(map[int]int)
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x <= len(matrix[y])-1; x++ {
			if matrix[y][x].flooded {
				if _, ok := basins[matrix[y][x].basin]; ok {
					basins[matrix[y][x].basin]++
					continue
				}
				basins[matrix[y][x].basin] = 1
			}
		}
	}
	var basinsSurfaces []int
	for _, v := range basins {
		basinsSurfaces = append(basinsSurfaces, v)
	}
	sort.Ints(basinsSurfaces)
	return basinsSurfaces
}

func threeLargestBasinMultiplied(basin []int) int {
	if len(basin) < 3 {
		log.Fatal("too less basins")
	}
	sum := 1
	for i := len(basin) - 1; i > len(basin)-4; i-- {
		sum *= basin[i]
	}
	return sum
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #9 A")
	data := tools.ReadLines("y2021/days/d09/data.txt")
	matrix := prepareData(data)
	process(matrix)
	fmt.Printf("result: %v\n", riskPoints(matrix))

	fmt.Println("DAY #9 B")
	makeRain(matrix)
	basins := countBasinSizes(matrix)
	fmt.Printf("result: %v\n", threeLargestBasinMultiplied(basins))
}
