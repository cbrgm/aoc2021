package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth         int
	horizontalPos int
}

type Command struct {
	name  string
	value int
}

func (p *Position) moveUp(v int) {
	p.depth = p.depth + v
}

func (p *Position) moveDown(v int) {
	p.depth = p.depth - v
}

func (p *Position) moveForward(v int) {
	p.horizontalPos = p.horizontalPos + v
}

func readInputFile(file *os.File) ([]Command, error) {
	scanner := bufio.NewScanner(file)
	var commands []Command
	for scanner.Scan() {
		if scanner.Text() != "" {
			command := strings.Split(scanner.Text(), " ")
			name := command[0]
			value, err := strconv.Atoi(command[1])
			if err != nil {
				return nil, err
			}

			commands = append(commands, Command{
				name:  name,
				value: value,
			})
		}
	}
	return commands, nil
}

func main() {

	currentPos := Position{
		depth:         0,
		horizontalPos: 0,
	}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	commands, err := readInputFile(f)
	if err != nil {
		panic(err)
	}

	for _, c := range commands {
		switch c.name {
		case "forward":
			currentPos.moveForward(c.value)
		case "up":
			currentPos.moveUp(c.value)
		case "down":
			currentPos.moveDown(c.value)
		}
	}

	fmt.Printf("Result: %v", (currentPos.horizontalPos*currentPos.depth) * -1)

}
