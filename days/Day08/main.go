package Day08

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

type line int
type acumulator int

type action interface {
	process(line, acumulator) (line, acumulator)
	wasHere() bool
}

//---
type nop struct {
	visited bool
}

func NewNop() *nop {
	return &nop{}
}
func (n *nop) process(l line, a acumulator) (line, acumulator) {
	n.visited = true
	return l + 1, a
}
func (n *nop) wasHere() bool {
	return n.visited
}

//----
type jmp struct {
	value   int
	visited bool
}

func NewJmp(v int) *jmp {
	return &jmp{value: v}
}
func (j *jmp) process(l line, acu acumulator) (line, acumulator) {
	j.visited = true
	return l + line(j.value), acu
}
func (j *jmp) wasHere() bool {
	return j.visited
}

//----
type acc struct {
	value   int
	visited bool
}

func NewAcc(v int) *acc {
	return &acc{value: v}
}
func (a *acc) process(l line, acu acumulator) (line, acumulator) {
	a.visited = true
	return l + 1, acu + acumulator(a.value)
}
func (a *acc) wasHere() bool {
	return a.visited
}

//----

type program struct {
	lines []action
}

func Goooo() {
	fmt.Println("--------- DAY 08 ---------")
	//lines := tools.ReadFile(("days/Day08/testInput.txt"))
	lines := tools.ReadFile(("days/Day08/Input.txt"))

	program := populateProgram(lines)
	acc, finished := run(program)
	fmt.Printf("Part1: finished: %v, accumulator: %d", finished, acc)
	//
	//for _, instruction := range program.lines {
	//	if in
	//	acc, finished = run(program)
	//
	//}
	fmt.Printf("Part2: finished: %v, accumulator: %d", finished, acc)
}

func run(program program) (acumulator, bool) {
	finished := true
	var line line
	var acc acumulator

	for {
		if program.lines[line].wasHere() {
			finished = false
			break
		}
		line, acc = program.lines[line].process(line, acc)
	}
	return acc, finished
}

func populateProgram(lines []string) program {
	program := program{}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		i, _ := strconv.Atoi(tokens[1])

		var instruction action

		switch tokens[0] {
		case "nop":
			instruction = NewNop()
		case "acc":
			instruction = NewAcc(i)
		case "jmp":
			instruction = NewJmp(i)
		default:
			panic("sth wrong")
		}
		program.lines = append(program.lines, instruction)
	}
	return program
}
