package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	return txtlines
}

func Goooo() {
	fmt.Println("--------- DAY 02 ---------")

	lines := readFile("days/day01/input.txt")
	ints := convertIntoInts(lines)

	for _, x := range ints {
		for _, y := range ints {
			if x+y == 2020 {
				fmt.Printf("%d + %d = %d -> %d\n", x, y, x+y, x*y)
			}
		}
	}
}

func convertIntoInts(lines []string) []int {
	ints := []int{}
	for _, s := range lines {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
