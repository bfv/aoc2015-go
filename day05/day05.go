package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\r\n")

	start := time.Now()
	a, b := solve(lines)
	ta := time.Since(start)

	fmt.Printf("day 05, a: %v, b: %v, t: %v\n", a, b, ta)

}

func solve(lines []string) (int, int) {

	niceA := 0
	niceB := 0

	for _, s := range lines {
		if checkNice(s) {
			niceA++
		}
		if checkNiceB(s) {
			niceB++
		}
	}

	return niceA, niceB
}

type tuple struct {
	count   int
	prevPos int
}

func checkNiceB(line string) bool {

	repeat := false
	double := false

	tuples := make(map[string]tuple)

	fmt.Println(line)
	for i := range line {
		if i > 0 {

			if i < len(line)-1 && line[i-1] == line[i+1] /*&& line[i] != line[i+1]*/ {
				fmt.Printf("  repeat: %v\n", string(line[i-1]))
				repeat = true
			}

			// prevent aaa type overlapping
			if i < len(line)-1 && line[i-1] == line[i] && line[i] == line[i+1] {
				continue
			}

			pair := string(line[i-1]) + string(line[i])
			if t, ok := tuples[pair]; !ok {
				t = tuple{1, i}
				tuples[pair] = t
			} else {
				t.count += 1
				t.prevPos = i
				tuples[pair] = t
			}
		}
	}

	for k, t := range tuples {
		if t.count > 1 {
			fmt.Printf("  double: %v\n", k)
			double = true
		}
	}

	return double && repeat
}

func checkNice(line string) bool {
	var prev rune
	vowels := 0
	same := false
	wrong := false
	for _, s := range line {
		if isVowel(s) {
			vowels++
		}
		if s == prev {
			same = true
		}
		wrong = wrong || isWrong(prev, s)
		prev = s
	}
	return vowels >= 3 && same && !wrong
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func isWrong(prev rune, this rune) bool {
	chars := string(prev) + string(this)
	switch chars {
	case "ab", "cd", "pq", "xy":
		return true
	}
	return false
}
