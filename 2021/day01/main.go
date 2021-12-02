package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var curr, prev, count int

	count-- // skips first measurement

	for scanner.Scan() {
		curr, _ = strconv.Atoi(scanner.Text())

		if curr > prev {
			count++
		}

		prev = curr
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var curr int
	var nums []int

	for scanner.Scan() {
		curr, _ = strconv.Atoi(scanner.Text())
		nums = append(nums, curr)
	}

	lenNums := len(nums) - 1
	var count, a, b int
	count-- // skips first measurement
	for i, _ := range nums {
		if i == 0 || i == lenNums {
			continue
		}

		a = sum(nums[i-1 : i+2])
		if a > b {
			count++
		}
		b = a
	}

	fmt.Println("Part 2:", count)
}

func sum(vals []int) int {
	var s int
	for _, val := range vals {
		s += val
	}
	return s
}
