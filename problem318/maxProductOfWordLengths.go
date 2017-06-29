package main

import (
	"strings"
	"fmt"
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

func maxProduct(words []string) int {
	wordMap := make(map[string][]string)

	for i:=0;i<len(words);i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if hasNoOverlapChar(words[i], words[j]) {
				if wordlist, ok := wordMap[words[i]]; ok {
					wordMap[words[i]] = append(wordlist, words[j])
				} else {
					wordMap[words[i]] = []string{words[j]}
				}
			}
		}
	}

	max := 0
	for key, wordlist := range wordMap {
		for _, word := range wordlist {
			product := len(key) * len(word)
			if product > max {
				max = product
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
