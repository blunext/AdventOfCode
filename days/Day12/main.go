package Day12

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
)

type direction struct {
	dirX, dirY int
}

var degrees = map[int]direction{
	0:   {1, 0},
	90:  {0, -1},
	180: {-1, 0},
	270: {0, 1},
}

type ship struct {
	x, y   int
	hading int
}

func newShip() *ship {
	return &ship{x: 0, y: 0, hading: 0}
}
func (s *ship) processIntruction(instr string) {
	letter := instr[:1]
	value, _ := strconv.Atoi(instr[1:])

	switch letter {
	case "F":
		s.x += degrees[s.hading].dirX * value
		s.y += degrees[s.hading].dirY * value
	case "L":
		s.hading -= value
		if s.hading < 0 {
			s.hading = s.hading + 360
		}
	case "R":
		s.hading += value
		if s.hading >= 360 {
			s.hading = s.hading - 360
		}
	case "N":
		s.y += value
	case "S":
		s.y -= value
	case "E":
		s.x += value
	case "W":
		s.x -= value
	default:
		panic("Errr....")
	}
}

func Goooo() {
	fmt.Println("--------- DAY 12 ---------")
	//lines := tools.ReadFile(("days/Day12/testInput.txt"))
	lines := tools.ReadFile(("days/Day12/input.txt"))
	ship := newShip()
	for _, instr := range lines {
		ship.processIntruction(instr)
	}
	fmt.Printf("x=%d, y=%d", ship.x, ship.y)
}
