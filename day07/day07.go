package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type wire struct {
	name string
}

type instruction struct {
	wireA   wire
	wireB   wire
	operand string
	unary   bool
	wireOut wire
	done    bool
}

var register map[string]uint16

type program []instruction

//type gates map[string]int

var booklet program

func main() {

	//gates = map[string]int{}
	booklet = program{}
	register = map[string]uint16{}

	buf, _ := ioutil.ReadFile("_input.txt")
	lines := strings.Split(string(buf), "\r\n")

	processInput(lines)

	solveA("a")

	fmt.Println(booklet)
}

func solveA(letter string) uint16 {
	a := iterateIntructions()
	displayRegister()
	return a
}

func iterateIntructions() uint16 {
	var a uint16
	found := false
	done := false
	for !found && !done {
		displayRegister()
		done = true
		for i, instr := range booklet {
			if !instr.done {
				booklet[i] = execute(instr)
				done = false
			}
		}
		a, found = register["a"]
	}
	return a
}

func execute(instr instruction) instruction {

	if !hasValue(instr.wireA) || (!instr.unary && !hasValue(instr.wireB)) {
		return instr
	}

	var v uint16
	switch instr.operand {
	case "AND":
		v = register[instr.wireA.name] & register[instr.wireB.name]
	case "OR":
		v = register[instr.wireA.name] | register[instr.wireB.name]
	case "NOT":
		v = ^register[instr.wireA.name]
	case "RSHIFT":
		v = register[instr.wireA.name] >> register[instr.wireB.name]
	case "LSHIFT":
		v = register[instr.wireA.name] << register[instr.wireB.name]
	}
	instr.done = true
	register[instr.wireOut.name] = v
	return instr
}

func hasValue(w wire) bool {
	_, valueFound := register[w.name]
	return valueFound
}

func processInput(lines []string) {

	for _, s := range lines {

		parts := strings.Split(s, " ")

		if v, err := strconv.Atoi(parts[0]); err == nil {
			register[parts[2]] = uint16(v)
		} else {
			instr := instruction{}

			switch len(parts) {
			case 3:
				instr.operand = "assign"
				instr.wireA.name = parts[0]
				instr.wireOut.name = parts[2]
				instr.unary = true
			case 4:
				instr.operand = parts[0]
				instr.wireA.name = parts[1]
				instr.wireOut.name = parts[3]
				instr.unary = true
			case 5:
				instr.wireA.name = parts[0]
				instr.operand = parts[1]
				instr.wireB.name = parts[2]
				instr.wireOut.name = parts[4]
				instr.unary = false
			}
			booklet = append(booklet, instr)
		}

	}
}

func displayRegister() {
	fmt.Println("\nregisters")
	for name, v := range register {
		fmt.Printf("%2v: %v\n", name, v)
	}
}
