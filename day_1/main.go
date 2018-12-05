package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func frequency(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(scanner.Text() + " is not a integer")
		}
		result += i
	}

	return result
}

func frequencyTwice(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputFreq []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(scanner.Text() + " is not a integer")
		}
		inputFreq = append(inputFreq, i)
	}

	prevFreqs := make(map[int]bool)
	prevFreqs[0] = true

	var currFreq int
	for i := 0; i%len(inputFreq) < len(inputFreq); i++ {
		currFreq += inputFreq[i%len(inputFreq)]

		if prevFreqs[currFreq] {
			return currFreq
		}

		prevFreqs[currFreq] = true
	}

	var result int
	return result
}

func main() {
	fmt.Println(frequency("input.txt"))
	fmt.Println(frequencyTwice("input.txt"))
}
