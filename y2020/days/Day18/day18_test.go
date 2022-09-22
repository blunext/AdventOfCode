package Day18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData1(t *testing.T) {
	in := "2 * 3 + (4 * 5)"
	tokens := preProcess(in)
	result, _ := process(tokens)
	assert.Equal(t, 26, result, fmt.Sprintf("error in %v", in))

	in = "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	tokens = preProcess(in)
	result, _ = process(tokens)
	assert.Equal(t, 437, result, fmt.Sprintf("error in %v", in))

	in = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	tokens = preProcess(in)
	result, _ = process(tokens)
	assert.Equal(t, 12240, result, fmt.Sprintf("error in %v", in))

	in = "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	tokens = preProcess(in)
	result, _ = process(tokens)
	assert.Equal(t, 13632, result, fmt.Sprintf("error in %v", in))

}
