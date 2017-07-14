package main

import (
	"fmt"
	"sort"
	"math"
)

func numToArrayInReverse(n int) []int {
	num := n
	arr := []int{}
	for {
		if num == 0 {
			break
		}
		arr = append(arr, num%10)
		num /= 10
	}
	return arr
}

func sortDescending(arr []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
}

func arrayInReverseToNum(arr []int) int {
	num := 0
	multiplier := 1
	for _, digit := range arr {
		num += digit * multiplier
		multiplier *= 10
	}

	if num > math.MaxInt32 {
		return -1
	}
	return num
}

func nextGreaterElement(n int) int {
	backToFront := numToArrayInReverse(n)
	swappedAt := math.MaxInt32
	swappedWith := 0
	limit := len(backToFront)
	for i:=0;i<limit-1;i++ {
		for j:=i+1;j<limit;j++ {
			if backToFront[i] > backToFront[j] {
				if j < swappedAt {
					swappedWith = i
					swappedAt = j
					limit = j
					break
				}
			}
		}
	}

	if swappedAt == math.MaxInt32 {
		return -1
	}

	backToFront[swappedWith], backToFront[swappedAt] = backToFront[swappedAt], backToFront[swappedWith]

	if swappedAt > 0 {
		slice1 := backToFront[:swappedAt]
		sortDescending(slice1)
		slice2 := backToFront[swappedAt:]
		backToFront = append(slice1, slice2...)
	}
	return arrayInReverseToNum(backToFront)
}

func main() {
	nextGreater := nextGreaterElement(12345)
	fmt.Println(12345, nextGreater, nextGreater == 12354)
	fmt.Println()
	nextGreater = nextGreaterElement(12354)
	fmt.Println(12354, nextGreater, nextGreater == 12435)
	fmt.Println()
	nextGreater = nextGreaterElement(21352)
	fmt.Println(21352, nextGreater, nextGreater == 21523)
	fmt.Println()
	nextGreater = nextGreaterElement(23152)
	fmt.Println(23152, nextGreater, nextGreater == 23215)
	fmt.Println()
	nextGreater = nextGreaterElement(23159)
	fmt.Println(23159, nextGreater, nextGreater == 23195)
	fmt.Println()
	nextGreater = nextGreaterElement(23451)
	fmt.Println(23451, nextGreater, nextGreater == 23514)
	fmt.Println()
	nextGreater = nextGreaterElement(21)
	fmt.Println(21, nextGreater, nextGreater == -1)
	fmt.Println()
	nextGreater = nextGreaterElement(1999999999)
	fmt.Println(1999999999, nextGreater, nextGreater == -1)
}
