package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var gridA map[string]int
var gridB map[string]int

type operation struct {
	x1, y1 int
	x2, y2 int
	action string
}

func main() {

	start := time.Now()

	gridA = make(map[string]int)
	gridB = make(map[string]int)

	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\r\n")

	for _, line := range lines {

		parts := strings.Split(line, " ")
		if parts[0] == "turn" {
			parts = parts[1:]
		}

		o := operation{action: parts[0]}
		o.x1, o.y1 = spiltXY(parts[1:2][0])
		o.x2, o.y2 = spiltXY(parts[3:][0])

		applyToGrid(o)
	}

	ta := time.Since(start)
	b := getBrightness()

	fmt.Printf("day 05, a: %v, b: %v, t: %v\n", len(gridA), b, ta)
}

func spiltXY(co string) (int, int) {
	var x, y int
	parts := strings.Split(co, ",")
	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])
	return x, y
}

func applyToGrid(o operation) {
	for x := o.x1; x <= o.x2; x++ {
		for y := o.y1; y <= o.y2; y++ {
			setCellA(x, y, o.action)
			setCellB(x, y, o.action)
		}
	}
}

// note: the grid only the on coordinates
func setCellA(x int, y int, action string) {
	co := fmt.Sprintf("%v,%v", x, y)
	_, ok := gridA[co]
	switch true {
	case ok && (action == "off" || action == "toggle"):
		delete(gridA, co)
	case ok && action == "on":
		gridA[co] = 1
	case !ok && (action == "on" || action == "toggle"):
		gridA[co] = 1
	}
}

func setCellB(x int, y int, action string) {
	co := fmt.Sprintf("%v,%v", x, y)
	_, ok := gridB[co]
	if !ok {
		gridB[co] = 0
	}
	switch action {
	case "off":
		gridB[co] = gridB[co] - 1
		if (gridB[co]) <= 0 {
			delete(gridB, co)
		}
	case "on":
		gridB[co] += 1
	case "toggle":
		gridB[co] += 2
	}
}

func getBrightness() int {
	b := 0
	for _, v := range gridB {
		b += v
	}
	return b
}
