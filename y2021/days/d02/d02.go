package d02

import (
	"fmt"
	"log"

	"AdventOfCode/tools"
)

type command struct {
	name                 string
	horizontalPos, depth int
}

var (
	horizontalPos = 0
	depth         = 0
	commands      = []command{
		{"forward", 1, 0},
		{"down", 0, 1},
		{"up", 0, -1},
	}
)

func process(data [][]string) int {
	for _, commandData := range data {
		command := getCommand(commandData[0])
		execute(command, tools.StrToInt(commandData[1]))
	}
	return horizontalPos * depth
}

func execute(command command, value int) {
	horizontalPos += command.horizontalPos * value
	depth += command.depth * value
}

func getCommand(commandName string) command {
	for _, c := range commands {
		if c.name == commandName {
			return c
		}
	}
	log.Fatalf("no commad found")
	return command{}
}

func Main() {
	data := tools.ReadByWords("y2021/days/d02/data.txt", 2)
	fmt.Println("DAY #2 A")
	fmt.Printf("result: %d\n", process(data))
}
