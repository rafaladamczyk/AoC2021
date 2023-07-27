package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func GetDirs() []Point {
	return []Point{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}
}

func GetInput(day int) []string {
	x := getDataFromUrl(fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day))
	return split(x)
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

func getDataFromUrl(inputUrl string) string {

	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5f3eb94f24f1ecbdc5d2d50e329c47bf4e31b1c3e2456c2a70e475708ad4d105e814fd162373b89ef231f430953e3fc57fbe3b760b5bdff39e",
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

	return string(bytes)
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
