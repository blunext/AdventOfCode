package Day06

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/tools"
)

func Goooo() {
	fmt.Println("--------- DAY 06 ---------")
	// lines := tools.ReadLines(("y2020/days/Day06/testInput.txt"))
	lines := tools.ReadLines(("y2020/days/Day06/input.txt"))
	entries := tools.CombineLines(lines)
	a := 0
	b := 0
	for _, entry := range entries {
		a += countLetters(entry)
		b += countSame(entry)
	}
	fmt.Printf("a=%d, b=%d\n", a, b)
}

func countLetters(entry string) int {
	count := 0
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		if strings.Contains(entry, string(letter)) {
			count += 1
		}
	}
	return count
}

func countSame(entry string) int {
	persons := strings.Split(entry, " ")
	pattern, _ := strconv.ParseInt("11111111111111111111111111", 2, 32)
	for _, p := range persons {
		pattern = pattern & convertToInt(p)
	}
	return strings.Count(strconv.FormatInt(pattern, 2), "1")
}

func convertToInt(p string) int64 {
	bits := "abcdefghijklmnopqrstuvwxyz"
	for _, letter := range p {
		bits = strings.Replace(bits, string(letter), "1", 1)
	}
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		bits = strings.Replace(bits, string(letter), "0", 1)
	}

	pattern, _ := strconv.ParseInt(bits, 2, 32)
	return pattern
}
