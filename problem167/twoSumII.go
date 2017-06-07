package main

import "fmt"

// given: the numbers array is sorted
//        no 2 numbers can be used twice
//        second's index > first's index
func twoSum1(numbers []int, target int) []int { // this is slow
	for i := 0; i < len(numbers); i++ {
		dist := 1

		for {
			if (i + dist >= len(numbers)) {
				break
			}
			first := numbers[i]
			second := numbers[i+dist]
			fmt.Println([]int{first, second})
			if (second + first > target) {
				break;
			}
			if (second + first == target) {
				return []int{i+1, i+dist+1}
			}
			dist++
		}
	}
	return []int{}
}

func twoSum(numbers []int, target int) []int { // this is fast
	fidx := 0
	bidx := len(numbers) - 1
	for {
		if (numbers[fidx] + numbers[bidx] > target) {
			bidx--
		}
		if (numbers[fidx] + numbers[bidx] < target) {
			fidx++
		}
		if (numbers[fidx] + numbers[bidx] == target) {
			return []int{fidx+1, bidx+1}
		}
		if (fidx >= bidx) {
			break
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15, 19}
	fmt.Println(twoSum(numbers, 9))
	fmt.Println()
	fmt.Println(twoSum(numbers, 23))
	fmt.Println()
	fmt.Println(twoSum(numbers, 22))
	fmt.Println()
	fmt.Println(twoSum(numbers, 26))
}
