package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
//done

func sendData(ch chan<- int) {
	i := 1
	for {
		ch <- i
		i++
		time.Sleep(400 * time.Millisecond)
	}
}

func receiveData(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	var n int
	fmt.Print("Enter N: ")
	if _, err := fmt.Scanf("%d", &n); err != nil {
		panic(err)
	}
	ch := make(chan int)
	go sendData(ch)
	go receiveData(ch)

	time.Sleep(time.Duration(n) * time.Second)
	close(ch)
}
