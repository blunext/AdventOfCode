package d01

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	depths := tools.GetSliceOfInts("test.txt")
	assert.Equal(t, 7, countDepths(depths))

	transformed := transformMeasurements(depths)
	assert.Equal(t, 5, countDepths(transformed))
}
