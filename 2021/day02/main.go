package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	string
	int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []Pair

	for scanner.Scan() {
		curr := strings.Split(scanner.Text(), " ")
		i, _ := strconv.Atoi(curr[1])
		input = append(input, Pair{curr[0], i})
	}

	part1(input)
	part2(input)
}

func part1(input []Pair) {
	var v, h int

	for _, item := range input {
		switch item.string {
		case "up":
			v -= item.int
		case "down":
			v += item.int
		case "forward":
			h += item.int
		}
	}

	fmt.Println("Part 1:", v*h)
}

func part2(input []Pair) {
	var v, h, aim int

	for _, item := range input {
		switch item.string {
		case "up":
			aim -= item.int
		case "down":
			aim += item.int
		case "forward":
			h += item.int
			v += aim * item.int
		}
	}

	fmt.Println("Part 2:", v*h)
}
