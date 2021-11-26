package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//var b int

	input := "yzbqklnj"

	start := time.Now()
	a := solve(input)
	ta := time.Since(start)

	fmt.Printf("day 02, a: %v, t: %v", a, ta)
}

func solve(input string) int {
	var i int

	for {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		s := hex.EncodeToString(hash[:])
		if strings.HasPrefix(s, "00000") {
			break
		}
		i++
	}
	return i
}
