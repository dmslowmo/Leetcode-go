package main

import "fmt"

func distributeCandies(candies []int) int {
	sister := make(map[int]int)
	for i:=0;i<len(candies);i++ {
		sister[candies[i]] = candies[i]
		if len(sister) == len(candies)/2 {
			break
		}
	}
	return len(sister)
}

func main() {
	candies := []int{1,1,2,2,3,3}
	fmt.Println(distributeCandies(candies))
}
