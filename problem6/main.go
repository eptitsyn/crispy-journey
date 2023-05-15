package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
//done

func chanGoroutine(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("goroutine exiting...")
			return
		default:
			fmt.Println("goroutine working...")
			time.Sleep(time.Second)
		}
	}
}

func ctxGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine exiting...")
			return
		default:
			// Do some work
			fmt.Println("goroutine working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stopCh := make(chan struct{})
	go chanGoroutine(stopCh)
	time.Sleep(2 * time.Second)
	close(stopCh)
	time.Sleep(1 * time.Second)
	fmt.Println("finish fisrt")

	//2nd
	ctx, cancel := context.WithCancel(context.Background())
	go ctxGoroutine(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("finish second")
}
