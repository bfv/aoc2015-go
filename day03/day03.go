package main

import (
	"fmt"
	"io/ioutil"
)

var gridA map[string]int
var gridB map[string]int

func main() {
	a := solveA("input.txt")
	b := solveB("input.txt")
	fmt.Printf("day 02, a: %v, b: %v", a, b)
}

func solveA(filename string) int {
	var a int
	var x, y int

	gridA = make(map[string]int)

	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)

	bringPresent(x, y, "a")
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
		bringPresent(x, y, "a")
	}

	// get answer a
	for _, v := range gridA {
		if v > 0 {
			a++
		}
	}

	return a
}

func solveB(filename string) int {
	var b int
	var dx, dy int
	var xs, ys int // santa
	var xr, yr int // robo

	gridB = make(map[string]int)
	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)

	bringPresent(0, 0, "b")
	for i, r := range s {
		dx = 0
		dy = 0
		switch r {
		case '^':
			dy = 1
		case '>':
			dx = 1
		case 'v':
			dy = -1
		case '<':
			dx = -1
		}
		if i%2 == 0 {
			// santa's turn
			xs += dx
			ys += dy
			bringPresent(xs, ys, "b")
		} else {
			// robo's turn
			xr += dx
			yr += dy
			bringPresent(xr, yr, "b")
		}
	}

	// get answer a
	for _, v := range gridB {
		if v > 0 {
			b++
		}
	}
	return b
}

func bringPresent(x int, y int, grid string) {
	co := fmt.Sprintf("%v,%v", x, y)
	if grid == "a" {
		gridA[co]++
	} else if grid == "b" {
		gridB[co]++
	} else {
		panic("wrong grid: " + grid)
	}
}
