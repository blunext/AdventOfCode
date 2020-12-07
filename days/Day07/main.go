package Day07

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"strings"
)

type bag struct {
	name     string
	contains []*bag
}

func newBag(name string) *bag {
	return &bag{name: name}
}

func (b *bag) find(bag *bag) bool {
	for _, bagItem := range b.contains {
		switch bagItem {
		case bag:
			//fmt.Printf("%v !!!! ", b.name)
			return true
		default:
			//fmt.Printf("%v > ", b.name)
			return bagItem.find(bag)
		}
	}
	return false
}

type bagsHolder struct {
	bags map[string]*bag
}

func newBagHolder() *bagsHolder {
	return &bagsHolder{bags: make(map[string]*bag)}
}
func (b *bagsHolder) addOuterBag(name string) *bag {
	if found, ok := b.bags[name]; ok {
		return found
	}
	bag := newBag(name)
	b.bags[name] = bag
	return bag
}

func (b *bagsHolder) addBag(bag *bag, name string) *bag {
	if found, ok := b.bags[name]; ok {
		for _, has := range bag.contains {
			if has == found {
				return found
			}
		}
		bag.contains = append(bag.contains, found)
		return found
	}
	newBag := newBag(name)
	b.bags[name] = newBag
	bag.contains = append(bag.contains, newBag)
	return newBag
}

func (b *bagsHolder) countBagsHolding(name string) int {
	i := 0
	if bagSearched, ok := b.bags[name]; ok {
		//fmt.Printf("%v > ", bagSearched.name)
		for _, bag := range b.bags {
			if bag.find(bagSearched) {
				i++
			}
			//fmt.Println()
		}
		return i
	}
	panic("cannot find bag " + name)
}

func Goooo() {
	fmt.Println("--------- DAY 07 ---------")
	//lines := tools.ReadFile(("days/Day07/testinput.txt"))
	lines := tools.ReadFile(("days/Day07/input.txt"))

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
			bag = holder.addBag(bag, (b[2:]))
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
