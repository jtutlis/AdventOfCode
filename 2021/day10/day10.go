package day10

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func Run(path string) {
	fmt.Println("day 10")
	start := time.Now()
	lines := GetInput(path)
	fmt.Println("part 1:", SolvePart1(lines))
	duration := time.Since(start)
	fmt.Println("duration", duration)
	start = time.Now()
	fmt.Println("part 2:", SolvePart2(lines))
	duration = time.Since(start)
	fmt.Println("duration", duration)
}

var (
	openersMap = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	errorScore = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	autoScore  = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
)

func SolvePart1(lines []string) int {
	var score int
	for _, line := range lines {
		var stack []rune
		for _, r := range line {
			if match, ok := openersMap[r]; ok {
				stack = append(stack, match)
			} else {
				if r != stack[len(stack)-1] {
					score += errorScore[r]
					break
				} else {
					if len(stack) > 0 {
						stack = stack[:len(stack)-1]
					}
				}
			}
		}
	}

	return score
}

//stolen from https://git.yetaga.in/alazyreader/AdventOfCode2021/src/branch/master/10/main.go
func pop(stack []rune) (rune, []rune) {
	pop := stack[len(stack)-1]
	return pop, stack[:len(stack)-1]
}

func SolvePart2(lines []string) int {
	var scores []int
	for _, line := range lines {
		stack := make([]rune, 0, 100)
		// var stack []rune
		var invalid bool
		for _, r := range line {
			var p rune
			// switch idea stolen from https://git.yetaga.in/alazyreader/AdventOfCode2021/src/branch/master/10/main.go
			switch r {
			case '[', '{', '<', '(':
				stack = append(stack, r)
			case '>':
				p, stack = pop(stack)
				if p != '<' {
					invalid = true
				}
			case ')':
				p, stack = pop(stack)
				if p != '(' {
					invalid = true
				}
			case ']':
				p, stack = pop(stack)
				if p != '[' {
					invalid = true
				}
			case '}':
				p, stack = pop(stack)
				if p != '{' {
					invalid = true
				}
			}
			// if match, ok := openersMap[r]; ok {
			// 	stack = append(stack, match)
			// } else {
			// 	if r != stack[len(stack)-1] {
			// 		invalid = true
			// 		break
			// 	} else {
			// 		if len(stack) > 0 {
			// 			stack = stack[:len(stack)-1]
			// 		}
			// 	}
			// }
		}
		if !invalid {
			var score int
			for i := len(stack) - 1; i >= 0; i-- {
				// score = score*5 + autoScore[stack[i]]
				score *= 5
				switch stack[i] {
				case '(':
					score += 1
				case '[':
					score += 2
				case '{':
					score += 3
				case '<':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func GetInput(path string) []string {
	input, _ := ioutil.ReadFile(path)
	return strings.Split(string(input), "\n")
}
