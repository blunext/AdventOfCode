package Day11

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
)

func Goooo() {
	fmt.Println("--------- DAY 11 ---------")

	//lines := tools.ConvertIntoInts(tools.ReadFile(("days/Day10/testInput.txt")))
	//lines := tools.ReadFile(("days/Day11/testInput.txt"))
	lines := tools.ReadFile(("days/Day11/input.txt"))

	waiting := newWaitingArea()
	waiting.addRow()

	// add 1st row above + 1 floor left + 1 floor right
	for i := 0; i < len(lines[0])+2; i++ {
		waiting.addSeat(string("."))
	}

	for _, line := range lines {
		waiting.addRow()
		waiting.addSeat(string(".")) //left flour
		for _, in := range line {
			waiting.addSeat(string(in))
		}
		waiting.addSeat(string(".")) //right flour
	}

	// add last empty row  + 1 floor left + 1 floor right
	waiting.addRow()
	for i := 0; i < len(lines[0])+2; i++ {
		waiting.addSeat(string("."))
	}

	printState(waiting)
	fmt.Println()

	for {
		waiting.clearMovement()
		for r := 0; r < waiting.rowsCount(); r++ {
			for c := 0; c < waiting.colCount(); c++ {
				waiting.makeDecision(r, c)
			}
		}

		waiting.flipMarked()

		//printState(waiting)

		if !waiting.peopleMoved() {
			break
		}
		//time.Sleep(time.Second)
		//fmt.Println()
	}

	fmt.Println(countOcupied(waiting))
}

func printState(waiting *waitingArea) {
	for r := 0; r < waiting.rowsCount(); r++ {
		for c := 0; c < waiting.colCount(); c++ {
			fmt.Print(waiting.printSeat(r, c))
		}
		fmt.Println()
	}
}

func countOcupied(waiting *waitingArea) int {
	count := 0
	waiting.traverseAll(func(r, c int) {
		if waiting.seats[r][c].state == occupied {
			count++
		}
	})
	return count
}
