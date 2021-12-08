package day07

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Run(path string) {
	fmt.Println("day07")
	nums := GetInput(path)
	// fmt.Println(nums)
	fmt.Println("part 1:", SolvePart1(nums))
	fmt.Println("part 2:", SolvePart2(nums))
}

func SolvePart1(nums []int) int {
	median := Median(nums)
	var sum int
	for _, n := range nums {
		sum += AbsDiffInt(median, n)
	}
	return sum
}

func SolvePart2(nums []int) int {
	mean := Mean(nums)
	var sum int
	for _, n := range nums {
		sum += Summation(AbsDiffInt(mean, n))
	}
	return sum
}

func GetInput(path string) []int {
	input, _ := ioutil.ReadFile(path)
	stringInput := strings.Split(string(input), ",")
	var nums []int
	for _, item := range stringInput {
		i, _ := strconv.Atoi(item)
		nums = append(nums, i)
	}
	sort.Ints(nums)
	var testNums []int
	for _, i := range nums {
		testNums = append(testNums, int(i))
	}

	return testNums
}
func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Summation(n int) int {
	n++
	return (n * (n - 1)) / 2
}

func Mean(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += int(n)
	}
	return sum / len(nums)
}

// nums needs to be sorted before calling this
func Median(nums []int) int {
	// sort.Ints(nums)
	l := len(nums)
	middle := l / 2
	if l%2 == 1 {
		return nums[middle]
	} else {
		return (nums[middle-1] + nums[middle]) / 2
	}
}
