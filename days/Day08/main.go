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
	line []action
}

func Goooo() {
	fmt.Println("--------- DAY 08 ---------")
	//lines := tools.ReadFile(("days/Day08/testInput.txt"))
	lines := tools.ReadFile(("days/Day08/Input.txt"))

	program := populateProgram(lines)

	var line line
	var acc acumulator

	for {
		if program.line[line].wasHere() {
			fmt.Printf("accumulator: %d", acc)
			break
		}
		line, acc = program.line[line].process(line, acc)
	}

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
		program.line = append(program.line, instruction)
	}
	return program
}
