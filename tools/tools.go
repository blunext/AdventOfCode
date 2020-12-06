package tools

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []string {
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

func ConvertIntoInts(lines []string) []int {
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

func CombineLines(lines []string) []string {
	set := []string{}
	new := true
	for _, line := range lines {
		if line == "" {
			new = true
			continue
		}
		if new {
			set = append(set, line)
			new = false
			continue
		}
		set[len(set)-1] = set[len(set)-1] + " " + line
	}
	return set
}
