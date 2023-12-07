package d07

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	data := tools.ReadLines("test.txt")
	crabs := tools.ConvertCommaSeparatedStrIntoInts(data[0])
	assert.Len(t, crabs, 10)

	assert.Equal(t, float64(37), processLinearFuel(crabs))
	assert.Equal(t, float64(168), processNonLinearFuel(crabs))
}
