package d08

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"AdventOfCode/tools"
)

// a: 8     8888
// b: 6    6    8
// c: 8    6    8
// d: 7     7777
// e: 4    4    9
// f: 9    4    9
// g: 7     7777

var digitMap = map[int]int{
	467889:  0,
	89:      1,
	47788:   2,
	77889:   3,
	6789:    4,
	67789:   5,
	467789:  6,
	889:     7,
	4677889: 8,
	677889:  9,
}

func processA(rawData []string) int {
	sum := 0
	for _, line := range rawData {
		segments := strings.Split(line, "|")
		output := readSegments(segments[1], 4)
		sum += countKnownDigits(output)
	}
	return sum
}

func processB(rawData []string) int {
	sum := 0
	for _, line := range rawData {
		segments := strings.Split(line, "|")
		input := readSegments(segments[0], 10)
		output := readSegments(segments[1], 4)
		frequency := calculateFrequency(input)
		applyFrequency(output, frequency)
		sorted := softNumbers(output)
		result := mapNumbers(sorted)
		// fmt.Println(result)
		sum += result
	}
	return sum
}

func mapNumbers(arr []string) int {
	result := ""
	for _, d := range arr {
		if digit, ok := digitMap[tools.StrToInt(d)]; ok {
			result = result + fmt.Sprint(digit)
		} else {
			log.Fatal("no digit found")
		}
	}
	return tools.StrToInt(result)
}

func softNumbers(input []string) []string {
	var result = []string{}
	for _, entry := range input {
		var tmp []int
		for _, d := range entry {
			tmp = append(tmp, tools.StrToInt(string(d)))
		}
		sort.Ints(tmp)
		result = append(result, tools.ConvertIntSliceToString(tmp))
	}
	return result
}

func applyFrequency(output []string, frequency map[string]int) {
	for i := range output {
		for char, f := range frequency {
			output[i] = strings.ReplaceAll(output[i], char, fmt.Sprint(f))
		}
	}
}

func calculateFrequency(input []string) map[string]int {
	var frequency = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
	}
	for _, digit := range input {
		for i := 0; i < len(digit); i++ {
			frequency[string(digit[i])]++
		}
	}
	return frequency
}

func countKnownDigits(output []string) int {
	/*
		1: 2*
		4: 4*
		7: 3*
		8: 7*
	*/
	c := 0
	for _, s := range output {
		if len(s) == 2 || len(s) == 4 || len(s) == 3 || len(s) == 7 {
			c++
		}
	}
	return c
}

func readSegments(data string, n int) []string {
	input, err := tools.ReadWordsFromLine(strings.NewReader(data), n)
	if err != nil {
		log.Fatal(err)
	}
	if len(input) != n {
		log.Fatal("wrong number of digits read")
	}
	return input
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #8 A")
	data := tools.ReadFile("y2021/days/d08/data.txt")
	fmt.Printf("result: %v\n", processA(data))
	fmt.Println("DAY #8 B")
	fmt.Printf("result: %v\n", processB(data))
}
