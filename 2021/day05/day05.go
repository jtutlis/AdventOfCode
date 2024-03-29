package day05

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Run(path string) {
	fmt.Println("~~~day 05~~~")
	lines := GetLines(path)
	SolveFast(lines, false, true)
	SolveFast(lines, true, true)
}

type Line struct {
	ax, ay, bx, by int
}

func SolveFast(lines []Line, diagonals bool, print bool) int {
	var hitPoints [990][991]uint8
	var count int
	var dx, dy int
	var x1, x2, y1, y2 int
	r := true

	for _, line := range lines {
		x1, y1, x2, y2 = line.ax, line.ay, line.bx, line.by
		dx, dy = GetDelta(x1, x2), GetDelta(y1, y2)

		if dx != 0 && dy != 0 && !diagonals {
			continue
		}

		r = true
		for r {
			if x1 == x2 && y1 == y2 {
				r = false
			}

			if hitPoints[x1][y1] == 0 {
				hitPoints[x1][y1] = 1
			} else if hitPoints[x1][y1] == 1 {
				count++
				hitPoints[x1][y1] = 2
			}
			x1, y1 = x1+dx, y1+dy
		}
	}
	if print {
		if diagonals {
			fmt.Println("Part 2:", count)
		} else {
			fmt.Println("Part 1:", count)
		}
	}

	return count

}

func GetDelta(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return -1
	}
}

func GetLines(path string) []Line {
	input, _ := ioutil.ReadFile(path)
	stringInput := strings.Split(string(input), "\n")
	var lines []Line
	for _, item := range stringInput {

		z := strings.Split(string(item), " -> ")
		xp1 := strings.Split(string(z[0]), ",")
		xp2 := strings.Split(string(z[1]), ",")

		ax, _ := strconv.Atoi(xp1[0])
		ay, _ := strconv.Atoi(xp1[1])
		bx, _ := strconv.Atoi(xp2[0])
		by, _ := strconv.Atoi(xp2[1])

		line := Line{ax, ay, bx, by}
		lines = append(lines, line)
	}
	return lines
}
