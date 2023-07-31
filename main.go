package main

import (
	"fmt"
	"raf/aoc2021/day10"
	"raf/aoc2021/day11"
	"raf/aoc2021/day12"
	"raf/aoc2021/day13"
	"raf/aoc2021/day9"
	"time"
)

func main() {

	total := time.Now()
	fmt.Println("Less go!")

	run(day9.Run, 9)
	run(day10.Run, 10)
	run(day11.Run, 11)
	run(day12.Run, 12)
	run(day13.Run, 13)

	fmt.Printf("\nTotal time: %s", time.Since(total))
}

func run(runFunc func() (string, string), day int) {
	start := time.Now()
	ans1, ans2 := runFunc()
	elapsed := time.Since(start)
	fmt.Printf("Day %d \t| Part1: %s\t| Part2: %s\t| %s\n", day, ans1, ans2, elapsed)
}
