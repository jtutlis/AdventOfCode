package day12

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Run(path string) {
	fmt.Println("day 12")
	start := time.Now()
	caveMap := GetInput(path)
	fmt.Println("part 1:", SolvePart1(caveMap))
	duration := time.Since(start)
	fmt.Println("duration", duration)
	start = time.Now()
	fmt.Println("part 2:", SolvePart2(caveMap))
	duration = time.Since(start)
	fmt.Println("duration", duration)
}

type Cave struct {
	name  string
	small bool // true if large cave, false is small
}

func SolvePart1(caveMap map[string][]Cave) int {
	var count int

	visited := make(map[string]struct{})
	visited["start"] = struct{}{}
	count += visitCavePart1(Cave{"start", false}, visited, caveMap)

	return count
}

func visitCavePart1(cave Cave, visited map[string]struct{}, caveMap map[string][]Cave) int {
	if cave.name == "end" {
		return 1
	}
	var count int
	for _, newCave := range caveMap[cave.name] {
		if _, ok := visited[newCave.name]; !ok {
			newVisited := copyVisitedMap(visited)
			if newCave.small {
				newVisited[newCave.name] = struct{}{}
			}
			count += visitCavePart1(newCave, newVisited, caveMap)
		}

	}
	return count
}

func copyVisitedMap(originalMap map[string]struct{}) map[string]struct{} {
	newMap := make(map[string]struct{})
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func SolvePart2(caveMap map[string][]Cave) int {
	var count int

	visited := make(map[string]struct{})
	visited["start"] = struct{}{}
	count += visitCavePart2(Cave{"start", false}, visited, caveMap, false)

	return count
}

func visitCavePart2(cave Cave, visited map[string]struct{}, caveMap map[string][]Cave, twice bool) int {
	if cave.name == "end" {
		return 1
	}
	var count int
	for _, newCave := range caveMap[cave.name] {
		if _, ok := visited[newCave.name]; !ok {
			newVisited := copyVisitedMap(visited)
			if newCave.small {
				newVisited[newCave.name] = struct{}{}
			}
			count += visitCavePart2(newCave, newVisited, caveMap, twice)
		} else if !twice && newCave.small && newCave.name != "start" {
			newVisited := copyVisitedMap(visited)
			count += visitCavePart2(newCave, newVisited, caveMap, true)
		}

	}
	return count
}

func GetInput(path string) map[string][]Cave {
	input, _ := ioutil.ReadFile(path)
	caveMap := make(map[string][]Cave)
	for _, line := range strings.Split(string(input), "\n") {
		caves := strings.Split(line, "-")
		a := Cave{caves[0], strings.ToLower(caves[0]) == caves[0]}
		b := Cave{caves[1], strings.ToLower(caves[1]) == caves[1]}

		caveMap[a.name] = append(caveMap[a.name], b)
		caveMap[b.name] = append(caveMap[b.name], a)
	}
	// fmt.Println(caveMap)
	return caveMap
}
