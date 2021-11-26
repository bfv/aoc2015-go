package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b int

	input := "yzbqklnj"
	a := solve(input)

	fmt.Printf("day 02, a: %v, b: %v", a, b)
}

func solve(input string) int {
	var i int
	var a int
	for a == 0 {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		s := hex.EncodeToString(hash[:])
		if strings.HasPrefix(s, "00000") {
			a = i
		}
		i++
	}
	return a
}
