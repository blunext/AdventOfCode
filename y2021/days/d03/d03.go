package d03

import (
	"fmt"
	"strconv"

	"AdventOfCode/tools"
)

func process(data []int64, bits int, reverseCheck bool) int64 {
	result := ""

	for i := bits - 1; i >= 0; i-- {
		s0, s1 := 0, 0
		for _, num := range data {
			b := tools.GetBit(uint64(num), i)
			switch b {
			case 0:
				s0++
			case 1:
				s1++
			}
		}
		switch reverseCheck {
		case true:
			if s0 < s1 {
				result = result + "0"
			} else {
				result = result + "1"
			}
		default:
			if s0 > s1 {
				result = result + "0"
			} else {
				result = result + "1"
			}
		}
	}
	base10, _ := strconv.ParseInt(result, 2, bits+1)
	// fmt.Println(result, base10)
	return base10
}

func Main() {
	data := tools.ConvertBitsIntoInt64(tools.ReadFile("y2021/days/d03/data.txt"))
	fmt.Println("DAY #3 A")
	gamma := process(data, 12, false)
	epsilon := process(data, 12, true)
	fmt.Printf("result: %d\n", gamma*epsilon)
}
