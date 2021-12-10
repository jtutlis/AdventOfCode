package day09

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Run(path string) {
	fmt.Println("day 9")
	matrix := GetInput(path)
	fmt.Println("part 1:", SolvePart1(matrix))
	fmt.Println("part 2:", SolvePart2(matrix))
}

type Pair struct {
	j, i int8
}

var Dirs = []Pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func SolvePart1(matrix [][]int8) int {
	var risk int

	xBound := int8(len(matrix[0]))
	yBound := int8(len(matrix))

	for j := range matrix {
		for i := range matrix[j] {
			currValue := matrix[j][i]
			var greaterOrEqual bool
			for _, dir := range Dirs {
				y, x := dir.j+int8(j), dir.i+int8(i)
				if x >= 0 && x != xBound && y >= 0 && y != yBound {
					checkedValue := matrix[y][x]
					if checkedValue <= currValue {
						greaterOrEqual = true
						break
					}
				}

			}
			if !greaterOrEqual {
				risk += int(currValue) + 1
			}
		}
	}

	return risk
}

func SolvePart2(matrix [][]int8) int {
	bound := int8(len(matrix))
	var first, second, third, skipY, skipX int // we can skip ever few rows and cols since we are looking for the largest basins
	var alreadySearched [100][100]bool
	for j := range matrix {
		if skipY == 6 {
			skipY = 0
		} else {
			skipY++
			continue
		}
		for i := range matrix[j] {
			if skipX == 6 {
				skipX = 0
			} else {
				skipX++
				continue
			}
			if alreadySearched[j][i] || matrix[j][i] == 9 {
				continue
			}

			alreadySearched[j][i] = true
			sum := search(matrix, int8(j), int8(i), bound, &alreadySearched)
			if sum > first {
				third = second
				second = first
				first = sum
			} else if sum > second {
				third = sum
				second = sum
			} else if sum > third {
				third = sum
			}
		}
	}

	return (first + 1) * (second + 1) * (third + 1)
}

func search(matrix [][]int8, y, x, bound int8, alreadySearched *[100][100]bool) int {
	var space int
	for _, dir := range Dirs {
		var newY, newX int8
		if dir.i == 0 {
			newY, newX = dir.j+y, x
			if newY == -1 || newY == bound {
				continue
			}
		} else {
			newY, newX = y, dir.i+x
			if newX == -1 || newX == bound {
				continue
			}
		}

		if matrix[newY][newX] != 9 && !alreadySearched[newY][newX] {
			alreadySearched[newY][newX] = true
			space++
			space += search(matrix, newY, newX, bound, alreadySearched)
		}
	}

	return space
}

func GetInput(path string) [][]int8 {
	var matrix [][]int8
	input, _ := ioutil.ReadFile(path)
	for i, row := range strings.Split(string(input), "\n") {
		matrix = append(matrix, []int8{})
		for _, r := range row {
			matrix[i] = append(matrix[i], int8(r-'0'))
		}
	}
	return matrix
}
