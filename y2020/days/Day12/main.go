package Day12

import (
	"fmt"
	"strconv"

	"AdventOfCode/tools"
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
	x, y             int
	hading           int
	wPointX, wPointY int
}

func newShip() *ship {
	return &ship{wPointX: 10, wPointY: 1}
}
func (s *ship) move(instr string) {
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
func (s *ship) rotateWaypoint(d int, clockWise int) {
	switch d {
	case 90:
		x := s.wPointX
		s.wPointX = s.wPointY * clockWise
		s.wPointY = -x * clockWise
	case 180:
		s.wPointX = -s.wPointX
		s.wPointY = -s.wPointY
	case 270:
		x := s.wPointX
		s.wPointX = -s.wPointY * clockWise
		s.wPointY = x * clockWise
	}
}

func (s *ship) moveWithWaypoint(instr string) {
	letter := instr[:1]
	value, _ := strconv.Atoi(instr[1:])
	switch letter {
	case "F":
		s.x += s.wPointX * value
		s.y += s.wPointY * value
	case "R":
		s.rotateWaypoint(value, 1)
	case "L":
		s.rotateWaypoint(value, -1)
	case "N":
		s.wPointY += value
	case "S":
		s.wPointY -= value
	case "E":
		s.wPointX += value
	case "W":
		s.wPointX -= value
	default:
		panic("Errr....")
	}
}
func Goooo() {
	fmt.Println("--------- DAY 12 ---------")
	// lines := tools.ReadLines(("y2020/days/Day12/testInput.txt"))
	lines := tools.ReadLines(("y2020/days/Day12/input.txt"))
	ship := newShip()
	for _, instr := range lines {
		ship.move(instr)
	}
	fmt.Printf("x=%d, y=%d\n", ship.x, ship.y)

	ship = newShip()
	for _, instr := range lines {
		ship.moveWithWaypoint(instr)
	}
	fmt.Printf("x=%d, y=%d\n", ship.x, ship.y)
}
