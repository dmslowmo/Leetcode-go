package main

import (
	"math"
	"fmt"
)

func judgeSquareSum(c int) bool {

	cf := float64(c)
	for a := float64(0);a*a<=float64(c/2);a++ {
		b := math.Floor(math.Sqrt(cf - a*a))
		if int(a*a+b*b) == c {
			fmt.Println("a=", int(a), " b=", int(b))
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(5, ": ", judgeSquareSum(5))
	fmt.Println(10, ": ", judgeSquareSum(10))
	fmt.Println(100, ": ", judgeSquareSum(100))
	fmt.Println(101, ": ", judgeSquareSum(101))
	fmt.Println(99, ": ", judgeSquareSum(99))
	fmt.Println(2, ": ", judgeSquareSum(2))
}
