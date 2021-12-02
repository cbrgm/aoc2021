package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	depths, err := readInputFile(f)
	if err != nil {
		panic(err)
	}

	var increased int
	var decreased int

	for i, v := range depths {
		if i == 0 {
			continue
		}
		if v < depths[i-1] {
			decreased++
		} else {
			increased++
		}
	}

	fmt.Printf("Increased measurements: %v \n", increased)
	fmt.Printf("Decreased measurements: %v \n", decreased)
}

func readInputFile(file *os.File) ([]int, error) {
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
