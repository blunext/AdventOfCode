package d03

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestDataA(t *testing.T) {
	data := tools.ConvertBitsIntoInt64(tools.ReadFile("test.txt"))
	assert.Equal(t, 12, len(data))
	assert.Equal(t, int64(22), process(data, 5, false))
	assert.Equal(t, int64(9), process(data, 5, true))

	assert.Equal(t, int64(23), processWithReduction(data, 5, false))
	assert.Equal(t, int64(10), processWithReduction(data, 5, true))

}
