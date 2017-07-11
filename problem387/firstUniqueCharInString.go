package main

import (
	"fmt"
	"math"
)

func firstUniqueChar(s string) int {
	if s == "" {
		return -1
	}
	nonRepeated := make(map[byte]int)
	for i:=0;i<len(s);i++ {
		if _, ok := nonRepeated[s[i]]; ok {
			nonRepeated[s[i]] = -1
			continue
		}
		nonRepeated[s[i]] = i
	}
	min := math.MaxInt64
	for _,v := range nonRepeated {
		if v != -1 && v < min {
			min = v
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
