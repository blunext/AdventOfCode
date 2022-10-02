package d10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	data := tools.ReadFile("test.txt")
	assert.Equal(t, 26397, process(data))
}
