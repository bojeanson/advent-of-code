package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func IntIteratorFromFile(filename string) <-chan int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ch := make(chan int)
	go func() {
		for scanner.Scan() {
			inputIntAsInt64, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			ch <- int(inputIntAsInt64)
		}
		close(ch)
	}()
	return ch
}
