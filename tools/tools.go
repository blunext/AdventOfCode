package tools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetSliceOfInts(path string) []int {
	return ConvertStrArrayIntoInts(ReadFile(path))
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines
}

func ReadByWords(path string, len int) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines [][]string

	for scanner.Scan() {
		r := strings.NewReader(scanner.Text())
		words, err := ReadWordsFromLine(r, len)
		if err != nil {
			fmt.Printf("Fscanf err: %v\n", err)
		}
		var lineArr []string
		for _, t := range words {
			lineArr = append(lineArr, t)
		}
		txtlines = append(txtlines, lineArr)
	}
	return txtlines
}

func ReadWordsFromLine(r io.Reader, n int) ([]string, error) {
	in := make([]string, n)
	for i := range in {
		_, err := fmt.Fscan(r, &in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func ConvertStrArrayIntoInts(lines []string) []int {
	ints := []int{}
	for _, s := range lines {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func parseBits(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 16)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ConvertBitsIntoInt64(lines []string) []int64 {
	var ints []int64
	for _, s := range lines {
		ints = append(ints, parseBits(s))
	}
	return ints
}

func CombineLines(lines []string) []string {
	set := []string{}
	new := true
	for _, line := range lines {
		if line == "" {
			new = true
			continue
		}
		if new {
			set = append(set, line)
			new = false
			continue
		}
		set[len(set)-1] = set[len(set)-1] + " " + line
	}
	return set
}

func GetBit(val uint64, index int) uint64 {
	mask := uint64(1) << index
	return (val & mask) >> index
}

func SetBit(val uint64, index int) uint64 {
	return val | (1 << index)
}

func ConvertCommaSeparatedStrIntoInts(line string) []int {
	ints := []int{}
	for _, s := range strings.Split(line, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ConvertStrOfDigitsToInts(line string) []int {
	ints := []int{}
	for _, character := range line {
		i, err := strconv.Atoi(string(character))
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ConvertCommaSeparatedStrIntoInts8(line string) []int8 {
	var ints []int8
	for _, s := range strings.Split(line, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, int8(i))
	}
	return ints
}

func ConvertIntSliceToString(ints []int) string {
	var result string
	for _, s := range ints {
		result = result + strconv.Itoa(s)
	}
	return result
}
func ConvertIntSliceToFloat64(ints []int) []float64 {
	var result []float64
	for i := range ints {
		result = append(result, float64(ints[i]))
	}
	return result
}

func RemoveSlice(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("String conv to int err: ", err)
	}
	return i
}
