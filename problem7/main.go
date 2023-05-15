package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/
//done

func writeToMap(m map[int]string, key int, value string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	m[key] = value
	mu.Unlock()
}

func main() {
	m := make(map[int]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(m, i, strconv.Itoa(i), &wg, &mu)
	}

	wg.Wait()

	fmt.Println(m)
}
