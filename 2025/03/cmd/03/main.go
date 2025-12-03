package main

import "fmt"
import "advent_03/internal"

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	totalVoltage := 0

	for scanner.Scan() {
		line := scanner.Text()
		battery := internal.NewBank(line)
		voltage := battery.OutputVoltage()
		totalVoltage += voltage

	}
	totalVoltage2 := 0

	scanner2, closer2 := internal.Scan("input2.txt")
	defer closer2.Close()
	for scanner2.Scan() {
		line := scanner2.Text()
		bigBattery := internal.NewBigBank(line)
		voltage := bigBattery.OutputVoltage(12)
		totalVoltage2 += voltage
	}

	fmt.Printf("Total Output Voltage: %d\n", totalVoltage)
	fmt.Printf("Total Output Voltage 2: %d\n", totalVoltage2)

}
