package day13

import (
	"fmt"
	. "raf/aoc2021/inputFetcher"
	"strconv"
	"strings"
)

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(13)
	return part1(input), part2(input)
}

type Fold struct {
	alongX bool
	value  int
}

func parseInput(lines []string) ([][]byte, []Fold) {

	cutoff := 0
	for i, line := range lines {
		if strings.Contains(line, "fold along") {
			cutoff = i
			break
		}
	}

	gridString := lines[0:cutoff]
	grid := makeGrid(gridString)
	foldStrings := lines[cutoff:]
	folds := make([]Fold, 0, len(foldStrings))

	for _, line := range foldStrings {
		xIndex := strings.Index(line, "x=")
		if xIndex > 0 {
			val, _ := strconv.ParseInt(string(line[xIndex+2:]), 10, 0)
			folds = append(folds, Fold{alongX: true, value: int(val)})
		} else {
			yIndex := strings.Index(line, "y=")
			val, _ := strconv.ParseInt(string(line[yIndex+2:]), 10, 0)
			folds = append(folds, Fold{alongX: false, value: int(val)})
		}
	}

	return grid, folds
}

func makeGrid(lines []string) [][]byte {

	points := make([]Point, 0, len(lines))
	maxY, maxX := 0, 0

	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		if x > maxX {
			maxX = x
		}
		y, _ := strconv.Atoi(split[1])
		if y > maxY {
			maxY = y
		}
		points = append(points, Point{X: x, Y: y})
	}

	result := make([][]byte, maxY+1)
	for y := range result {
		result[y] = make([]byte, maxX+1)
		for x := range result[y] {
			result[y][x] = '.'
		}
	}

	for _, point := range points {
		result[point.Y][point.X] = '#'
	}

	return result
}

func part1(lines []string) string {
	grid, folds := parseInput(lines)
	fold := folds[0]
	if fold.alongX {
		grid = flipVertically(grid, fold.value)
	} else {
		grid = flipHorizontally(grid, fold.value)
	}

	return fmt.Sprintf("%d", countDots(grid))
}

func part2(lines []string) string {
	grid, folds := parseInput(lines)
	for _, fold := range folds {
		if fold.alongX {
			grid = flipVertically(grid, fold.value)
		} else {
			grid = flipHorizontally(grid, fold.value)
		}
	}

	print(grid)
	return fmt.Sprintf("%d", countDots(grid))
}

func countDots(grid [][]byte) int {
	acc := 0
	for _, row := range grid {
		acc += strings.Count(string(row), "#")
	}

	return acc
}

func flipVertically(grid [][]byte, xAxis int) [][]byte {
	firstWidth := xAxis
	secondWidth := len(grid[0]) - 1 - xAxis

	difference := firstWidth - secondWidth
	if difference < 0 {
		difference *= -1
	}

	firstIsBigger := true
	biggerWidth := firstWidth
	if secondWidth > firstWidth {
		biggerWidth = firstWidth
		firstIsBigger = false
	}

	result := make([][]byte, len(grid))
	for y := range result {
		result[y] = make([]byte, biggerWidth)
	}

	//first half
	start := 0
	if !firstIsBigger {
		start = difference
	}

	for y := 0; y < len(grid); y++ {
		for x := start; x < biggerWidth; x++ {
			result[y][x] = grid[y][x-start]
		}
	}

	//second half
	toFlip := grid[:]
	for y := range toFlip {
		toFlip[y] = grid[y][xAxis+1:]
	}

	start = 0
	if firstIsBigger {
		start = difference
	}

	height := len(toFlip)
	width := len(toFlip[0])
	for y := 0; y < height; y++ {
		for x := start; x < biggerWidth; x++ {
			result[y][x] = combineChars(result[y][x], toFlip[y][width-1-x+start])
		}
	}

	return result
}

func flipHorizontally(grid [][]byte, yAxis int) [][]byte {

	firstHeight := yAxis
	secondHeight := len(grid) - 1 - yAxis

	difference := firstHeight - secondHeight
	if difference < 0 {
		difference *= -1
	}

	firstIsBigger := true
	biggerHeight := firstHeight
	if secondHeight > firstHeight {
		biggerHeight = secondHeight
		firstIsBigger = false
	}

	result := make([][]byte, biggerHeight)
	for y := range result {
		result[y] = make([]byte, len(grid[0]))
		for x := range result[y] {
			result[y][x] = '.'
		}
	}

	//first half
	start := 0
	if !firstIsBigger {
		start = difference
	}

	for y := start; y < biggerHeight; y++ {
		for x := 0; x < len(grid[0]); x++ {
			result[y][x] = grid[y-start][x]
		}
	}

	//second half
	toFlip := grid[yAxis+1:]

	start = 0
	if firstIsBigger {
		start = difference
	}

	width := len(toFlip[0])
	height := len(toFlip)
	for y := start; y < biggerHeight; y++ {
		for x := 0; x < width; x++ {
			result[y][x] = combineChars(result[y][x], toFlip[height-1-y+start][x])
		}
	}

	return result
}

func combineChars(first, second byte) byte {
	if first == '#' || second == '#' {
		return '#'
	} else {
		return '.'
	}
}

func print(grid [][]byte) {
	fmt.Println("-------------------------------------")
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------------")
}
