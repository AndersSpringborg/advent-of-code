package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigBank(t *testing.T) {
	testCases := []struct {
		bank            BigBank
		expectedVoltage int
	}{
		{BigBank{"811111111111119"}, 811111111119},
		{BigBank{"234234234234278"}, 434234234278},
		{BigBank{"818181911112111"}, 888911112111},
	}

	for _, tc := range testCases {
		result := tc.bank.OutputVoltage(12)
		assert.Equal(t, tc.expectedVoltage, result)
	}
}
