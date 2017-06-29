package main

import (
	"strings"
	"fmt"
	"sort"
)

func hasNoOverlapChar(word1 string, word2 string) bool {
	//for _, char := range word1 {
	//	if strings.Contains(word2, string(char)) {
	//		return false
	//	}
	//}
	//return true
	fmt.Println(word1, word2, strings.ContainsAny(word1, word2))
	return !strings.ContainsAny(word1, word2)
}

// A data structure to hold key/value pairs
type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func maxProduct(words []string) int {
	sortedByLength := make(map[string]int)
	for _, word := range words {
		sortedByLength[word] = len(word)
	}
	// sort the sequence based on values ie occurrences
	p := make(PairList, len(sortedByLength))

	i := 0
	for k, v := range sortedByLength {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p)) // sorted from longest to shortest word

	max := 0
	limit := len(p) - 1
	for i:=0;i<=limit;i++ {
		for j := 0; j <= limit; j++ {
			if i == j {
				continue
			}
			if len(p[i].Key) * len(p[j].Key) > max {
				if hasNoOverlapChar(p[i].Key, p[j].Key) {
					max = len(p[i].Key) * len(p[j].Key)
					limit = j
					break
				}
			}

		}
	}

	return max
}

func main() {
	fmt.Println(hasNoOverlapChar("hello", "ole"))
	fmt.Println(hasNoOverlapChar("hello", "abba"))

	fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}))
}
