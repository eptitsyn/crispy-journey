package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
//done

func createSet(sequence []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range sequence {
		set[item] = true
	}
	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := createSet(sequence)
	for item := range set {
		fmt.Println(item)
	}
}
