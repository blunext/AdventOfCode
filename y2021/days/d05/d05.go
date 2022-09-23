package d05

import (
	"fmt"
	"log"
	"math"

	"AdventOfCode/tools"
)

type point struct {
	x, y int
}
type vector struct {
	start, end point
}

type board struct {
	fields map[point]int
}

func getVectors(data []string) []vector {
	var vectors []vector
	for _, line := range data {
		vec := vector{}
		n, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &vec.start.x, &vec.start.y, &vec.end.x, &vec.end.y)
		if err != nil {
			log.Fatalln(err)
		}
		if n != 4 {
			log.Fatalf("wrong number of fields")
		}
		vectors = append(vectors, vec)
	}
	return vectors
}

func mapVectors(vectors []vector, diagonals bool) board {
	board := board{fields: map[point]int{}}
	for _, v := range vectors {
		if diagonals || (v.start.x == v.end.x) || (v.start.y == v.end.y) {
			points := getPath(v, diagonals)
			for _, p := range points {
				if _, ok := board.fields[p]; ok {
					board.fields[p]++
				} else {
					board.fields[p] = 1
				}
			}
		}
	}
	return board
}

func getPath(vect vector, diagonals bool) []point {
	var path []point
	switch {
	case vect.start.x == vect.end.x:
		if vect.start.y <= vect.end.y {
			for i := vect.start.y; i <= vect.end.y; i++ {
				p := point{x: vect.start.x, y: i}
				path = append(path, p)
			}
		} else {
			for i := vect.start.y; i >= vect.end.y; i-- {
				p := point{x: vect.start.x, y: i}
				path = append(path, p)
			}
		}
	case vect.start.y == vect.end.y:
		if vect.start.x <= vect.end.x {
			for i := vect.start.x; i <= vect.end.x; i++ {
				p := point{x: i, y: vect.start.y}
				path = append(path, p)
			}
		} else {
			for i := vect.start.x; i >= vect.end.x; i-- {
				p := point{x: i, y: vect.start.y}
				path = append(path, p)
			}
		}
	case diagonals:
		if diagonalVector(vect) {
			xStep := 0
			if vect.start.x >= vect.end.x {
				xStep = -1
			} else {
				xStep = 1
			}
			yStep := 0
			if vect.start.y >= vect.end.y {
				yStep = -1
			} else {
				yStep = 1
			}
			movablePoint := vect.start
			path = append(path, movablePoint)
			for movablePoint.x != vect.end.x {
				movablePoint.x += xStep
				movablePoint.y += yStep
				path = append(path, movablePoint)
			}
		}
	}
	return path
}

func countOverlaps(b board) int {
	count := 0
	for _, overlaps := range b.fields {
		if overlaps > 1 {
			count++
		}
	}
	return count
}

func diagonalVector(v vector) bool {
	xx := float64(v.start.x - v.end.x)
	yy := float64(v.start.y - v.end.y)
	if math.Abs(xx) == math.Abs(yy) {
		return true
	}
	return false
}

func Main() {
	fmt.Println("DAY #5 A")
	data := tools.ReadFile("y2021/days/d05/data.txt")
	vectors := getVectors(data)
	board := mapVectors(vectors, false)
	fmt.Printf("result: %d\n", countOverlaps(board))

	fmt.Println("DAY #5 A")
	board = mapVectors(vectors, true)
	fmt.Printf("result: %d\n", countOverlaps(board))
}
