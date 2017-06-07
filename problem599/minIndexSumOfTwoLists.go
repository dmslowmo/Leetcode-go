package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	restaurant := make(map[string]int)
	for i:=0; i<len(list1); i++ {
		restaurant[list1[i]] = i
	}
	common := []string{}
	for i:=0; i<len(list2); i++ {
		if index, ok := restaurant[list2[i]]; ok {
			restaurant[list2[i]] = index + i
			common = append(common, list2[i])
		}
	}
	min := restaurant[common[0]]
	mins := make(map[int][]string)
	for i:=0;i<len(common);i++ {
		if min >= restaurant[common[i]] {
			min = restaurant[common[i]]
			if arr, ok := mins[min]; ok {
				mins[min] = append(arr, common[i])
			} else {
				mins = make(map[int][]string)
				mins[min] = []string{common[i]}
			}
		}
	}
	//for k, v := range mins {
	//	fmt.Printf("key[%d] value[%s]\n", k, v)
	//}
	return mins[min]
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
	fmt.Println(findRestaurant([]string{"Shogun", "KFC", "Tapioca Express", "Burger King"}, []string{"KFC", "Shogun", "Burger King"}))
}
