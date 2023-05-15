package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(set1, set2 map[string]bool) map[string]bool {
	intersect := make(map[string]bool)
	for item := range set1 {
		if set2[item] {
			intersect[item] = true
		}
	}
	return intersect
}

func main() {
	set1 := map[string]bool{"cat": true, "dog": true, "mice": true}
	set2 := map[string]bool{"cat": true, "mice": true, "rat": true}
	result := intersection(set1, set2)
	for item := range result {
		fmt.Println(item)
	}
}
