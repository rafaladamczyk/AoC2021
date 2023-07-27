package day11

import (
	"fmt"
	. "raf/aoc2021/inputFetcher"
)

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(11)
	return part1(input), part2(input)
}

func part1(lines []string) string {
	levels := Make2DintArray(lines)
	allDirs := GetDirs(true)

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += step(levels, allDirs)
	}

	return fmt.Sprintf("%d", flashes)
}

func part2(lines []string) string {
	levels := Make2DintArray(lines)
	allDirs := GetDirs(true)

	for i := 0; i < 9999999; i++ {
		step(levels, allDirs)
		// fmt.Printf("Step %d", i+1)
		// print(levels)
		// fmt.Println()
		// fmt.Println()
		if Sum(levels) == 0 {
			return fmt.Sprintf("%d", i+1)
		}
	}

	return ""
}

func print(levels [][]int) {
	for y, row := range levels {
		fmt.Println()
		for x := range row {
			fmt.Printf(" %d ", levels[y][x])
		}
	}
}

func step(levels [][]int, dirs []Point) int {
	flashes := 0
	for y, row := range levels {
		for x := range row {
			levels[y][x] += 1
			if levels[y][x] > 9 {
				flashes += flash(levels, dirs, x, y)
			}
		}
	}

	for y, row := range levels {
		for x := range row {
			if levels[y][x] >= 100 {
				levels[y][x] = 0
			}
		}
	}

	return flashes
}

func flash(levels [][]int, dirs []Point, x, y int) int {
	if levels[y][x] >= 100 {
		return 0 // we already flashed
	}

	sum := 1
	levels[y][x] = 100 //flash
	for _, n := range GetNeighbors(levels, Point{X: x, Y: y}, true) {
		levels[n.Y][n.X] += 1
		if levels[n.Y][n.X] > 9 {
			sum += flash(levels, dirs, n.X, n.Y)
		}
	}

	return sum
}
