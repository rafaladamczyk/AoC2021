package main

import (
	"fmt"
	"sort"
)

func Day9() {
	//input := GetExampleInput()
	input := GetInput(9)
	fmt.Println("Less go!")
	fmt.Printf("Day 9 Part 1: %s\n", part1(input))
	fmt.Printf("Day 9 Part 2: %s\n", part2(input))
}

func part1(lines []string) string {
	m := Make2DintArray(lines)
	sum := 0
	for y, row := range m {
		for x, val := range row {
			if isLowPoint(m, x, y) {
				sum += val + 1
			}
		}
	}

	return fmt.Sprint(sum)
}

func part2(lines []string) string {
	m := Make2DintArray(lines)
	lowPoints := make([]Point, 0, 100)

	for y, row := range m {
		for x := range row {
			if isLowPoint(m, x, y) {
				lowPoints = append(lowPoints, Point{x: x, y: y})
			}
		}
	}

	basinSizes := make([]int, len(lowPoints))
	seen := make([][]bool, len(m))
	for i := range seen {
		seen[i] = make([]bool, len(m[0]))
	}

	for i, p := range lowPoints {
		basinSizes[i] = getBasinSize(m, seen, p)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	total := basinSizes[0] * basinSizes[1] * basinSizes[2]

	return fmt.Sprintf("%d", total)
}

func getBasinSize(m [][]int, seen [][]bool, p Point) int {
	myvalue := m[p.y][p.x]
	size := 1
	seen[p.y][p.x] = true

	for _, n := range getNeighbors(m, p) {
		if seen[n.y][n.x] {
			continue
		}
		theirValue := m[n.y][n.x]
		if theirValue < 9 && myvalue < theirValue {
			size += getBasinSize(m, seen, n)
		}
	}
	return size
}

func isLowPoint(m [][]int, x, y int) bool {
	for _, n := range getNeighbors(m, Point{x: x, y: y}) {
		if m[n.y][n.x] <= m[y][x] {
			return false
		}
	}

	return true
}

func getNeighbors(m [][]int, p Point) []Point {
	neighbors := make([]Point, 0, 4)
	for _, d := range GetDirs() {
		candidate := Point{x: d.x + p.x, y: d.y + p.y}
		if isValid(m, candidate) {
			neighbors = append(neighbors, candidate)
		}
	}

	return neighbors
}

func isValid(m [][]int, p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(m[0]) && p.y < len(m)
}
