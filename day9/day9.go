package day9

import (
	"fmt"
	. "raf/aoc2021/inputFetcher"
	"sort"
)

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(9)
	return part1(input), part2(input)
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
				lowPoints = append(lowPoints, Point{X: x, Y: y})
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
	myvalue := m[p.Y][p.X]
	size := 1
	seen[p.Y][p.X] = true

	for _, n := range GetNeighbors(m, p, false) {
		if seen[n.Y][n.X] {
			continue
		}
		theirValue := m[n.Y][n.X]
		if theirValue < 9 && myvalue < theirValue {
			size += getBasinSize(m, seen, n)
		}
	}
	return size
}

func isLowPoint(m [][]int, x, y int) bool {
	for _, n := range GetNeighbors(m, Point{X: x, Y: y}, false) {
		if m[n.Y][n.X] <= m[y][x] {
			return false
		}
	}

	return true
}
