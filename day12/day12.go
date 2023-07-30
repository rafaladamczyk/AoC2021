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

	normalChan := make(chan int)
	go func() { normalChan <- len(findAllPaths(start, end, make([]*Node, 0, 1000))) }()

	smallChans := make([]chan int, len(smalls))

	for i, small := range smalls {
		smallChans[i] = make(chan int)
		go func(smallNode *Node, sc chan int) {
			acc := 0
			for _, path := range findAllPaths2(start, end, make([]*Node, 0, 1000), smallNode.name) {
				if count(path, smallNode) == 2 {
					acc++
				}
			}

			sc <- acc
		}(small, smallChans[i])
	}

	result := <-normalChan
	for _, smallChan := range smallChans {
		result += <-smallChan
	}

	return fmt.Sprintf("%d", result)
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
		if dest.isBig || (dest.name == small && count(visited, dest) <= 1) || !contains(visited, dest) {
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
