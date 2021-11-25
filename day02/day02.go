package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\r\n")
	a, b := solve(lines)
	fmt.Printf("day 02, a: %v, b: %v", a, b)
}

func solve(lines []string) (int, int) {
	a, b := 0, -1

	for _, line := range lines {
		p := getInts(line)
		p1 := 2 * (p[0]*p[1] + p[0]*p[2] + p[1]*p[2])
		p2 := getSmallestArea(p)
		a += p1 + p2
	}

	return a, b
}

func getInts(line string) []int {
	ints := [3]int{}

	partsString := strings.Split(line, "x")
	for i, part := range partsString {
		ints[i], _ = strconv.Atoi(part)
	}
	return ints[:]
}

func getSmallestArea(sides []int) int {

	var first, second = math.MaxInt, math.MaxInt
	for _, v := range sides {
		if v < first {
			second = first
			first = v
		} else if v < second {
			second = v
		}
	}
	return first * second
}
