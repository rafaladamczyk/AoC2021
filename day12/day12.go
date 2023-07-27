package day12

import (
	"fmt"
	. "raf/aoc2021/inputFetcher"
	"strings"
)

type Node struct {
	destinations []*Node
	name         string
	isBig        bool
}

func (src *Node) AddEdge(dest *Node) {
	src.destinations = append(src.destinations, dest)
}

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(12)
	return part1(input), part2(input)
}

func part1(lines []string) string {
	m := buildMap(lines)
	start := m["start"]
	end := m["end"]

	paths := findAllPaths(start, end, make([]*Node, 0, 1000))
	return fmt.Sprintf("%d", len(paths))
}

func part2(lines []string) string {
	m := buildMap(lines)
	start := m["start"]
	end := m["end"]

	smalls := make([]*Node, 0)
	for _, n := range m {
		if !n.isBig && n.name != "start" {
			smalls = append(smalls, n)
		}
	}

	normalPaths := findAllPaths(start, end, make([]*Node, 0, 1000))

	acc := 0
	for _, small := range smalls {
		for _, path := range findAllPaths2(start, end, make([]*Node, 0, 1000), small.name) {
			if count(path, small) == 2 {
				acc++
			}
		}
	}

	return fmt.Sprintf("%d", len(normalPaths)+acc)
}
func findAllPaths(start, end *Node, visited []*Node) [][]*Node {
	if start == end {
		return [][]*Node{{end}}
	}

	if !start.isBig {
		visited = append(visited, start)
	}

	paths := make([][]*Node, 0)

	for _, dest := range start.destinations {
		if dest.isBig || !contains(visited, dest) {
			v := make([]*Node, len(visited))
			copy(v, visited)
			results := findAllPaths(dest, end, v)
			for _, path := range results {
				newPath := make([]*Node, 0, len(path)+1)
				newPath = append(newPath, start)
				newPath = append(newPath, path...)
				paths = append(paths, newPath)
			}
		}
	}

	return paths
}

func findAllPaths2(start, end *Node, visited []*Node, small string) [][]*Node {
	if start == end {
		return [][]*Node{{end}}
	}

	if !start.isBig {
		visited = append(visited, start)
	}

	paths := make([][]*Node, 0)

	for _, dest := range start.destinations {
		visitedCount := count(visited, dest)
		if dest.isBig || visitedCount == 0 || (dest.name == small && visitedCount == 1) {
			v := make([]*Node, len(visited))
			copy(v, visited)
			for _, path := range findAllPaths2(dest, end, v, small) {
				newPath := make([]*Node, 0, len(path)+1)
				newPath = append(newPath, start)
				newPath = append(newPath, path...)
				paths = append(paths, newPath)
			}
		}
	}

	return paths
}

func contains(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func count(nodes []*Node, node *Node) int {
	acc := 0
	for _, n := range nodes {
		if n == node {
			acc++
		}
	}
	return acc
}

func buildMap(lines []string) map[string]*Node {
	m := make(map[string]*Node)

	for _, l := range lines {
		s := strings.Split(l, "-")
		first := getOrCreateNode(m, s[0])
		second := getOrCreateNode(m, s[1])
		first.AddEdge(second)
		second.AddEdge(first)
	}

	return m
}

func getOrCreateNode(m map[string]*Node, name string) *Node {
	n := m[name]
	if n == nil {
		newNode := &Node{name: name, destinations: make([]*Node, 0), isBig: isBigCavern(name)}
		m[name] = newNode
		return newNode
	} else {
		return n
	}
}

func isBigCavern(s string) bool {
	return strings.ToUpper(s) == s
}

func printPath(p []*Node) {
	for _, n := range p {
		fmt.Printf("%s -> ", n.name)
	}
	fmt.Println()
}
