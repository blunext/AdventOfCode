package d09

import (
	"fmt"
	"time"

	"AdventOfCode/tools"
)

type point struct {
	val    int8
	lowest bool
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
			line = append(line, point{val: int8(v)})
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

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #9 A")
	data := tools.ReadFile("y2021/days/d09/data.txt")
	matrix := prepareData(data)
	process(matrix)
	fmt.Printf("result: %v\n", riskPoints(matrix))
}
