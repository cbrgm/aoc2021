package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth         int
	horizontalPos int
	aim           int
}

type Command struct {
	name  string
	value int
}

func (p *Position) moveUp(v int) {
	p.depth = p.depth - v
}

func (p *Position) moveUpWithAim(v int) {
	p.aim = p.aim - v
}

func (p *Position) moveDown(v int) {
	p.depth = p.depth + v
}

func (p *Position) moveDownWithAim(v int) {
	p.aim = p.aim + v
}

func (p *Position) moveForward(v int) {
	p.horizontalPos = p.horizontalPos + v
}

func (p *Position) moveForwardWithAim(v int) {
	p.horizontalPos = p.horizontalPos + v
	p.depth = p.depth + (p.aim * v)
}

func readInputFile(file *os.File) ([]Command, error) {
	scanner := bufio.NewScanner(file)
	var commands []Command
	for scanner.Scan() {
		c, err := parseCommandFromString(scanner.Text())
		if err != nil {
			return nil, err
		}
		commands = append(commands, c)
	}
	return commands, nil
}

func parseCommandFromString(s string) (Command, error) {
	if s == "" {
		return Command{}, errors.New("empty string")
	}
	command := strings.Split(s, " ")
	name := command[0]
	value, err := strconv.Atoi(command[1])
	if err != nil {
		return Command{}, err
	}
	return Command{
		name:  name,
		value: value,
	}, nil
}

func execWithoutAim(currentPos Position, commands []Command) {
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

	fmt.Printf("Result: %v \n", (currentPos.horizontalPos*currentPos.depth))
}

func execWithAim(currentPos Position, commands []Command) {
	for _, c := range commands {
		switch c.name {
		case "forward":
			currentPos.moveForwardWithAim(c.value)
		case "up":
			currentPos.moveUpWithAim(c.value)
		case "down":
			currentPos.moveDownWithAim(c.value)
		}
	}

	fmt.Printf("Result: %v \n", (currentPos.horizontalPos*currentPos.depth))
}

func main() {

	// task 01

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

	execWithoutAim(currentPos, commands)

	// task 02

	currentPos = Position{
		depth:         0,
		horizontalPos: 0,
		aim:           0,
	}

	execWithAim(currentPos, commands)

}
