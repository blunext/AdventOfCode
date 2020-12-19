package Day16

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type rules struct {
	//name string
	range1min, range1max, range2min, range2max int
}

func Goooo() {
	fmt.Println("--------- DAY 16 ---------")
	//lines := tools.ReadFile(("days/Day16/testInput.txt"))
	lines := tools.ReadFile(("days/Day16/input.txt"))

	rulesCollection, lineIndex := getRules(lines)

	if lines[lineIndex+1] != "your ticket:" {
		panic("omg...")
	}
	lineIndex += 4
	if lines[lineIndex] != "nearby tickets:" {
		panic("omg again...")
	}
	lineIndex++
	errorRate := 0
	for ; lineIndex < len(lines); lineIndex++ {
		neabyTicket := tools.ConvertCommaSeparatedStrIntoInts(lines[lineIndex])
		for _, num := range neabyTicket {
			if !checkRules(rulesCollection, num) {
				//fmt.Println("no: ", num)
				errorRate += num
			}
		}
	}

	fmt.Printf("Part 1: %d\n", errorRate)

}

func checkRules(rulesCollection []rules, num int) bool {
	for _, rule := range rulesCollection {
		if (num >= rule.range1min && num <= rule.range1max) || (num >= rule.range2min && num <= rule.range2max) {
			return true
		}
	}
	return false
}

func getRules(lines []string) ([]rules, int) {

	rulesCollection := []rules{}
	i := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		r := rules{}
		newLine := line[strings.Index(line, ":")+1:]
		_, err := fmt.Sscanf(newLine, " %d-%d or %d-%d", &r.range1min, &r.range1max, &r.range2min, &r.range2max)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s: %d-%d or %d-%d\n", rules.range1min, rules.range1max, rules.range2min, rules.range2max)
		rulesCollection = append(rulesCollection, r)
		i++
	}
	return rulesCollection, i
}
