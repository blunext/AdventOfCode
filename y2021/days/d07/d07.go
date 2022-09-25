package d07

import (
	"fmt"
	"math"
	"sort"
	"time"

	"gonum.org/v1/gonum/stat"

	"AdventOfCode/tools"
)

func processLinearFuel(crabs []int) float64 {

	floats := tools.ConvertIntSliceToFloat64(crabs)
	sort.Float64s(floats)
	median := stat.Quantile(0.5, stat.Empirical, floats, nil)
	sum := distance(floats, median, func(f float64) float64 {
		return f
	})
	fmt.Println("sum: ", sum)
	return sum
}

func distance(data []float64, from float64, factor func(float64) float64) float64 {
	sum := float64(0)
	for _, crab := range data {
		sum += factor(math.Abs(from - crab))
	}
	return sum
}

func processNonLinearFuel(crabs []int) float64 {
	floats := tools.ConvertIntSliceToFloat64(crabs)
	mean := stat.Mean(floats, nil)

	fuelFunc := func(distance float64) float64 {
		return distance * (distance + 1) / 2
	}

	sumMin := distance(floats, math.Floor(mean), fuelFunc)
	sumMax := distance(floats, math.Ceil(mean), fuelFunc)

	return math.Min(sumMin, sumMax)
}

func Main() {
	timeStart := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(timeStart))
	}()
	fmt.Println("DAY #7 A")
	data := tools.ReadFile("y2021/days/d07/data.txt")

	crabs := tools.ConvertCommaSeparatedStrIntoInts(data[0])
	fmt.Printf("result: %v\n", processLinearFuel(crabs))

	fmt.Println("DAY #7 B")
	crabs = tools.ConvertCommaSeparatedStrIntoInts(data[0])
	fmt.Printf("result: %v\n", int(processNonLinearFuel(crabs)))
}
