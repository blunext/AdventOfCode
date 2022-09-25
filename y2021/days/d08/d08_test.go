package d08

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"AdventOfCode/tools"
)

func TestData(t *testing.T) {
	/*
		0: 6
		1: 2*
		2: 5
		3: 5
		4: 4*
		5: 5
		6: 6
		7: 3*
		8: 7*
		9: 6
	*/
	data := tools.ReadFile("test.txt")
	assert.Equal(t, 26, processA(data))
	assert.Equal(t, 61229, processB(data))

}
