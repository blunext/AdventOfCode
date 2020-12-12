package Day11

type state int

const (
	floor state = iota
	empty
	occupied
)

type seat struct {
	state state
	flip  bool
}

type waitingArea struct {
	moved bool
	seats [][]seat
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

func (w *waitingArea) makeDecision(r, c int) {
	if w.seats[r][c].state == occupied {
		if w.countOccupiedAround(r, c) >= 4 {
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
	for i := r - 1; i <= r+1; i++ {
		if w.seats[i][c-1].state == occupied {
			count++
		}
		if w.seats[i][c+1].state == occupied {
			count++
		}
	}
	if w.seats[r-1][c].state == occupied {
		count++
	}
	if w.seats[r+1][c].state == occupied {
		count++
	}
	return count
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
