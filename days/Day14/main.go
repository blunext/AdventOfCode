package Day14

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

func convertMask(m string) (uint64, uint64) {
	return 0, 0
}

func Goooo() {
	fmt.Println("--------- DAY 14 ---------")
	//lines := tools.ReadFile(("days/Day14/testInput.txt"))
	lines := tools.ReadFile(("days/Day14/input.txt"))

	memory := make(map[int64]int64)
	andMask := int64(0)
	orMask := int64(0)
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			andMask, orMask = getMasks(line)
			//fmt.Printf("andMask: %d, orMask: %d\n", andMask, orMask)
			continue
		}
		mem, val := getInstruction(line)

		memory[mem] = val&andMask | orMask
		//fmt.Printf("mem: %d, val, %d\n", mem, val)
	}

	i := int64(0)
	for _, v := range memory {
		i += v
	}

	fmt.Printf("part 1: %d\n", i)

}

func getMasks(line string) (int64, int64) {
	t := strings.Split(line, "=")
	t[1] = strings.TrimSpace(t[1])

	andMask := strings.ReplaceAll(t[1], "X", "1")
	orMask := strings.ReplaceAll(t[1], "X", "0")

	andVal, _ := strconv.ParseInt(andMask, 2, 64)
	orVal, _ := strconv.ParseInt(orMask, 2, 64)

	return andVal, orVal
}

func getInstruction(line string) (int64, int64) {
	t := strings.Split(line, "=")
	m := strings.Replace(strings.TrimSpace(t[0]), "mem[", "", 1)
	m = strings.Replace(m, "]", "", 1)

	mem, _ := strconv.Atoi(m)
	val, _ := strconv.Atoi(strings.TrimSpace(t[1]))
	return int64(mem), int64(val)
}
