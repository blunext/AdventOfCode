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
	getValue() int
}

type program struct {
	lines []action
}

//NOP
type nop struct {
	value int
}

func NewNop(v int) *nop {
	return &nop{value: v}
}
func (n *nop) process(l line, a acumulator) (line, acumulator) {
	return l + 1, a
}
func (n *nop) getValue() int {
	return n.value
}

//JMP
type jmp struct {
	value int
}

func NewJmp(v int) *jmp {
	return &jmp{value: v}
}
func (j *jmp) process(l line, acu acumulator) (line, acumulator) {
	return l + line(j.value), acu
}
func (j *jmp) getValue() int {
	return j.value
}

//ACC
type acc struct {
	value int
}

func NewAcc(v int) *acc {
	return &acc{value: v}
}
func (a *acc) process(l line, acu acumulator) (line, acumulator) {
	return l + 1, acu + acumulator(a.value)
}
func (a *acc) getValue() int {
	return a.value
}

func Goooo() {
	fmt.Println("--------- DAY 08 ---------")
	//lines := tools.ReadFile(("days/Day08/testInput.txt"))
	lines := tools.ReadFile(("days/Day08/Input.txt"))

	myProgram := populateProgram(lines)

	acum, finished := run(myProgram)
	fmt.Printf("Part1: finished: %v, accumulator: %d\n", finished, acum)

	for i, instruction := range myProgram.lines {
		var prevAction action
		var newAction action

		switch v := instruction.(type) {
		case *acc:
			newAction, prevAction = v, v
		case *jmp:
			prevAction = v
			newAction = NewNop(v.getValue())
		case *nop:
			prevAction = v
			newAction = NewJmp(v.getValue())
		default:
			panic("sth wrong")
		}

		myProgram.lines[i] = newAction
		acum, finished = run(myProgram)
		if finished {
			break
		}
		myProgram.lines[i] = prevAction
	}
	fmt.Printf("Part2: finished: %v, accumulator: %d\n", finished, acum)
}

func run(program program) (acumulator, bool) {
	var ln line
	var acc acumulator
	finished := true
	visited := make(map[line]bool)
	for {
		if _, ok := visited[ln]; ok {
			finished = false
			break
		}
		visited[ln] = true
		ln, acc = program.lines[ln].process(ln, acc)

		if int(ln) >= len(program.lines) {
			finished = true
			break
		}
	}
	return acc, finished
}

func populateProgram(lines []string) program {
	program := program{}
	for _, line := range lines {
		tokens := strings.Split(line, " ")

		val, _ := strconv.Atoi(tokens[1])

		var instruction action

		switch tokens[0] {
		case "nop":
			instruction = NewNop(val)
		case "acc":
			instruction = NewAcc(val)
		case "jmp":
			instruction = NewJmp(val)
		default:
			panic("sth wrong")
		}
		program.lines = append(program.lines, instruction)

	}
	return program
}
