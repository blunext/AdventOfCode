package Day02

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	rules string
	valid bool
}

var testData1 = []test{
	{"1-3 a: abcde", true},
	{"1-3 b: cdefg", false},
	{"2-9 c: ccccccccc", true},
}

var testData2 = []test{
	{"1-3 a: abcde", true},
	{"1-3 b: cdefg", false},
	{"2-9 c: ccccccccc", false},
}

func TestData1(t *testing.T) {
	for _, v := range testData1 {
		if v.valid {
			assert.True(t, check1(v.rules), fmt.Sprintf("error in rule %v", v.rules))
			continue
		}
		assert.False(t, check1(v.rules), fmt.Sprintf("error in rule %v", v.rules))
	}
}

func TestData2(t *testing.T) {
	for _, v := range testData2 {
		if v.valid {
			assert.True(t, check2(v.rules), fmt.Sprintf("error in rule %v", v.rules))
			continue
		}
		assert.False(t, check2(v.rules), fmt.Sprintf("error in rule %v", v.rules))
	}
}
