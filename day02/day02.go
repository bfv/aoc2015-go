package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(b), "\r")

}
