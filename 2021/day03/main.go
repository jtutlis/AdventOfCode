package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// const MAX_BIT = 5
const MAX_BIT = 12

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	var nums []int64
	var curr int64

	for scanner.Scan() {
		curr, _ = strconv.ParseInt(scanner.Text(), 2, 64)

		nums = append(nums, curr)
	}

	part1(nums)
	part2(nums)
}

func part1(nums []int64) {

	var gammaArr [MAX_BIT]int

	for _, item := range nums {
		for i := 0; i < MAX_BIT; i++ {
			if (item>>i)&1 == 1 {
				gammaArr[MAX_BIT-1-i]++
			}
		}
	}
	var gamma int
	half := len(nums) / 2
	for _, item := range gammaArr {
		if item < half {
			gamma = gamma << 1
		} else {
			gamma = (gamma << 1) | 1
		}
	}
	epsilon := ^gamma & ((1 << MAX_BIT) - 1)
	fmt.Println("Part 1:", epsilon*gamma)
}

func part2(nums []int64) {
	oxyList := nums
	for i := int64(0); i < MAX_BIT; i++ {
		oxyList = filter(oxyList, i, mostCommonBit(oxyList, i, true))
	}

	co2List := nums
	for i := int64(0); i < MAX_BIT; i++ {
		co2List = filter(co2List, i, mostCommonBit(co2List, i, false))
		if len(co2List) == 1 {
			break
		}

	}

	fmt.Println("Part 2:", oxyList[0]*co2List[0])
}

func filter(nums []int64, index, requirement int64) []int64 {
	var filtered []int64
	for _, item := range nums {
		if (item>>(MAX_BIT-1-index))&1 == requirement {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func mostCommonBit(nums []int64, index int64, oxygen bool) int64 {
	var res int
	for _, item := range nums {
		if (item>>(MAX_BIT-1-index))&1 == 1 {
			res++
		}
	}

	if oxygen {
		if res < len(nums)-res {
			return 0
		} else {
			return 1
		}
	} else {
		if res == len(nums)-res {
			return 0
		} else if res < len(nums)-res {
			return 1
		} else {
			return 0
		}
	}
}
