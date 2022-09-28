package d09

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	data := tools.ReadFile("test.txt")
	width := 10
	assert.Equal(t, width, len(data[0]))
	matrix := prepareData(data)
	p9 := point{val: 9}
	assert.Equal(t, p9, matrix[0][0])
	assert.Equal(t, p9, matrix[0][5])
	assert.Equal(t, p9, matrix[0][width+1])
	assert.Equal(t, p9, matrix[1][0])
	assert.Equal(t, p9, matrix[1][width+1])
	assert.Equal(t, p9, matrix[3][0])
	assert.Equal(t, p9, matrix[3][width+1])

	process(matrix)
	assert.Equal(t, 15, riskPoints(matrix))

	makeRain(matrix)
	basins := countBasinSizes(matrix)
	assert.Len(t, basins, 4)

	assert.Equal(t, 1134, threeLargestBasinMultiplied(basins))

}
