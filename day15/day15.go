package day15

import (
	"container/heap"
	"fmt"
	. "raf/aoc2021/dataStructures"
	. "raf/aoc2021/inputFetcher"
	"strconv"
)

type Node struct {
	p           Point
	dist, value int
	via         *Node
}

func Run() (string, string) {
	//input := GetExampleInput()
	input := GetInput(15)
	return part1(input), part2(input)
}

func part1(lines []string) string {
	grid := makeGrid(lines)
	dest := pathfindPriorityQ(grid)
	return fmt.Sprintf("%d", dest.dist)
}

func part2(lines []string) string {
	grid := makeBigGrid(lines)
	dest := pathfindPriorityQ(grid)
	return fmt.Sprintf("%d", dest.dist)
}

func makeGrid(lines []string) [][]*Node {
	d := len(lines)

	grid := make([][]*Node, d)
	for y, l := range lines {
		grid[y] = make([]*Node, d)
		for x, c := range l {
			val, _ := strconv.Atoi(string(c))
			grid[y][x] = &Node{dist: int(^uint(0) >> 1), value: val, via: nil, p: Point{X: x, Y: y}}
		}
	}

	grid[0][0].dist = 0

	return grid
}

func makeBigGrid(lines []string) [][]*Node {
	d := len(lines)

	grid := make([][]*Node, 5*d)
	for y, l := range lines {
		for i := 0; i < 5; i++ {
			grid[y+i*d] = make([]*Node, 5*d)
		}
		for x, c := range l {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					val, _ := strconv.Atoi(string(c))
					adjusted := val + i + j
					for adjusted > 9 {
						adjusted -= 9
					}

					grid[y+i*d][x+j*d] = &Node{dist: int(^uint(0) >> 1), value: adjusted, via: nil, p: Point{X: x + j*d, Y: y + i*d}}
				}
			}

		}
	}

	grid[0][0].dist = 0

	//print(grid)
	return grid
}

func print(grid [][]*Node) {
	for _, row := range grid {
		for _, node := range row {
			fmt.Printf("%d", node.value)
		}
		fmt.Println()
	}
}

func pathfind(grid [][]*Node) *Node {

	d := len(grid) - 1
	q := make([]*Node, 0, len(grid)*len(grid))

	for _, row := range grid {
		q = append(q, row...)
	}

	for len(q) != 0 {
		currIndex := minDistance(q)
		curr := q[currIndex]
		if curr.p.X == d && curr.p.Y == d {
			return curr
		}

		q = remove(q, currIndex)
		neighbors := getNeighbors(curr, grid)
		for _, neigh := range neighbors {
			alt := curr.dist + neigh.value
			if alt < neigh.dist {
				neigh.dist = alt
				neigh.via = curr
			}
		}
	}

	return nil
}

func remove(q []*Node, i int) []*Node {
	q[i] = q[len(q)-1]
	return q[:len(q)-1]
}

func pathfindPriorityQ(grid [][]*Node) *Node {

	d := len(grid) - 1
	pq := make(PriorityQueue, 1)

	pq[0] = &QueueItem{Value: grid[0][0], Priority: 0, Index: 0}
	heap.Init(&pq)

	for len(pq) != 0 {
		item := heap.Pop(&pq).(*QueueItem)
		curr := item.Value.(*Node)

		if item.Priority <= int(curr.dist) {
			neighbors := getNeighbors(curr, grid)
			for _, neigh := range neighbors {
				alt := curr.dist + neigh.value
				if alt < neigh.dist {
					neigh.dist = alt
					neigh.via = curr
					heap.Push(&pq, &QueueItem{Value: neigh, Priority: int(alt)})
				}
			}
		}
	}

	return grid[d][d]
}

func getNeighbors(node *Node, grid [][]*Node) []*Node {

	d := len(grid)
	result := make([]*Node, 0, 4)
	for _, p := range GetDirs(false) {
		newPoint := Point{node.p.X + p.X, node.p.Y + p.Y}
		if newPoint.X < d && newPoint.X >= 0 && newPoint.Y < d && newPoint.Y >= 0 {
			candidate := grid[newPoint.Y][newPoint.X]
			result = append(result, candidate)
		}
	}

	return result
}

func minDistance(q []*Node) int {
	min := int(^uint(0) >> 1)
	minIndex := -1
	for i, n := range q {
		if n.dist < min {
			min = n.dist
			minIndex = i
		}
	}

	return minIndex
}
