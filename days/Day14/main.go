package Day14

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Goooo() {
	fmt.Println("--------- DAY 14 ---------")
	//lines := tools.ReadFile(("days/Day14/testInput.txt"))
	lines := tools.ReadFile(("days/Day14/input.txt"))

	i := part1(lines)
	fmt.Printf("part 1: %d\n", i)

	i = part2(lines)
	fmt.Printf("part 2: %d\n", i)
}

func part1(lines []string) uint64 {
	memory := make(map[uint64]uint64)
	andMask := uint64(0)
	orMask := uint64(0)
	for _, line := range lines {
		if strings.Contains(line, "mask") {
			andMask, orMask = getBitMasks(line)
			//fmt.Printf("andMask: %d, orMask: %d\n", andMask, orMask)
			continue
		}
		address, val := getInstruction(line)

		memory[address] = val&andMask | orMask
		//fmt.Printf("mem: %d, val, %d\n", mem, val)
	}

	i := uint64(0)
	for _, v := range memory {
		i += v
	}
	return i
}

func getBitMasks(line string) (uint64, uint64) {
	t := strings.Split(line, "=")
	t[1] = strings.TrimSpace(t[1])

	andMask := strings.ReplaceAll(t[1], "X", "1")
	orMask := strings.ReplaceAll(t[1], "X", "0")

	andVal, _ := strconv.ParseInt(andMask, 2, 64)
	orVal, _ := strconv.ParseInt(orMask, 2, 64)
	return uint64(andVal), uint64(orVal)
}

func getInstruction(line string) (uint64, uint64) {
	t := strings.Split(line, "=")
	m := strings.Replace(strings.TrimSpace(t[0]), "mem[", "", 1)
	m = strings.Replace(m, "]", "", 1)

	mem, _ := strconv.Atoi(m)
	val, _ := strconv.Atoi(strings.TrimSpace(t[1]))
	return uint64(mem), uint64(val)
}

func part2(lines []string) uint64 {
	memory := make(map[uint64]uint64)

	var mask string
	for _, line := range lines {
		maskPattern := maskRgx.FindStringSubmatch(line)
		if len(maskPattern) != 0 {
			mask = maskPattern[1]
			continue
		}

		memMatch := memRgx.FindStringSubmatch(line)

		address, _ := strconv.ParseUint(memMatch[1], 10, 64)

		addresses := []uint64{0}

		for i := 0; i < 36; i++ {
			bitIndex := 35 - i
			switch mask[i] {
			case '0':
				for j := range addresses {
					if tools.GetBit(address, bitIndex) == 1 {
						addresses[j] = tools.SetBit(addresses[j], bitIndex)
					}
				}
			case '1':
				for j := range addresses {
					addresses[j] = tools.SetBit(addresses[j], bitIndex)
				}
			case 'X':
				for j := range addresses {
					addresses = append(addresses, addresses[j])
					addresses[j] = tools.SetBit(addresses[j], bitIndex)
				}
			}
		}

		value, _ := strconv.ParseUint(memMatch[2], 10, 64)

		for _, address := range addresses {
			memory[address] = value
		}
	}

	var sum uint64
	for _, val := range memory {
		sum += val
	}
	return sum
}

var maskRgx = regexp.MustCompile("^mask = (.+)$")
var memRgx = regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")
