package d10

import (
	"fmt"
	"time"

	"AdventOfCode/tools"
)

type chank struct {
	start, end byte
}

type stack []chank

func (s *stack) Push(v chank) {
	*s = append(*s, v)
}

func (s *stack) Pop() chank {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

var brackets = map[byte]byte{'(': ')', '[': ']', '{': '}', '<': '>'}
var points = map[byte]int{')': 3, ']': 57, '}': 1197, '>': 25137}

func process(data []string) int {
	var illegalChars []byte

	for _, line := range data {
		stack := stack{}
		for i := range line {
			char := line[i]
			switch char {
			case '(', '[', '{', '<':
				chunk := chank{start: char, end: brackets[char]}
				stack.Push(chunk)
			case ')', ']', '}', '>':
				chunk := stack.Pop()
				if char != chunk.end {
					illegalChars = append(illegalChars, char)
					// fmt.Printf(" bad closing: %v, expected, %v ", string(chunk.end), string(char))
				}
			}
		}
	}
	sum := 0
	for _, char := range illegalChars {
		sum += points[char]
	}
	return sum
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #10 A")
	data := tools.ReadLines("y2021/days/d10/data.txt")

	fmt.Printf("result: %v\n", process(data))
}
