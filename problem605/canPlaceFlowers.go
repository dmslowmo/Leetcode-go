package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	zeros := 0
	flower := 0
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			flower++
		}
		return n - flower <= 0
	}
	for i:=0;i<2;i++ {
		if flowerbed[i] == 0 {
			zeros++
		} else {
			zeros = 0
		}
		if i == 1 && zeros == 2 {
			flower++
			zeros = 1
		}
	}
	if len(flowerbed) == 2 {
		return n - flower <= 0
	}

	for i:=2;i<len(flowerbed)-2;i++ {
		if flowerbed[i] == 0 {
			zeros++

		} else {
			zeros = 0
		}
		if zeros == 3 {
			flower++
			zeros = 1
		}
	}
	for i:=len(flowerbed)-2;i<len(flowerbed);i++ {
		if flowerbed[i] == 0 {
			zeros++

		} else {
			zeros = 0
		}
		if zeros == 3 {
			flower++
			zeros = 1
		} else if zeros == 2 {
			if i == len(flowerbed)-1 {
				flower++
			}
		}
	}
	return n - flower <= 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,0,1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,0,1}, 2) == true)
	fmt.Println(canPlaceFlowers([]int{0,0,1,0,1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{0,1,0,0}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0}, 2) == true)
	fmt.Println(canPlaceFlowers([]int{0}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1}, 0) == true)
	fmt.Println(canPlaceFlowers([]int{0,0}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{0,0,0,0}, 3) == false)
}
