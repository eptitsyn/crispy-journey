package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/
//done

func removeElement(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	index := 1
	numbers = removeElement(numbers, index)
	fmt.Println(numbers)
}
