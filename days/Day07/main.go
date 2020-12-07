package Day07

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	quantity int
}
type bagsHolder struct {
	bags map[string][]bag
}

func newBagHolder() *bagsHolder {
	return &bagsHolder{bags: make(map[string][]bag)}
}

func (h *bagsHolder) addInnerBag(outerBag, innerBag string, quantity int) {
	bag := bag{name: innerBag, quantity: quantity}
	h.bags[outerBag] = append(h.bags[outerBag], bag)
}

func (h *bagsHolder) countBagsHolding(name string) int {
	i := 0
	for holderName, _ := range h.bags {
		//fmt.Printf("%v: ", holderName)
		if h.check(holderName, name) {
			i++
		}
		//fmt.Println()
	}
	return i
}

func (h *bagsHolder) check(holder, name string) bool {
	bag, _ := h.bags[holder]
	for _, innerBag := range bag {
		//fmt.Printf(" %v ", bagName)

		if innerBag.name == name {
			//fmt.Printf(" ### OK ###")
			return true
		}
		//fmt.Printf(" >> ")
		if h.check(innerBag.name, name) {
			return true
		}
	}
	return false
}

func Goooo() {
	fmt.Println("--------- DAY 07 ---------")
	//lines := tools.ReadFile(("days/Day07/testinput.txt"))
	lines := tools.ReadFile(("days/Day07/input.txt"))

	holder := populateData(lines)
	fmt.Printf("count part a=%d\n", holder.countBagsHolding("shiny gold"))

	lines = tools.ReadFile(("days/Day07/testinput2.txt"))
	holder = populateData(lines)

}

func populateData(lines []string) *bagsHolder {
	holder := newBagHolder()
	for _, line := range lines {
		lineWithoutEndingDot := line[:len(line)-1]

		bagData := strings.Split(lineWithoutEndingDot, "contain")
		outerBag := clean(bagData[0])
		//holder.addOuterBag(outerBag)

		contains := strings.Split(bagData[1], ",")
		for _, innerBag := range contains {
			b := clean(innerBag)
			if b == "no other" {
				// was "no other bags" but bags word is removed
				continue
			}
			if b[1:2] != " " {
				// expected number as first byte, second is space
				panic(fmt.Sprintf("rule broken - first byte number, second byte space : %v", b))
			}
			//bag = bag.add(b[2:])
			quantity, _ := strconv.Atoi(b[:1])
			holder.addInnerBag(outerBag, b[2:], int(quantity))
		}
	}
	return holder
}

func clean(innerBag string) string {
	b := strings.TrimSuffix(innerBag, ".")
	b = strings.TrimSpace(b)
	b = strings.Replace(b, "bags", "", 1)
	b = strings.Replace(b, "bag", "", 1)
	b = strings.TrimSpace(b) // again
	return b
}
