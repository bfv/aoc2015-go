package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	value   int
	inputA  string
	inputB  string
	operand string
	target  string
}

type program []instruction

//type gates map[string]int

var booklet program

func main() {

	//gates = map[string]int{}
	booklet := program{}

	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\r\n")

	processInput(lines)

	solveA("a")

	fmt.Println()
	for _, instr := range booklet {
		if instr.operand == "assign" {
			fmt.Printf("assign %v -> %v\n", instr.value, instr.target)
		}
	}
}

func solveA(letter string) int {
	return findA(letter)
}

func findA(letter string) int {

}

func processInput(lines []string) {

	for _, s := range lines {
		// fmt.Println(s)
		parts := strings.Split(s, " ")
		instr := instruction{value: -1}
		if v, err := strconv.Atoi(parts[0]); err == nil {
			instr.value = v
			instr.operand = "assign"
			instr.target = parts[2]
		} else {
			if parts[0] == "lx" {
				fmt.Printf("found, len: %v\n", len(parts))
			}
			switch len(parts) {
			case 3:
				instr.operand = "assign"
				instr.inputA = parts[0]
				instr.target = parts[2]
			case 4:
				instr.operand = parts[0]
				instr.inputA = parts[1]
				instr.target = parts[3]
			case 5:
				instr.inputA = parts[0]
				instr.operand = parts[1]
				instr.inputB = parts[2]
				instr.target = parts[4]
			}
		}
		booklet = append(booklet, instr)
	}
}
