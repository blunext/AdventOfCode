package Day04

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type check func(s string) bool

var fields = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	"cid": false,
}

func Goooo() {
	fmt.Println("--------- DAY 04 ---------")

	//lines := tools.ReadFile(("days/Day04/testInput.txt"))
	lines := tools.ReadFile(("days/Day04/input.txt"))

	count := 0
	passports := processPassData(lines)
	for _, line := range passports {
		ok := true
		for field, must := range fields {
			if !must {
				continue
			}
			if !strings.Contains(line, field+":") {
				ok = false
				break
			}
		}
		if ok {
			count += 1
			//fmt.Println(line)
		}
	}

	fmt.Println(count)
	goooo2()
}

func goooo2() {
	fmt.Println("--------- DAY 04 b---------")
	//lines := tools.ReadFile(("days/Day04/input.txt"))

}

func processPassData(lines []string) []string {
	passports := []string{}
	new := true
	for _, line := range lines {
		if line == "" {
			new = true
			continue
		}
		if new {
			passports = append(passports, line)
			new = false
			continue
		}
		passports[len(passports)-1] = passports[len(passports)-1] + " " + line
	}
	return passports
}
