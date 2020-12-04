package Day03

import (
	"Go-AdventOfCode2020/tools"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData1(t *testing.T) {
	lines := tools.ReadFile(("testInput.txt"))
	for _, s := range slopes {
		no := check(lines, s.x, s.y)
		assert.Equal(t, s.testNo, no, fmt.Sprintf("%d, %d = %d ERR", s.x, s.y, no))
	}
}

func TestData2(t *testing.T) {

}
