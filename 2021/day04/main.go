package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coords struct {
	x, y int
}

type Board struct {
	arrBoard [5][5]int
	bitBoard [5][5]bool
	lookup   map[int]Coords
}

func main() {
	input, _ := ioutil.ReadFile("input")
	stringInput := strings.Split(string(input), "\n\n")

	var calledNums []int
	for _, item := range strings.Split(stringInput[0], ",") {
		i, _ := strconv.Atoi(item)
		calledNums = append(calledNums, i)
	}

	var boards []Board

	for _, board := range stringInput[1:] {
		var arrBoard [5][5]int
		var bitBoard [5][5]bool
		lookup := make(map[int]Coords)
		for j, row := range strings.Split(board, "\n") {
			currIndex := 0
			for _, item := range strings.Split(row, " ") {
				num, err := strconv.Atoi(item)
				if err == nil {
					arrBoard[j][currIndex] = num
					lookup[num] = Coords{j, currIndex}
					currIndex++
				}
			}
		}

		boards = append(boards, Board{arrBoard, bitBoard, lookup})
	}
	part1and2(calledNums, boards)
}

func part1and2(calledNums []int, boards []Board) {
	first := true
	last := false
	nextNonWinners := boards
	var nonWinners []Board
	for _, currCalledNum := range calledNums {
		nonWinners = nextNonWinners[:]
		nextNonWinners = []Board{}
		if len(nonWinners) == 1 {
			last = true
		}
		// fmt.Println("~~~", currCalledNum, first, last)

		for i := range nonWinners {
			// fmt.Println(i, nonWinners[i].bitBoard)
			bingo := false
			if val, ok := nonWinners[i].lookup[currCalledNum]; ok {
				nonWinners[i].bitBoard[val.x][val.y] = true
				bingo = checkForBingo(&nonWinners[i], val.x, val.y)
				if bingo {
					if first {
						fmt.Println("Part 1:", sumOfUnmarkedNums(&nonWinners[i])*currCalledNum)
						first = false
					}
					if last {
						fmt.Println("Part 2:", sumOfUnmarkedNums(&nonWinners[i])*currCalledNum)
						return
					}
				}
			}
			if !bingo {
				nextNonWinners = append(nextNonWinners, nonWinners[i])
			}
		}
	}
}

func checkForBingo(board *Board, x, y int) bool {
	// hor
	bingo := true
	for i := 0; i < 5; i++ {
		if !board.bitBoard[i][y] {
			bingo = false
			break
		}
	}

	if bingo {
		return true
	}

	bingo = true
	for i := 0; i < 5; i++ {
		if !board.bitBoard[x][i] {
			bingo = false
			break
		}
	}

	return bingo
}

func sumOfUnmarkedNums(board *Board) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !board.bitBoard[x][y] {
				sum += board.arrBoard[x][y]
			}
		}
	}
	return sum
}
