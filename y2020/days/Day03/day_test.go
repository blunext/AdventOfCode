package Day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData1(t *testing.T) {
	lines := tools.ReadLines(("testInput.txt"))
	for _, s := range slopes {
		no := check(lines, s.x, s.y)
		assert.Equal(t, s.testNo, no, fmt.Sprintf("%d, %d = %d ERR", s.x, s.y, no))
	}
}
