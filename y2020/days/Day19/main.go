package Day19

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/tools"
)

type rule struct {
	set  [][]int
	base string
}

func Goooo() {
	fmt.Println("--------- DAY 19 ---------")
	// lines := tools.ReadLines(("y2020/days/Day19/testInput.txt"))
	lines := tools.ReadLines(("y2020/days/Day19/input.txt"))

	rules, index := initRules(lines)

	i := checkRules(index, lines, rules)

	fmt.Printf("part 1: %d", i)

}

func checkRules(index int, lines []string, rules map[int]rule) int {
	no := 0
	index++
	for ; index < len(lines); index++ {
		// fmt.Printf("########## check %s\n", lines[index])
		if ok, i := check(lines[index], rules, 0, 0); ok {
			if i == len(lines[index]) {
				// fmt.Println(lines[index], i)
				no++
			}
		}
	}
	return no
}

func check(s string, rules map[int]rule, ruleNo int, index int) (bool, int) {
	// fmt.Printf("ruleNo: %d, letter: %s\n", ruleNo, s[index:index+1])
	// if index == len(s)-1 {
	//	return true, 0
	// }
	shift := 0
	r := rules[ruleNo]
	if r.base == s[index:index+1] {
		// fmt.Printf("%s - ok\n", s[index:index+1])
		return true, 1
	}

	for _, ruleSet := range r.set {
		i := 0
		ok := true
		for _, no := range ruleSet {
			if ok, shift = check(s, rules, no, index+i); !ok {
				// ok = false
				break
			}
			i += shift
		}
		if ok {
			// fmt.Printf("%s - ok\n", s[index:index+1])
			return true, i
		}
	}
	// fmt.Printf("%s - no ok\n", s[index:index+1])
	return false, 0
}

func initRules(lines []string) (map[int]rule, int) {
	rules := make(map[int]rule)
	index := 0
	line := ""
	for index, line = range lines {
		if line == "" {
			break
		}
		t := strings.Split(line, ":")
		no, _ := strconv.Atoi(t[0])

		if strings.Contains(t[1], "\"") {
			var base string
			_, err := fmt.Sscanf(strings.ReplaceAll(t[1], "\"", ""), " %s", &base)
			if err == nil {
				rules[no] = rule{base: base}
				continue
			}
		}

		rule := rule{}
		or := strings.Split(t[1], "|")
		for _, ruleSet := range or {
			set := strings.Split(strings.TrimSpace(ruleSet), " ")
			var andRules []int
			for _, s := range set {
				ruleNo, _ := strconv.Atoi(s)
				andRules = append(andRules, ruleNo)
			}
			if len(andRules) > 0 {
				rule.set = append(rule.set, andRules)
			}
		}
		rules[no] = rule
	}
	return rules, index
}
