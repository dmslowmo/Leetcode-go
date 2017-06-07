package main

import (
	"fmt"
	"sort"
)
func arrayPairSum(nums []int) int {
	sortedNum := nums
	sort.Ints(sortedNum)
	var sum = 0
	for i:=0;i<len(nums);i+=2 {
		sum += sortedNum[i]
	}
	return sum
}

func main() {
	nums := []int{1,4,3,2}
	fmt.Println(arrayPairSum(nums))
}
