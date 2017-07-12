package main

import (
	"fmt"
	"math"
	"strings"
)

func firstUniqueChar(s string) int {
	if s == "" {
		return -1
	}
	var count [26]int
	for i:=0;i<len(s);i++ {
		count[int(s[i])-97]++
	}
	min := math.MaxInt64
	for index,v := range count {
		if v == 1 {
			pos := strings.IndexByte(s, byte(index+97))
			if pos < min {
				min = pos
			}
		}
	}

	if min == math.MaxInt64 {
		return -1
	}
	return min
}

func main() {
	fmt.Println(firstUniqueChar("leetcode"))
	fmt.Println(firstUniqueChar("leetcodel"))
}
