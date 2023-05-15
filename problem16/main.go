package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)  // Before pi
		quickSort(arr, pi+1, high) // After pi
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{9, 4, 7, 2, 1, 5, 8, 3, 6}
	fmt.Println("Input :", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:", arr)
}
