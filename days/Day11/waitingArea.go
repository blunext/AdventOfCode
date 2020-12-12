package Day11

type state int

const (
	floor state = iota
	empty
	occupied
)

type vector struct {
	x, y int
}

var directions = []vector{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

type seat struct {
	state state
	flip  bool
}

type waitingArea struct {
	moved      bool
	seats      [][]seat
	lookFurher bool
}

func newWaitingArea() *waitingArea {
	return &waitingArea{seats: [][]seat{}}
}

func (w *waitingArea) addSeat(in string) {
	var occ state
	switch in {
	case ".":
		occ = floor
	case "L":
		occ = empty
	case "#":
		occ = occupied
	default:
		panic("bum....")
	}
	s := seat{state: occ}

	row := &w.seats[len(w.seats)-1]
	*row = append(*row, s)
}

func (w *waitingArea) addRow() {
	w.seats = append(w.seats, []seat{})
}

func (w *waitingArea) makeDecision(r, c, occupiedAcceptance int, lookFurther bool) {
	w.lookFurher = lookFurther
	if w.seats[r][c].state == occupied {
		if w.countOccupiedAround(r, c) >= occupiedAcceptance {
			w.changeState(r, c)
		}
	}
	if w.seats[r][c].state == empty {
		if w.countOccupiedAround(r, c) == 0 {
			w.changeState(r, c)
		}
	}
}

func (w *waitingArea) changeState(r int, c int) {
	w.moved = true
	w.seats[r][c].flip = true
}

func (w *waitingArea) countOccupiedAround(r, c int) int {
	count := 0
	for _, v := range directions {
		if w.lookIntoVectorDirection(r, c, v) {
			count++
		}
	}

	return count
}

func (w *waitingArea) lookIntoVectorDirection(r, c int, v vector) bool {
	newR := r + v.y
	newC := c + v.x
	if newR < 0 || newR >= w.rowsCount() {
		return false
	}
	if newC < 0 || newC >= w.colCount() {
		return false
	}
	switch w.seats[newR][newC].state {
	case occupied:
		return true
	case empty:
		return false
	default: // floor
		if w.lookFurher {
			return w.lookIntoVectorDirection(newR, newC, v)
		}
		return false
	}
}

func (w *waitingArea) traverseAll(process func(i, y int)) {
	for r := 0; r < w.rowsCount(); r++ {
		for c := 0; c < w.colCount(); c++ {
			process(r, c)
		}
	}
}

func (w *waitingArea) flipMarked() {
	w.traverseAll(func(r, c int) {
		if w.seats[r][c].flip {
			switch w.seats[r][c].state {
			case occupied:
				w.seats[r][c] = seat{state: empty}
			case empty:
				w.seats[r][c] = seat{state: occupied}
			}
		}
	})
}

func (w *waitingArea) clearMovement() {
	w.traverseAll(func(r, c int) {
		if w.seats[r][c].flip {
			w.seats[r][c] = seat{flip: false}
		}
	})
	w.moved = false
}

func (w *waitingArea) printSeat(r, c int) string {
	switch w.seats[r][c].state {
	case occupied:
		return "#"
	case empty:
		return "L"
	case floor:
		return "."
	default:
		return "ERROR"
	}
}

func (w *waitingArea) rowsCount() int {
	return len(w.seats)
}
func (w *waitingArea) colCount() int {
	return len(w.seats[0])
}

func (w *waitingArea) peopleMoved() bool {
	return w.moved
}
