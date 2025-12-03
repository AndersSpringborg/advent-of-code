package internal

import (
	"log"
	"strconv"
	"strings"
)

type Bank struct {
	batteries string
}

func NewBank(batteries string) *Bank {
	return &Bank{
		batteries: batteries,
	}
}

func (b *Bank) OutputVoltage() int {
	voltages := parseBatteriesToIntArray(b.batteries)
	left := 0
	right := 0

	for i, voltage := range voltages {
		if voltage > left && i < len(voltages)-1 {
			left = voltage
			right = voltages[i+1]
		} else if voltage > right {
			right = voltage
		}

	}

	num := b.joltFromJumbers(left, right)

	return num
}

func (b *Bank) joltFromJumbers(left int, right int) int {
	togehterNumber := strconv.Itoa(left) + strconv.Itoa(right)
	num, err := strconv.Atoi(togehterNumber)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseBatteriesToIntArray(input string) []int {
	voltagesAsString := strings.Split(input, "")
	voltages := make([]int, len(voltagesAsString))
	for i, v := range voltagesAsString {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		voltages[i] = num
	}

	return voltages
}
