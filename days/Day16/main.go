package Day16

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type rules struct {
	name                                       string
	range1min, range1max, range2min, range2max int
}

type ticket struct {
	valid  bool
	fields []int
}

type tickets struct {
	collection []*ticket
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

	nerbyTickets := getNearbyTickets(lineIndex, lines, rulesCollection)
	//errorRate, ticketsToRemove := funcName(lineIndex, lines, rulesCollection)
	errorRate := checkTicketValidity(nerbyTickets, rulesCollection)

	removeInvalidTickets(&nerbyTickets)

	fmt.Printf("Part 1: %d\n", errorRate) //25895

}

func removeInvalidTickets(nerbyTickets *tickets) {
	newTickets := []*ticket{}
	for _, t := range nerbyTickets.collection {
		if t.valid {
			newTickets = append(newTickets, t)
		}
	}
	nerbyTickets.collection = newTickets
}

func checkTicketValidity(nerbyTickets tickets, rulesCollection []rules) int {
	errorRate := 0
	for _, ticket := range nerbyTickets.collection {
		for _, num := range ticket.fields {
			if !checkRules(rulesCollection, num) {
				//fmt.Println("no: ", num)
				errorRate += num
				ticket.valid = false
			}
		}
	}
	return errorRate
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func getNearbyTickets(lineIndex int, lines []string, rulesCollection []rules) tickets {
	neabyTickets := tickets{}
	for ; lineIndex < len(lines); lineIndex++ {
		f := tools.ConvertCommaSeparatedStrIntoInts(lines[lineIndex])
		t := ticket{valid: true, fields: f}
		neabyTickets.collection = append(neabyTickets.collection, &t)
	}

	return neabyTickets
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

		if strings.Contains(line[:strings.Index(line, ":")], " ") {
			line = strings.Replace(line, " ", "_", 1)
		}
		line = strings.Replace(line, ":", "", 1)
		_, err := fmt.Sscanf(line, "%s %d-%d or %d-%d", &r.name, &r.range1min, &r.range1max, &r.range2min, &r.range2max)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s: %d-%d or %d-%d\n", rules.range1min, rules.range1max, rules.range2min, rules.range2max)
		rulesCollection = append(rulesCollection, r)
		i++
	}
	return rulesCollection, i
}
