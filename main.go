package main

import (
	"fmt"
	"raf/aoc2021/day10"
	"raf/aoc2021/day11"
	"raf/aoc2021/day9"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("Less go!")

	ans1, ans2 := day9.Run()
	fmt.Printf("Day 9 Part 1: %s\n", ans1)
	fmt.Printf("Day 9 Part 2: %s\n", ans2)

	ans1, ans2 = day10.Run()
	fmt.Printf("Day 10 Part 1: %s\n", ans1)
	fmt.Printf("Day 10 Part 2: %s\n", ans2)

	ans1, ans2 = day11.Run()
	fmt.Printf("Day 11 Part 1: %s\n", ans1)
	fmt.Printf("Day 11 Part 2: %s\n", ans2)

	elapsed := time.Since(start)
	fmt.Printf("\nFinished in %s", elapsed)

}
