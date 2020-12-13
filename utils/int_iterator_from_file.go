package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func IntIteratorFromFile(filename string) <-chan int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ch := make(chan int)
	go func(inputFile *os.File) {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			inputIntAsInt64, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			ch <- int(inputIntAsInt64)
		}
		close(ch)
	}(file)
	time.Sleep(100 * time.Millisecond)
	return ch
}
