package Day04

import (
	"fmt"
	"strings"

	"AdventOfCode/tools"
)

type passport map[string]string

type params struct {
	must  bool
	check checker
}

var fieldsDefinition = map[string]params{
	"byr": {true, NewChecker(digit(), length(4), span(1920, 2002))},
	"iyr": {true, NewChecker(digit(), length(4), span(2010, 2020))},
	"eyr": {true, NewChecker(digit(), length(4), span(2020, 2030))},
	"hgt": {true, NewChecker(height())},
	"hcl": {true, NewChecker(hairColor())},
	"ecl": {true, NewChecker(eyeColor())},
	"pid": {true, NewChecker(digit(), length(9))},
	// "cid": false,
}

func Goooo() {
	fmt.Println("--------- DAY 04 ---------")

	// lines := tools.ReadFile(("y2020/days/Day04/testInput.txt"))
	lines := tools.ReadFile(("y2020/days/Day04/input.txt"))

	count := 0
	passports := tools.CombineLines(lines)
	for _, line := range passports {
		ok := true
		for field, param := range fieldsDefinition {
			if !param.must {
				continue
			}
			if !strings.Contains(line, field+":") {
				ok = false
				break
			}
		}
		if ok {
			count += 1
			// fmt.Println(line)
		}
	}

	fmt.Println(count)
	goooo2()
}

func goooo2() {
	fmt.Println("--------- DAY 04 b---------")
	lines := tools.ReadFile(("y2020/days/Day04/input.txt"))
	// lines := tools.ReadFile(("y2020/days/Day04/testInputb.txt"))
	count := 0
	passLines := tools.CombineLines(lines)
	for _, line := range passLines {
		pass := lineToPassword(line)
		if checkFields(pass) {
			// fmt.Println(line)
			count += 1
		}
	}
	fmt.Println(count)

}

func checkFields(pass passport) bool {
	for defKey, defParam := range fieldsDefinition {
		if !defParam.must {
			continue
		}
		passFieldValue, ok := pass[defKey]
		if !ok {
			return false
		}
		if !defParam.check(passFieldValue) {
			return false
		}
	}
	return true
}

func lineToPassword(line string) passport {
	pass := make(passport)
	lineParts := strings.Split(line, " ")
	for _, fullField := range lineParts {
		fieldParts := strings.Split(fullField, ":")
		pass[fieldParts[0]] = fieldParts[1]
	}
	return pass
}
