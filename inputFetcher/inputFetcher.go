package inputFetcher

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func GetDirs(includeDiagonals bool) []Point {
	if includeDiagonals {

		return []Point{
			{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1},
			{X: 1, Y: 1}, {X: 1, Y: -1}, {X: -1, Y: 1}, {X: -1, Y: -1}}
	} else {
		return []Point{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}
	}
}

func GetInput(day int) []string {
	fName := fmt.Sprintf("inputs/day%d.txt", day)
	data, e := os.ReadFile(fName)
	if e != nil && os.IsNotExist(e) {
		data = getDataFromUrl(fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day))
		err := os.WriteFile(fName, data, 0)
		check(err)
	}

	return split(string(data))
}

func GetExampleInput() []string {
	bytes, err := os.ReadFile("example.txt")
	check(err)
	contents := string(bytes)
	return split(contents)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDataFromUrl(inputUrl string) []byte {
	session, found := os.LookupEnv("aoc_session")
	if !found {
		panic(errors.New("did not find session in env variables"))
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", inputUrl, nil)
	req.AddCookie(cookie)
	r, err := client.Do(req)
	check(err)
	defer r.Body.Close()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Response status:", r.Status)
		log.Fatalln(err)
	}

	return bytes
}

func split(input string) []string {
	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return filter(lines, func(s string) bool { return len(s) > 0 })
}

func filter(ss []string, test func(string) bool) (ret []string) {
	b := ss[:0] // in place filtering, take another slice that references the same array

	for _, x := range ss {
		if test(x) {
			b = append(b, x)
		}
	}

	return b
}

func Make2DintArray(lines []string) [][]int {
	width := len(lines[0])
	height := len(lines)
	var nums = make([][]int, height)
	for y, l := range lines {
		nums[y] = make([]int, width)
		for x, c := range l {
			nums[y][x], _ = strconv.Atoi(string(c))
		}
	}

	return nums
}

func Make2DCharArray(lines []string) [][]uint8 {
	width := len(lines[0])
	height := len(lines)
	var chars = make([][]uint8, width*height)
	for y, l := range lines {
		chars[y] = make([]uint8, width)
		for x, c := range l {
			chars[y][x] = uint8(c)
		}
	}

	return chars
}

func GetNeighbors(m [][]int, p Point, includeDiagonals bool) []Point {
	neighbors := make([]Point, 0, 4)
	for _, d := range GetDirs(includeDiagonals) {
		candidate := Point{X: d.X + p.X, Y: d.Y + p.Y}
		if isValid(m, candidate) {
			neighbors = append(neighbors, candidate)
		}
	}

	return neighbors
}

func isValid(m [][]int, p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < len(m[0]) && p.Y < len(m)
}

func Sum(a [][]int) int {
	acc := 0
	for _, row := range a {
		for _, v := range row {
			acc += v
		}
	}
	return acc
}
