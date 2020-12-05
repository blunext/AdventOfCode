package Day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData41(t *testing.T) {

	number := getNumber("FBFBBFF")
	assert.Equal(t, 44, number, "not 44")

	number = getNumber("RLR")
	assert.Equal(t, 5, number, "not 5")
}

func TestData42(t *testing.T) {

	pass := "BFFFBBFRRR"
	seatId := getSeatId(pass)
	assert.Equal(t, 567, seatId, "err TestData42")

	pass = "FFFBBBFRRR"
	seatId = getSeatId(pass)
	assert.Equal(t, 119, seatId, "err TestData42")

	pass = "BBFFBBFRLL"
	seatId = getSeatId(pass)
	assert.Equal(t, 820, seatId, "err TestData42")
}
