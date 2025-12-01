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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("current Dial: %d\n", dial.Index)
		fmt.Printf("Line is: %s\n", line)
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

	scanner2, closer2 := Scan("input2.txt")
	defer closer2.Close()

	dialMethod := NewDial0x434C49434B(50)
	for scanner2.Scan() {
		line := scanner2.Text()
		fmt.Printf("current Dial: %d\n", dialMethod.Index)
		fmt.Printf("Line is: %s\n", line)
		rotationSymbol, steps := line[0:1], line[1:]
		stepsNumber, err := strconv.Atoi(steps)
		if err != nil {
			log.Fatalf("Failed to convert steps to number: %v", err)
		}

		if rotationSymbol == "R" {
			dialMethod.Rotate(Right, stepsNumber)
		} else if rotationSymbol == "L" {
			dialMethod.Rotate(Left, stepsNumber)
		}
	}

	log.Printf("Final Dial Index: %d", dial.Index)
	log.Printf("Answer is : %d", dial.ZeroIndex)
	log.Printf("Final Dial0x434C49434B Index: %d", dialMethod.Index)
	log.Printf("Answer 2 is : %d", dialMethod.ZeroIndex)
}
