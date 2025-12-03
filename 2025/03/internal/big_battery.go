package internal

import (
	"log"
	"strconv"
)

type BigBank struct {
	batteries string
}

func NewBigBank(batteries string) *BigBank {
	return &BigBank{
		batteries: batteries,
	}
}

func (b *BigBank) OutputVoltage(n int) int {
	batteries := make([]int, n)

	voltages := parseBatteriesToIntArray(b.batteries)

	for i, _ := range voltages {
		if i+n == len(voltages)+1 {
			break
		}
		for j := 0; j < n; j++ {
			index := i + j
			if voltages[index] > batteries[j] {
				batteries[j] = voltages[index]
				// move the rest down to this
				for k := j + 1; k < n; k++ {
					k_index := i + k
					batteries[k] = voltages[k_index]
				}
			}
		}

	}

	num := b.voltageFromJumbers(batteries...)

	return num
}

func (b *BigBank) voltageFromJumbers(nums ...int) int {
	var combined string
	for _, num := range nums {
		combined += strconv.Itoa(num)
	}

	result, err := strconv.Atoi(combined)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
