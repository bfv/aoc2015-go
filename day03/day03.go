package main

import (
	"fmt"
	"io/ioutil"
)

var grid map[string]int

func main() {
	a, b := solve("input.txt")
	fmt.Printf("day 02, a: %v, b: %v", a, b)
}

func solve(filename string) (int, int) {
	var a, b int
	var x, y int

	grid = make(map[string]int)

	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)

	bringPresent(x, y)
	for _, r := range s {
		switch r {
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		case '<':
			x--
		}
		bringPresent(x, y)
	}

	// get answer a
	for _, v := range grid {
		if v > 0 {
			a++
		}
	}

	return a, b
}

func bringPresent(x int, y int) {
	co := fmt.Sprintf("%v,%v", x, y)
	grid[co]++
}
