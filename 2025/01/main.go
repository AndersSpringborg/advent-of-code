package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	scanner, closer := Scan("input.txt")
	defer closer.Close()

	dial := NewDial(50)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("current Dial: %d\n", dial.Index)
		fmt.Printf("Line is: %s\n", line)
		counter++
		rotationSymbol, steps := line[0:1], line[1:]
		stepsNumber, err := strconv.Atoi(steps)
		uintStempNumber := uint(stepsNumber)
		if err != nil {
			log.Fatalf("Failed to convert steps to number: %v", err)
		}

		if rotationSymbol == "R" {
			dial.Rotate(Right, uintStempNumber)
		} else if rotationSymbol == "L" {
			dial.Rotate(Left, uintStempNumber)
		}
	}

	log.Printf("Final Dial Index: %d", dial.Index)
	log.Printf("Answer is : %d", dial.ZeroIndex)
}
