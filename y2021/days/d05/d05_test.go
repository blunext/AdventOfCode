package d05

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestDataA(t *testing.T) {
	data := tools.ReadLines("test.txt")
	vectors := getVectors(data)
	assert.Len(t, vectors, 10)

	v := vector{start: point{x: 1, y: 5}, end: point{1, 0}}
	assert.Len(t, getPath(v, false), 6)

	v = vector{start: point{x: 1, y: 1}, end: point{1, 1}}
	assert.Len(t, getPath(v, false), 1)

	v = vector{start: point{x: 0, y: 1}, end: point{10, 1}}
	assert.Len(t, getPath(v, false), 11)

	board := mapVectors(vectors, false)

	assert.Equal(t, 5, countOverlaps(board))

	v = vector{start: point{x: 0, y: 0}, end: point{10, 10}}
	assert.True(t, diagonalVector(v))
	v.end.y = 5
	assert.False(t, diagonalVector(v))

	v = vector{start: point{x: 10, y: 10}, end: point{5, 5}}
	assert.True(t, diagonalVector(v))

	v = vector{start: point{x: 9, y: 7}, end: point{7, 9}}
	assert.True(t, diagonalVector(v))

	board = mapVectors(vectors, true)
	assert.Equal(t, 12, countOverlaps(board))
}
