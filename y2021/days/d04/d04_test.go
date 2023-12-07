package d04

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestDataA(t *testing.T) {
	data := tools.ReadLines("test.txt")
	numbers := tools.ConvertCommaSeparatedStrIntoInts(data[0])
	assert.Len(t, numbers, 27)
	boards := populateBoards(data[1:])

	assert.Len(t, boards, 3)

	assert.Equal(t, 4512, game(boards, numbers, true))
	boards = populateBoards(data[1:])
	assert.Equal(t, 1924, game(boards, numbers, false))
}
