package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/
//done

func reverseWords(input string) string {
	words := strings.Split(input, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	reversed := strings.Join(words, " ")
	return reversed
}

func main() {
	input := "snow dog sun"
	result := reverseWords(input)
	fmt.Printf("%s - %s\n", input, result)
}
