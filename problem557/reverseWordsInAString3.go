package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	chars := []rune(s)
	for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	reversed := ""
	for i:=0; i<len(words); i++ {
		reversed += reverse(words[i])
		if i != len(words)-1 {
			reversed += " "
		}
	}
	return reversed
}

func main() {
	fmt.Println(reverseWords("abcd efgh dcba hgfe"))
}
