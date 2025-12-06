package internal

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Scan(path string) (*bufio.Scanner, io.Closer) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner, file
}
