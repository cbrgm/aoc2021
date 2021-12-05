package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Num   int
	Bingo bool
}

func loadMatrix(scanner *bufio.Scanner) ([][]BingoNumber, error) {
	matrix := make([][]BingoNumber, 0)
	for !lineIsBlank(scanner.Text()) {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		numbers := make([]BingoNumber, 0)

		for _, n := range lineParts {
			if n == "" {
				continue
			}
			number, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, BingoNumber{Num: number})
		}
		matrix = append(matrix, numbers)
		scanner.Scan()
	}
	return matrix, nil
}

func lineIsBlank(s string) bool {
	return s == ""
}

func loadInput(path string) ([]int, [][][]BingoNumber, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	lineParts := strings.Split(line, ",")

	bingoNumbers := make([]int, 0)
	for _, part := range lineParts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, nil, err
		}
		bingoNumbers = append(bingoNumbers, number)
	}

	tables := make([][][]BingoNumber, 0)
	scanner.Scan()
	for scanner.Scan() {
		m, err := loadMatrix(scanner)
		if err != nil {
			return nil, nil, err
		}
		tables = append(tables, m)
	}
	return bingoNumbers, tables, nil
}

func checkWinner(table [][]BingoNumber, i, j int) bool {
	rows := true
	for z := 0; z < len(table[0]); z++ {
		if !table[i][z].Bingo {
			rows = false
			break
		}
	}
	if rows {
		return true
	}

	cols := true
	for z := 0; z < len(table); z++ {
		if !table[z][j].Bingo {
			cols = false
			break
		}
	}

	return cols
}

func playout(table [][]BingoNumber, bingoNumber int) bool {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			if bingoNumber == table[i][j].Num {
				table[i][j].Bingo = true
				if checkWinner(table, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func findWinner(bingoNumbers []int, tables [][][]BingoNumber) ([][]BingoNumber, int) {
	for _, bingoNumber := range bingoNumbers {
		for _, table := range tables {
			if playout(table, bingoNumber) {
				return table, bingoNumber
			}
		}
	}
	return nil, -1
}

func main() {
	bingoNumbers, tables, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	winner, bingoNumber := findWinner(bingoNumbers, tables)
	sum := 0
	for _, row := range winner {
		for _, num := range row {
			if !num.Bingo {
				sum += num.Num
			}
		}
	}
	fmt.Printf("Result: %v", bingoNumber * sum)
}
