package Day07

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type bag struct {
	name     string
	contains *bag
}

func newBag(name string) *bag {
	return &bag{name: name}
}

func (b *bag) find(name string) bool {
	switch {
	case b.name == name:
		return true
	case b.contains == nil:
		return false
	default:
		return b.contains.find(name)
	}
}

func (b *bag) add(name string) *bag {
	bag := newBag(name)
	b.contains = bag
	return bag
}

type bagsHolder struct {
	bags map[string]*bag
}

func newBagHolder() *bagsHolder {
	return &bagsHolder{bags: make(map[string]*bag)}
}
func (h *bagsHolder) addOuterBag(name string) *bag {
	if found, ok := h.bags[name]; ok {
		panic("outter bag found")
		return found
	}
	bag := newBag(name)
	h.bags[name] = bag
	return bag
}

func (h *bagsHolder) countBagsHolding(name string) int {
	i := 0
	for _, bag := range h.bags {
		if bag.find(name) {
			i++
		}
	}
	return i
}

func Goooo() {
	fmt.Println("--------- DAY 07 ---------")
	lines := tools.ReadFile(("days/Day07/testinput.txt"))
	//lines := tools.ReadFile(("days/Day07/input.txt"))

	holder := newBagHolder()
	for _, line := range lines {
		lineWithoutEndingDot := line[:len(line)-1]

		bagData := strings.Split(lineWithoutEndingDot, "contain")
		bag := holder.addOuterBag(clean(bagData[0]))

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
			bag = bag.add(b[2:])
		}
	}

	fmt.Printf("count=%d\n", holder.countBagsHolding("shiny gold"))
}

func clean(innerBag string) string {
	b := strings.TrimSuffix(innerBag, ".")
	b = strings.TrimSpace(b)
	b = strings.Replace(b, "bags", "", 1)
	b = strings.Replace(b, "bag", "", 1)
	b = strings.TrimSpace(b) // again
	return b
}
