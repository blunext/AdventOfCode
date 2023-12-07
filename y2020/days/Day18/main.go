package Day18

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/tools"
)

type operation func(x, y int) int
type operator func(string) operation

func Goooo() {
	fmt.Println("--------- DAY 18 ---------")

	// lines := tools.ReadLines(("y2020/days/Day18/testInput.txt"))
	lines := tools.ReadLines(("y2020/days/Day18/input.txt"))

	sum := 0
	for _, line := range lines {
		tokens := preProcess(line)
		result, _ := process(tokens)
		sum += result
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func preProcess(in string) []string {
	in = strings.ReplaceAll(in, "(", "( ")
	in = strings.ReplaceAll(in, ")", " )")
	tokens := strings.Split(in, " ")
	return tokens
}

func process(in []string) (int, int) {
	val := 0
	oper := add()
	for i := 0; i < len(in); i++ {
		token := in[i]
		number, err := strconv.Atoi(token)
		if err != nil {
			switch in[i] {
			case "+":
				oper = add()
				continue
			case "*":
				oper = mul()
				continue
			case "(":
				innerValue, shift := process(in[i+1:])
				val = oper(val, innerValue)
				i += shift + 1
				continue
			case ")":
				return val, i
			default:
				panic("Omg...")
			}
		}
		val = oper(val, number)
	}
	return val, len(in) - 1
}

func add() operation {
	return func(x, y int) int {
		return x + y
	}
}

func mul() operation {
	return func(x, y int) int {
		return x * y
	}
}
