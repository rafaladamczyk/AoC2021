package day14

import (
	"fmt"
	. "raf/aoc2021/inputFetcher"
	"strings"
)

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(14)
	return part1(input), part2(input)
}

func part1(lines []string) string {
	return computePolymers(lines, 10)
}

func part2(lines []string) string {
	return computePolymers(lines, 40)
}

func computePolymers(lines []string, steps int) string {
	m := make(map[string]int64)
	initial := lines[0]

	for i := 0; i < len(initial)-1; i++ {
		s := initial[i : i+2]
		m[s] += 1
	}

	for n := 0; n < steps; n++ {
		changes := make(map[string]int64)
		for l := 1; l < len(lines); l++ {
			line := lines[l]
			split := strings.Split(line, " -> ")
			existing := split[0]
			inbetween := split[1]
			howManyExist := m[existing]
			if howManyExist > 0 {
				changes[existing] -= howManyExist
				first := string(existing[0]) + inbetween
				second := inbetween + string(existing[1])
				changes[first] += howManyExist
				changes[second] += howManyExist
			}
		}

		for key, value := range changes {
			m[key] += value
		}
	}

	counts := make(map[string]int64)
	for key, value := range m {
		counts[string(key[1])] += value
	}

	max := int64(0)
	min := int64(^uint64(0) >> 1)
	for _, value := range counts {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return fmt.Sprintf("%d", max-min)
}
