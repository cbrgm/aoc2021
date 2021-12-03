package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	// task 01

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	matrix := readInputFile(f)

	println(len(matrix))
	println(len(matrix[0]))

	gamma := calculateGamma(matrix)
	println(binaryToDecimal(gamma))

	epsilon := calculateEpsilon(matrix)
	println(binaryToDecimal(epsilon))

	// final answer
	res := binaryToDecimal(gamma) * binaryToDecimal(epsilon)
	println(res)

	// task 02

	oxygenRating := calcOxygenGeneratorRating(matrix)
	println(binaryToDecimal(oxygenRating))

	co2ScrubberRating := calcCO2ScrubberRating(matrix)
	println(binaryToDecimal(co2ScrubberRating))

	// final answer
	res = binaryToDecimal(oxygenRating) * binaryToDecimal(co2ScrubberRating)
	println(res)

}

func readInputFile(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var matrix [][]string
	for scanner.Scan() {
		matrix = append(matrix, parseLineFromString(scanner.Text()))
	}
	return matrix
}

func parseLineFromString(s string) []string {
	return strings.Split(s, "")
}

func binaryToDecimal(s string) int64 {
	if i, err := strconv.ParseInt(s, 2, 64); err != nil {
		panic("failed to parse int from string")
	} else {
		return i
	}
}

func calculateEpsilon(matrix [][]string) string {
	res := ""
	for x := 0; x < len(matrix[0]); x++ {
		column := []string{}
		for y, _ := range matrix {
			column = append(column, matrix[y][x])
		}

		res += leastCommonBit(column)
	}
	return res
}

func calculateGamma(matrix [][]string) string {
	res := ""
	for x := 0; x < len(matrix[0]); x++ {
		column := []string{}
		for y, _ := range matrix {
			column = append(column, matrix[y][x])
		}

		res += mostCommonBit(column)
	}
	return res
}

func calcOxygenGeneratorRating(matrix [][]string) string {
	res := ""
	initialLength := len(matrix)
	for pos := 0; pos < initialLength; pos++ {

		column := []string{}
		for row, _ := range matrix {
			column = append(column, matrix[row][pos])
		}

		if mostCommonBit(column) == "1" {
			matrix = keepMatrixRowsOnes(pos, matrix)
		} else {
			matrix = keepMatrixRowsZeros(pos, matrix)
		}

		// if there's one row left, return
		if len(matrix) == 1 {
			res = strings.Join(matrix[0], "")
			break
		}
	}
	return res
}

func keepMatrixRowsZeros(pos int, matrix [][]string) [][]string {
	newMatrix := [][]string{}
	for row, _ := range matrix {
		if matrix[row][pos] == "0" {
			newMatrix = append(newMatrix, matrix[row])
		}
	}
	return newMatrix
}

func keepMatrixRowsOnes(pos int, matrix [][]string) [][]string {
	newMatrix := [][]string{}
	for row, _ := range matrix {
		if matrix[row][pos] == "1" {
			newMatrix = append(newMatrix, matrix[row])
		}
	}
	return newMatrix
}

func calcCO2ScrubberRating(matrix [][]string) string {
	res := ""
	initialLength := len(matrix)
	for pos := 0; pos < initialLength; pos++ {

		column := []string{}
		for row, _ := range matrix {
			column = append(column, matrix[row][pos])
		}

		if leastCommonBit(column) == "0" {
			matrix = keepMatrixRowsZeros(pos, matrix)
		} else {
			matrix = keepMatrixRowsOnes(pos, matrix)
		}

		// if there's one row left, return
		if len(matrix) == 1 {
			res = strings.Join(matrix[0], "")
			break
		}
	}
	return res
}

func mostCommonBit(s []string) string {
	var (
		countZero int
		countOne  int
	)
	for _, v := range s {
		switch v {
		case "0":
			countZero++
		case "1":
			countOne++
		}
	}

	if countZero < countOne {
		return "1"
	} else if countZero > countOne {
		return "0"
	} else {
		return "1"
	}
}

func leastCommonBit(s []string) string {
	var (
		countZero int
		countOne  int
	)
	for _, v := range s {
		switch v {
		case "0":
			countZero++
		case "1":
			countOne++
		}
	}

	if countZero > countOne {
		return "1"
	} else if countZero < countOne {
		return "0"
	} else {
		return "0"
	}
}
