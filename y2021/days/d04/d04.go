package d04

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"AdventOfCode/tools"
)

type field struct {
	number int
	marked bool
}

type board struct {
	fields    [][]field
	winOrder  int
	winNumber int
	winner    bool
}

func populateBoards(input []string) []board {
	var boards []board
	var theBoard board
	i := 0
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		var fields []field
		nums, err := tools.ReadWordsFromLine(strings.NewReader(line), 5)
		if err != nil {
			fmt.Errorf("error: %w\n", err)
		}
		for _, num := range nums {
			field := field{number: tools.StrToInt(num)}
			fields = append(fields, field)
		}
		theBoard.fields = append(theBoard.fields, fields)
		i++
		if i == 5 {
			boards = append(boards, theBoard)
			i = 0
			theBoard = board{}
		}
	}
	return boards
}

func game(boards []board, numbers []int, firstLast bool) int {
	winOrder := 0
	for _, number := range numbers {
		markFields(boards, number)
		markWinnerBoard(boards, number, winOrder)
		winOrder++
	}

	winnerId := 0
	switch firstLast {
	case true:
		order := 999999999
		for id := range boards {
			if !boards[id].winner {
				continue
			}
			if order > boards[id].winOrder {
				order = boards[id].winOrder
				winnerId = id
			}
		}
	default:
		order := -1
		for id := range boards {
			if !boards[id].winner {
				continue
			}
			if order < boards[id].winOrder {
				order = boards[id].winOrder
				winnerId = id
			}
		}
	}

	sum := sumWinnerUnmarkedFields(boards[winnerId])
	fmt.Printf("sum: %d, num: %d", sum, boards[winnerId].winNumber)
	draw(boards[winnerId])
	return sum * boards[winnerId].winNumber
}

func sumWinnerUnmarkedFields(winner board) int {
	sum := 0
	for _, fields := range winner.fields {
		for fi := range fields {
			if !fields[fi].marked {
				sum += fields[fi].number
			}
		}
	}
	return sum
}

func markWinnerBoard(boards []board, number int, order int) {
	for id := range boards {
		if boards[id].winner {
			continue
		}
		for x := 0; x < 5; x++ {
			marked := 0
			for y := 0; y < 5; y++ {
				if boards[id].fields[x][y].marked {
					marked++
				}
			}
			if marked == 5 {
				boards[id].winner = true
				boards[id].winNumber = number
				boards[id].winOrder = order
			}
		}

		for y := 0; y < 5; y++ {
			marked := 0
			for x := 0; x < 5; x++ {
				if boards[id].fields[x][y].marked {
					marked++
				}
			}
			if marked == 5 {
				boards[id].winner = true
				boards[id].winNumber = number
				boards[id].winOrder = order
			}
		}
	}
}

func markFields(boards []board, number int) {
	for _, board := range boards {
		if board.winner {
			continue
		}
		for _, fields := range board.fields {
			for fi := range fields {
				if fields[fi].number == number {
					fields[fi].marked = true
				}
			}
		}
	}
}

func draw(board board) {
	fmt.Println()
	c := color.New(color.FgBlue)
	for _, fields := range board.fields {
		for fi := range fields {
			if fields[fi].marked {
				c.Printf("%d ", fields[fi].number)
			} else {
				fmt.Printf("%d ", fields[fi].number)
			}
		}
		fmt.Println()
	}
}

func Main() {
	fmt.Println("DAY #4 A")
	data := tools.ReadLines("y2021/days/d04/data.txt")
	numbers := tools.ConvertCommaSeparatedStrIntoInts(data[0])
	boards := populateBoards(data[1:])
	result := game(boards, numbers, true)
	fmt.Printf("result: %d\n", result)
	fmt.Println("DAY #4 B")
	boards = populateBoards(data[1:])
	result = game(boards, numbers, false)
	fmt.Printf("result: %d\n", result)
}
