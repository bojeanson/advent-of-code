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

	ch := make(chan int)
	go func(inputFile *os.File) {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			inputIntAsInt64, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			ch <- int(inputIntAsInt64)
		}
		defer file.Close()
		close(ch)
	}(file)
	time.Sleep(100 * time.Millisecond)
	return ch
}

func StringIteratorFromFile(filename string) <-chan string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string)
	go func(inputFile *os.File) {
		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
		defer file.Close()
	}(file)
	time.Sleep(100 * time.Millisecond)
	return ch
}
