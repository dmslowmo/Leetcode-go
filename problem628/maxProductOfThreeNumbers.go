package main

import (
	"sort"
	"fmt"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	prodOfTwo := nums[0] * nums[1]
	if nums[len(nums)-1] >=0 {
		if nums[len(nums)-3]*nums[len(nums)-2] > prodOfTwo {
			prodOfTwo = nums[len(nums)-3] * nums[len(nums)-2]
		}
	} else {
		if nums[len(nums)-3]*nums[len(nums)-2] < prodOfTwo {
			prodOfTwo = nums[len(nums)-3] * nums[len(nums)-2]
		}
	}

	return prodOfTwo * nums[len(nums)-1]
}

func main() {
	var nums []int
	var max int
	nums = []int{1,2,3,4}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == 24)

	nums = []int{4,2,3,4}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == 48)

	nums = []int{1,4,4,4}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == 64)

	nums = []int{-1,-2,-3,-4}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == -6)

	nums = []int{-1,-2,-3,-4,60}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == 720)

	nums = []int{722,634,-504,-379,163,-613,-842,-578,750,951,-158,30,-238,-392,-487,-797,-157,-374,999,-5,-521,-879,-858,382,626,803,-347,903,-205,57,-342,186,-736,17,83,726,-960,343,-984,937,-758,-122,577,-595,-544,-559,903,-183,192,825,368,-674,57,-959,884,29,-681,-339,582,969,-95,-455,-275,205,-548,79,258,35,233,203,20,-936,878,-868,-458,-882,867,-664,-892,-687,322,844,-745,447,-909,-586,69,-88,88,445,-553,-666,130,-640,-918,-7,-420,-368,250,-786}
	fmt.Println(nums)
	max = maximumProduct(nums)
	fmt.Println(max, max == 943695360)
}
