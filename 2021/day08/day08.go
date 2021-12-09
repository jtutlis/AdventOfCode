package day08

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Line struct {
	in, out []string
}

func Run(path string) {
	fmt.Println("day 8")
	lines := GetInput(path)
	fmt.Println("part 1:", SolvePart1(lines))
	fmt.Println("part 2:", SolvePart2(lines))
}

func SolvePart1(lines []Line) int {
	var count int
	for _, line := range lines {
		for _, word := range line.out {
			switch len(word) {
			case 2, 3, 4, 7:
				count++
			}

		}
	}
	return count
}

var runeMap = map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6}

func SolvePart2(lines []Line) int {

	// we need to solve the mapping
	// generate possible mappings based on 1, 4, 7, 8
	// then reduce, like a constraint satisfaction problem?

	/*
		 aaaa  	     0000
		b    c      1    2
		b    c      1    2
		 dddd   ->   3333
		e    f      4    5
		e    f      4    5
		 gggg        6666
	*/

	/*
		how many times each apppears in all 10 digits
		'0->8 solved by unqiue value of 1 & 7
		'1->6 unique
		'2->8 solved after '0
		'3->7 solved from 4
		'4->4 unique
		'5->9 unique
		'6->7 solved after '3
	*/
	var sum int

	for _, line := range lines {
		var possibleMappings [7]int
		var possibleZero [7]int
		var possibleFour [7]int
		var mapping [7]int
		var zero int
		var currLetter int
		for _, word := range line.in {
			l := len(word)
			for _, r := range word {
				currLetter = runeMap[r]
				possibleMappings[currLetter]++
				switch l {
				case 2, 3: // "1"
					possibleZero[currLetter]++
				case 4: // "4"
					possibleFour[currLetter]++
				}
			}
		}

		for index, value := range possibleZero {
			if value == 1 {
				mapping[index] = 0
				zero = index
			}
		}
		for index, value := range possibleMappings {
			switch value {
			case 9:
				mapping[index] = 5
			case 4:
				mapping[index] = 4
			case 6:
				mapping[index] = 1
			case 8:
				if index != zero {
					mapping[index] = 2
				}
			case 7:
				if possibleFour[index] != 0 {
					mapping[index] = 3
				} else {
					mapping[index] = 6
				}
			}
		}
		var num int
		for _, word := range line.out {
			l := len(word)

			switch l {
			case 2: // "1"
				num = num*10 + 1
			case 3: // "7"
				num = num*10 + 7
			case 4: // "4"
				num = num*10 + 4
			case 7: // "8"
				num = num*10 + 8
			default:
				num = num*10 + getNum(word, mapping)
			}
		}
		sum += num
	}
	return sum
}

func getNum(word string, mapping [7]int) int {

	if len(word) == 5 { // 2,3,5
		for _, r := range word {
			index := runeMap[r]
			if mapping[index] == 4 {
				return 2
			}
			if mapping[index] == 1 {
				return 5
			}
		}
		return 3

	} else { // 0,6,9
		var sorn, n bool
		for _, r := range word {
			index := runeMap[r]
			if mapping[index] == 3 { // 6 or 9
				sorn = true
			} else if mapping[index] == 2 {
				n = true
				if sorn {
					return 9
				}
			}
		}
		if sorn {
			if n {
				return 9
			} else {
				return 6
			}
		} else {
			return 0
		}
	}
}

func GetInput(path string) []Line {
	var lines []Line
	input, _ := ioutil.ReadFile(path)
	for _, l := range strings.Split(string(input), "\n") {
		stringInput := strings.Split(l, " | ")
		line := Line{strings.Split(string(stringInput[0]), " "), strings.Split(string(stringInput[1]), " ")}
		lines = append(lines, line)
	}
	return lines
}
