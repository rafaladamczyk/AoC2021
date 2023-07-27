package day10

import (
	"fmt"
	. "raf/aoc2021/dataStructures"
	. "raf/aoc2021/inputFetcher"
	"sort"
)

type tokenPair struct {
	first  byte
	second byte
	score  int
}

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(10)
	return part1(input), part2(input)
}

func part1(lines []string) string {
	s := make(Stack, 0)
	tokenPairs := []tokenPair{
		{first: '(', second: ')', score: 3},
		{first: '[', second: ']', score: 57},
		{first: '{', second: '}', score: 1197},
		{first: '<', second: '>', score: 25137}}
	sum := 0

	for _, row := range lines {
		for _, c := range row {
			tp := getPairByEndChar(tokenPairs, byte(c))
			if tp != nil {
				if len(s) > 0 {
					var e byte
					s, e = s.Pop()
					if e != tp.first {
						sum += tp.score
					}
				}
			} else {
				s = s.Push(byte(c))
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func part2(lines []string) string {
	tokenPairs := []tokenPair{
		{first: '(', second: ')', score: 1},
		{first: '[', second: ']', score: 2},
		{first: '{', second: '}', score: 3},
		{first: '<', second: '>', score: 4}}
	scores := make([]int, 0)

	for _, row := range lines {
		score := 0
		s := make(Stack, 0)
		for _, c := range row {
			tp := getPairByEndChar(tokenPairs, byte(c))
			if tp != nil {
				if len(s) > 0 {
					var e byte
					s, e = s.Pop()
					if e != tp.first {
						s = s[:0] // corrupted row, empty the stack
						break
					}
				}
			} else {
				s = s.Push(byte(c))
			}
		}

		if len(s) > 0 {
			for i := len(s); i > 0; i-- {
				var c byte
				s, c = s.Pop()
				tp := getPairByStartChar(tokenPairs, c)
				score = score*5 + int(tp.score)
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	middle := int(len(scores) / 2)

	return fmt.Sprintf("%d", scores[middle])
}

func getPairByEndChar(s []tokenPair, x byte) *tokenPair {
	for _, b := range s {
		if b.second == x {
			return &b
		}
	}

	return nil
}

func getPairByStartChar(s []tokenPair, x byte) *tokenPair {
	for _, b := range s {
		if b.first == x {
			return &b
		}
	}

	return nil
}
