package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBank(t *testing.T) {
	testCases := []struct {
		bank            Bank
		expectedVoltage int
	}{
		{Bank{"99"}, 99},
		{Bank{"987654321111111"}, 98},
		{Bank{"811111111111119"}, 89},
		{Bank{"234234234234278"}, 78},
		{Bank{"818181911112111"}, 92},
	}

	for _, tc := range testCases {
		result := tc.bank.OutputVoltage()
		assert.Equal(t, tc.expectedVoltage, result)
	}
}
