package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Run(path string) {
	fmt.Println("day 11")
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

type Point struct {
	y, x int
}

var Dirs = []Point{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

func SolvePart1(matrix [10][10]int8) int {
	var count int
	flashNum := int8(9) // passing this unsued var make its 9000ns faster lmao
	toFlash := make([]Point, 0, 100)
	var curr Point
	for step := 0; step < 100; step++ {

		flashedCount := 0
		for j := range matrix {
			for i := range matrix[j] {
				if matrix[j][i] == -1 {
					matrix[j][i] = 1
				} else {
					matrix[j][i]++
				}

				if matrix[j][i] > 9 {
					matrix[j][i] = -1
					flashedCount++
					toFlash = append(toFlash, Point{j, i})
				}
			}
		}

		for len(toFlash) > 0 {
			curr, toFlash = toFlash[len(toFlash)-1], toFlash[:len(toFlash)-1]
			flashedCount += flash(curr.y, curr.x, &matrix, flashNum)

		}
		count += flashedCount
	}

	return count
}

func flash(y, x int, matrix *[10][10]int8, flashNum int8) int {
	var count int

	for _, dir := range Dirs {
		newY, newX := y+dir.y, x+dir.x
		if !(newY == -1 || newY == 10 || newX == -1 || newX == 10) && matrix[newY][newX] != -1 {
			matrix[newY][newX]++
			if matrix[newY][newX] > 9 {
				matrix[newY][newX] = -1
				count += flash(newY, newX, matrix, flashNum) + 1
			}

		}
	}
	return count
}

func SolvePart2(matrix [10][10]int8) int {
	flashNum := int8(9) // passing this unsued var make its 9000ns faster lmao
	toFlash := make([]Point, 0, 100)
	var curr Point
	for step := 0; step < 1000; step++ {

		flashedCount := 0
		for j := range matrix {
			for i := range matrix[j] {
				if matrix[j][i] == -1 {
					matrix[j][i] = 1
				} else {
					matrix[j][i]++
				}

				if matrix[j][i] > 9 {
					matrix[j][i] = -1
					flashedCount++
					toFlash = append(toFlash, Point{j, i})
				}
			}
		}

		for len(toFlash) > 0 {
			curr, toFlash = toFlash[len(toFlash)-1], toFlash[:len(toFlash)-1]
			flashedCount += flash(curr.y, curr.x, &matrix, flashNum)
		}

		if flashedCount == 100 {
			return step + 1
		}
	}

	return -1
}

func GetInput(path string) [10][10]int8 {
	input, _ := ioutil.ReadFile(path)
	var matrix [10][10]int8
	for j, line := range strings.Split(string(input), "\n") {
		for i, item := range line {
			matrix[j][i] = int8(item - '0')
		}
	}

	return matrix
}
