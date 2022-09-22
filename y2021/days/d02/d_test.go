package d02

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	data := tools.ReadByWords("test.txt", 2)
	assert.Equal(t, 6, len(data))
	assert.Equal(t, 150, process(data))
}
