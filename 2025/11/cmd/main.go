package main

import (
	"advent_11/internal"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	sg := internal.NewServerGraph()

	for scanner.Scan() {
		line := scanner.Text()
		sg.ParseLine(line)
	}

	totalDevices := sg.TotalPathsFromYouToOut()
	println("Total devices connected from 'YOU' to 'OUT':", totalDevices)
	totalDevicesDcc := sg.TotalPathsFromYouToOutWithFCCandDAC()
	println("Total devices connected from 'YOU' to 'OUT' with FCC and DAC:", totalDevicesDcc)

}
