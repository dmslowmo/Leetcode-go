package main

import (
	"fmt"
	"math"
)

func getFactors(num int) []int {
	factors := make([]int, 2)
	a := num
	b := int(math.Floor(math.Sqrt(float64(num))))
	for {
		if num % b == 0 {
			factors[0] = b
			a = num/b
			factors[1] = a
			fmt.Println(factors)
		}
		b = b + 1
		if a < b {
			break
		}
	}

	return factors
}

func constructRectangle(area int) []int {
	if area == 2 {
		return []int{2,1}
	}
	factors := getFactors(area)
	L := factors[0]
	W := factors[1]
	return []int{L, W}
}

func main() {
	lxw := constructRectangle(6)
	fmt.Println(lxw)
	lxw = constructRectangle(9)
	fmt.Println(lxw)
	lxw = constructRectangle(18)
	fmt.Println(lxw)
	lxw = constructRectangle(10000000)
	fmt.Println(lxw)
	lxw = constructRectangle(2)
	fmt.Println(lxw)
	lxw = constructRectangle(4)
	fmt.Println(lxw)

}
