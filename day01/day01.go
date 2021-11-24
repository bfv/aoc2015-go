package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	a, b := solve("input.txt")
	fmt.Printf("day 01, a: %v, b: %v", a, b)
}

func solve(filename string) (int, int) {
	var a, b int

	buf, _ := ioutil.ReadFile(filename)
	s := string(buf)
	for i, r := range s {
		if r == '(' {
			a++
		} else if r == ')' {
			a--
		}
		if b == 0 && a == -1 {
			b = i + 1
		}
	}

	return a, b
}
