package d06

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	data := tools.ReadFile("test.txt")
	fishes := tools.ConvertCommaSeparatedStrIntoInts8(data[0])
	assert.Len(t, fishes, 5)

	assert.Equal(t, 26, process(fishes, 18))
	assert.Equal(t, 5934, process(fishes, 80))
	assert.Equal(t, 26984457539, process(fishes, 256))

}
